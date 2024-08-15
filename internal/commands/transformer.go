package commands

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
)

type Transformer interface {
	Transform(tokenChain tokenizer.TokenChain) transformer.Pattern
}
