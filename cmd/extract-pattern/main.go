package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

func main() {
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
	fmt.Printf("sample: %v\n", opts.Sample)
}

type Ops struct {
	Sample string `short:"s" long:"sample" description:"Sample name to extract pattern from" required:"true"`
}
