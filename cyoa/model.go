package main

import (
	"errors"
	"fmt"
	"strings"
)

type Book map[string]*Story

func (b Book) String() string {
	var lines []string
	for k, v := range b {
		lines = append(lines, fmt.Sprintf("%v:\n  %v", k, v))
	}
	return strings.Join(lines, "\n")
}

type Story struct {
	Title       string       `json:"title,omitempty"`
	Description []string     `json:"story,omitempty"`
	Options     []*ArcOption `json:"options,omitempty"`
}

func (s Story) String() string {
	return fmt.Sprintf("'Title': %v\n  'Description': %v\n  'Options': %v\n", s.Title, s.Description, s.Options)
}

type ArcOption struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

func (ao ArcOption) String() string {
	return fmt.Sprintf("'Text': %v\n'Arc':%v\n", ao.Text, ao.Arc)
}

func (s *Book) getArc(arcName string) (*Story, error) {
	for currArc, arcData := range *s {
		if currArc == arcName {
			return arcData, nil
		}
	}
	return nil, errors.New("Arc " + arcName + " not found in stories.")
}
