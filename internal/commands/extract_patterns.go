package commands

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
)

// ExtractPatterns is a command that uses every provided transformer to extract a pattern from a given name
type ExtractPatterns struct {
	transformers []Transformer
}

// NewExtractPatterns creates a new ExtractPatterns command
func NewExtractPatterns(transformers ...Transformer) *ExtractPatterns {
	return &ExtractPatterns{transformers: transformers}
}

func (pe *ExtractPatterns) Execute(name string) ([]string, error) {
	tokenizedSample, err := tokenizer.Tokenize(name)
	if err != nil {
		return nil, err
	}

	patterns := make([]string, len(pe.transformers))
	for i, transformer := range pe.transformers {
		pattern := transformer.Transform(*tokenizedSample)
		patterns[i] = pattern.Capitalize()
	}

	return patterns, nil
}
