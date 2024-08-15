package transformer

import (
	"fmt"
	"strings"
)

type Pattern interface {
	String() string
	Capitalize() string
}

type pattern struct {
	sequences []patternSequence
}

func (p *pattern) add(sequence patternSequence) {
	p.sequences = append(p.sequences, sequence)
}

func (p *pattern) Capitalize() string {
	return fmt.Sprintf("!%s", p.String())
}

func (p *pattern) String() string {
	var patternBuilder strings.Builder
	for _, sequence := range p.sequences {
		patternBuilder.WriteString(string(sequence))
	}
	return patternBuilder.String()
}
