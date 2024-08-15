package main

import (
	"fmt"
	"os"

	"golang.org/x/term"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/ui"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/commands"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
	"github.com/jessevdk/go-flags"
)

func main() {
	opts := readOptionsOrFail()

	extractPatterns := commands.NewExtractPatterns(
		transformer.NewNaiveTransformer(),
		transformer.NewVowelProximityTransformer(),
	)
	patterns, err := extractPatterns.Execute(opts.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var renderer Renderer
	switch opts.NumberOfOutputsToGenerate {
	case 0:
		renderer = ui.NewTitleValueRenderer("PATTERNS:", commaSeparatedValues(patterns))
	default:
		table := new(ui.TwoColumnsTableBuilder).
			WithMaxWidth(getTerminalWidth()).
			WithTitles("PATTERN", "EXAMPLES")

		for _, pattern := range patterns {
			nameExamples, err := commands.GenerateExamples(pattern, opts.NumberOfOutputsToGenerate)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			table.WithRow(pattern, nameExamples)
		}
		renderer = table.Build()
	}

	fmt.Println(renderer.Render())
}

func readOptionsOrFail() Ops {
	var opts Ops
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	return opts
}

func getTerminalWidth() int {
	totalMaxWidth, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		totalMaxWidth = 80
	}
	return totalMaxWidth
}

func commaSeparatedValues(values []string) string {
	valuesStr := ""
	for i, value := range values {
		valuesStr += value
		if i < len(values)-1 {
			valuesStr += ", "
		}
	}
	return valuesStr
}

type Ops struct {
	Name                      string `short:"n" long:"name" description:"Sample name to extract pattern from" required:"true"`
	NumberOfOutputsToGenerate uint   `short:"o" long:"number-of-outputs" description:"Number of outputs to generate" default:"0"`
}

type Renderer interface {
	Render() string
}
