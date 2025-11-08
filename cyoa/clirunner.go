package main

import (
	"fmt"
	"os"
)

type CliRunner struct{}

func (cr CliRunner) Start(provider *Provider) {
	cr.displayArcText(provider, "intro")
}

func (cr CliRunner) displayArcText(provider *Provider, arcName string) {
	arc, err := provider.WriteTemplatedText(os.Stdout, arcName)
	if err != nil {
		panic(err)
	}
	if len(arc.Options) == 0 {
		return
	}
	fmt.Print("Your Option: ")
	var optionNum int
	fmt.Scan(&optionNum)
	for idx, option := range arc.Options {
		if idx == optionNum {
			cr.displayArcText(provider, option.Arc)
		}
	}
}
