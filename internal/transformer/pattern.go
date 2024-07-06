package transformer

import "strings"

type Pattern struct {
	sequences []patternSequence
}

func (p *Pattern) add(sequence patternSequence) {
	p.sequences = append(p.sequences, sequence)
}

func (p *Pattern) String() string {
	var patternBuilder strings.Builder
	for _, sequence := range p.sequences {
		patternBuilder.WriteString(string(sequence))
	}
	return patternBuilder.String()
}
