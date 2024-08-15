package commands

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
)

// ExtractPattern is a command that extracts a pattern from a given name using the provided transformer
type ExtractPattern struct {
	transformer Transformer
}

// NewExtractPattern creates a new ExtractPattern command
func NewExtractPattern(transformer Transformer) *ExtractPattern {
	return &ExtractPattern{transformer: transformer}
}

func (pe *ExtractPattern) Execute(name string) (string, error) {
	tokenizedSample, err := tokenizer.Tokenize(name)
	if err != nil {
		return "", err
	}

	pattern := pe.transformer.Transform(*tokenizedSample)
	return pattern.Capitalize(), nil
}
