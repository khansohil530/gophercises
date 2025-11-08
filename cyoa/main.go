package main

import (
	"flag"
)

func main() {
	var filename string
	var mode int
	fileType := 0
	parseFlags(&filename, &mode)
	book, err := getBookFromFile(filename, fileType)
	errExit(err)

	// 2. Parse the book into interactive story on browser using html/template
	var runner Runner
	provider := Provider{Book: book}
	switch mode {
	case 1:
		runner = CliRunner{}
		provider.ProviderType = CLI_PROVIDER
	default:
		runner = WebRunner{}
		provider.ProviderType = WEB_PROVIDER
	}
	err = provider.Initialize()
	errExit(err)

	runner.Start(&provider)

}

func parseFlags(filename *string, mode *int) {
	flag.StringVar(filename, "file", "gopher.json", "path to storyline file")
	flag.IntVar(mode, "mode", 0, "run in web (0) or cli (1) mode (default web)")
	flag.Parse()
}

func errExit(err error) {
	if err != nil {
		panic(err)
	}
}
