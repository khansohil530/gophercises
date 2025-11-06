package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	filename, timeLimit := parseFlags()

	csvParser := CSVParser{}
	lines := parseFile(filename, csvParser)

	problems := parseLines(lines)

	timeout := startTimer(timeLimit)

	var score int
	answerChan := make(chan string)
loop:
	for i, p := range problems {
		go func() {
			userAnswer := playQuestion(p.question, i+1)
			answerChan <- userAnswer
		}()
		select {
		case <-timeout:
			fmt.Printf("Times up. ")
			break loop
		case answer := <-answerChan:
			if answer == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", score, len(problems))

}

func parseFlags() (string, int) {
	var filename string
	var timeLimit int
	flag.StringVar(&filename, "csv", "questions.csv",
		"csv filename containing list of questions and answers in format 'question,answer'")
	flag.IntVar(&timeLimit, "limit", 30, "time limit for quiz in seconds")
	flag.Parse()
	return filename, timeLimit
}

type FileParser interface {
	Parse(io.Reader) ([][]string, error)
}

type CSVParser struct{}

func (CSVParser) Parse(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	return reader.ReadAll()
}

func parseFile(filename string, parser FileParser) [][]string {
	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	lines, err := parser.Parse(file)
	checkError(err)
	return lines
}

func parseLines(lines [][]string) []*problem {
	problems := make([]*problem, len(lines))
	for i, line := range lines {
		problems[i] = &problem{question: line[0], answer: strings.TrimSpace(line[1])}
	}
	return problems
}

func playQuestion(question string, questionNo int) string {
	var userInput string
	fmt.Printf("Question #%d: %s = \n", questionNo, question)
	fmt.Scanf("%s\n", &userInput)
	return userInput
}

type problem struct {
	question, answer string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func startTimer(timeLimit int) <-chan time.Time {
	return time.After(time.Duration(timeLimit) * time.Second)
}
