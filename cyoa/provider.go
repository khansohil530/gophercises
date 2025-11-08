package main

import (
	"fmt"
	"html/template"
	"io"
)

type StoryNotFound struct {
	Code    int
	Message string
}

func (e *StoryNotFound) Error() string {
	return fmt.Sprintf("Error code %d: %s", e.Code, e.Message)
}

type ProviderType int

const (
	WEB_PROVIDER ProviderType = iota
	CLI_PROVIDER
)

type Provider struct {
	Book         *Book
	ProviderType ProviderType
	tpl          *template.Template
}

func (p *Provider) Initialize() error {
	templateName := "web.tpl"
	if p.ProviderType == CLI_PROVIDER {
		templateName = "cli.tpl"
	}

	t := template.New(templateName)
	tpl, err := t.ParseFiles(templateName)
	if err != nil {
		return err
	}
	p.tpl = tpl
	return nil
}

func (p *Provider) WriteTemplatedText(w io.Writer, arcName string) (*Story, error) {
	arc, err := p.Book.getArc(arcName)
	if err != nil {
		return nil, err
	}

	err = p.tpl.Execute(w, arc)
	if err != nil {
		return nil, err
	}
	return arc, nil

}
