package transformer

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
)

type NaiveTransformer struct {
}

func NewNaiveTransformer() *NaiveTransformer {
	return &NaiveTransformer{}
}

func (t *NaiveTransformer) Transform(tokenChain tokenizer.TokenChain) string {
	token := tokenChain.Tokens()[0]
	var pattern string
	switch token {
	case tokenizer.TokenVowel:
		pattern = "v"
	case tokenizer.TokenConsonant:
		pattern = "c"
	case tokenizer.TokenTildeN:
		pattern = "(<c>|ñ)"
	case tokenizer.TokenCedilla:
		pattern = "(<c>|ç)"
	default:
		panic("unhandled default case")
	}

	return pattern
}
