package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SupportedFileType int

type UnSupportedFileType struct {
	Code    int
	Message string
}

func (e *UnSupportedFileType) Error() string {
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}

const (
	JSON_FILETYPE SupportedFileType = iota
)

type Parser interface {
	Parse(data []byte) (*Book, error)
}

type JsonParser struct{}

func (jp JsonParser) Parse(data []byte) (*Book, error) {
	var book Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func getBookFromFile(filename string, fileType int) (*Book, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parser, err := getParser(SupportedFileType(fileType))
	if err != nil {
		return nil, err
	}

	book, err := parser.Parse(data)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func getParser(fileType SupportedFileType) (Parser, error) {
	switch fileType {
	case JSON_FILETYPE:
		return JsonParser{}, nil

	default:
		return nil, &UnSupportedFileType{Code: 400, Message: fmt.Sprintf("Input file type not supported")}
	}
}
