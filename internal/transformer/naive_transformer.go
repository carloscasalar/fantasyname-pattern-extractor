package transformer

import (
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
)

type NaiveTransformer struct {
}

func NewNaiveTransformer() *NaiveTransformer {
	return &NaiveTransformer{}
}

func (t *NaiveTransformer) Transform(tokenChain tokenizer.TokenChain) Pattern {
	token := tokenChain.Tokens()[0]
	var pattern Pattern
	switch token {
	case tokenizer.TokenVowel:
		pattern.Add("v")

	case tokenizer.TokenVowelAcuteAccented:
		pattern.Add("(<v>|á|é|í|ó|ú)")
	case tokenizer.TokenConsonant:
		pattern.Add("c")
	case tokenizer.TokenTildeN:
		pattern.Add("(<c>|ñ)")
	case tokenizer.TokenCedilla:
		pattern.Add("(<c>|ç)")
	default:
		panic("unhandled default case")
	}

	return pattern
}
