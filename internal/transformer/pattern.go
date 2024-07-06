package transformer

import "strings"

type Pattern struct {
	sequences []string
}

func (p *Pattern) Add(sequence string) {
	p.sequences = append(p.sequences, sequence)
}

func (p *Pattern) String() string {
	var patternBuilder strings.Builder
	for _, sequence := range p.sequences {
		patternBuilder.WriteString(sequence)
	}
	return patternBuilder.String()
}
