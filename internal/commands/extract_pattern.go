package commands

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
)

type Transformer interface {
	Transform(tokenChain tokenizer.TokenChain) transformer.Pattern
}

type ExtractPattern struct {
	transformer Transformer
}

func NewExtractPattern(transformer Transformer) *ExtractPattern {
	return &ExtractPattern{transformer: transformer}
}

func (pe *ExtractPattern) Execute(name string) (string, error) {
	tokenizedSample, err := tokenizer.Tokenize(name)
	if err != nil {
		return "", err
	}

	pattern := pe.transformer.Transform(*tokenizedSample)
	return pattern.String(), nil
}
