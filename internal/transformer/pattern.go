package transformer

import "strings"

type Pattern struct {
	parts []string
}

func (p *Pattern) Add(patternPart string) {
	p.parts = append(p.parts, patternPart)
}

func (p *Pattern) String() string {
	var patternBuilder strings.Builder
	for _, part := range p.parts {
		patternBuilder.WriteString(part)
	}
	return patternBuilder.String()
}
