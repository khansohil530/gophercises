package main

import (
	"fmt"
	"net/http"
)

type WebRunner struct{}

func (wr WebRunner) Start(provider *Provider) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		wr.rootHandler(provider, writer, request)
	})
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func (wr WebRunner) rootHandler(p *Provider, w http.ResponseWriter, r *http.Request) {
	arcQuery := r.URL.Query()["arc"]
	currArc := "intro"
	if len(arcQuery) > 0 {
		currArc = arcQuery[0]
	}

	_, err := p.WriteTemplatedText(w, currArc)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
