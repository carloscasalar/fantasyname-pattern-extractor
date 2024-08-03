package main

import (
	"fmt"
	"os"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"

	"github.com/jessevdk/go-flags"
)

func main() {
	opts := readOptionsOrFail()

	tokenizedSample, err := tokenizer.Tokenize(opts.Sample)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pattern := transformer.NewNaiveTransformer().Transform(*tokenizedSample)

	fmt.Printf("Pattern: !%v\n", pattern.String())
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
	Sample string `short:"s" long:"sample" description:"Sample name to extract pattern from" required:"true"`
}
