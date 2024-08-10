package commands

import (
	"fmt"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
)

type Transformer interface {
	Transform(tokenChain tokenizer.TokenChain) transformer.Pattern
}

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
	patternWithFirstLetterUppercased := fmt.Sprintf("!%s", pattern.String())
	return patternWithFirstLetterUppercased, nil
}
