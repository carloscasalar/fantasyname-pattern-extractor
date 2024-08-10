package main

import (
	"fmt"
	"os"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/commands"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"

	"github.com/jessevdk/go-flags"
)

func main() {
	opts := readOptionsOrFail()

	extractPattern := commands.NewExtractPattern(transformer.NewNaiveTransformer())
	pattern, err := extractPattern.Execute(opts.Name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.NumberOfOutputsToGenerate == 0 {
		titleBox := lipgloss.NewStyle().
			Bold(true).
			PaddingLeft(1).
			Foreground(lipgloss.AdaptiveColor{Light: "202", Dark: "252"})
		patternBox := lipgloss.NewStyle().
			MarginLeft(1).
			Foreground(lipgloss.AdaptiveColor{Light: "#3C3C3C", Dark: "#04B575"})
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Left, titleBox.Render("PATTERN:"), patternBox.Render(pattern)))
		return
	}

	nameExamples, err := commands.GenerateExamples(pattern, opts.NumberOfOutputsToGenerate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows := [][]string{
		{pattern, nameExamples},
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
