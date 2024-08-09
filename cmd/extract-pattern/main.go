package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"

	"github.com/jessevdk/go-flags"

	"github.com/s0rg/fantasyname"
)

func main() {
	opts := readOptionsOrFail()

	tokenizedSample, err := tokenizer.Tokenize(opts.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pattern := transformer.NewNaiveTransformer().Transform(*tokenizedSample)
	capitalizedPattern := fmt.Sprintf("!%s", pattern.String())

	if opts.NumberOfOutputsToGenerate == 0 {
		titleBox := lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).
			Foreground(lipgloss.AdaptiveColor{Light: "202", Dark: "252"})
		patternBox := lipgloss.NewStyle().
			MarginLeft(1).
			Foreground(lipgloss.AdaptiveColor{Light: "#3C3C3C", Dark: "#04B575"})
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left, titleBox.Render("PATTERN:"), patternBox.Render(capitalizedPattern)))
		return
	}
	examples := make([]string, opts.NumberOfOutputsToGenerate)

	for i := 0; i < int(opts.NumberOfOutputsToGenerate); i++ {
		gen, err := fantasyname.Compile(capitalizedPattern, fantasyname.Collapse(true), fantasyname.RandFn(rand.Intn))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		examples[i] = gen.String()
	}

	rows := [][]string{
		{capitalizedPattern, commaSeparated(examples)},
	}

	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(lipgloss.AdaptiveColor{Light: "202", Dark: "252"}).Bold(true)
	patternRow := baseStyle.Foreground(lipgloss.AdaptiveColor{Light: "#3C3C3C", Dark: "#04B575"})
	columnExample := patternRow.MaxWidth(120).Italic(true)
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}
			if col == 0 {
				return patternRow
			}
			return columnExample
		}).
		Headers("PATTERN", "EXAMPLES").
		Rows(rows...)
	fmt.Println(t)
}

func commaSeparated(examples []string) string {
	if len(examples) == 0 {
		return ""
	}

	joinedString := new(strings.Builder)
	for i, example := range examples {
		joinedString.WriteString(example)
		if i < len(examples)-1 {
			joinedString.WriteString(", ")
		}
	}

	return joinedString.String()
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

type Ops struct {
	Name                      string `short:"n" long:"name" description:"Sample name to extract pattern from" required:"true"`
	NumberOfOutputsToGenerate uint   `short:"o" long:"number-of-outputs" description:"Number of outputs to generate" default:"0"`
}
