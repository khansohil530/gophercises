package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/khansohil530/gophercises/urlshort"
)

func main() {
	var ymlFilename string
	parseFlags(&ymlFilename)

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yml, err := readFileBytes(ymlFilename)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseFlags(ymlFilename *string) {
	flag.StringVar(ymlFilename, "yml", "urls.yml", "YAML Filename of format \n- path: /example \n  url: https://example.com ")
	flag.Parse()
}

func readFileBytes(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}
