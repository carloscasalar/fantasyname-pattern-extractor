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
	case tokenizer.TokenVowelGraveAccented:
		pattern.Add("(<v>|à|è|ì|ò|ù)")
	case tokenizer.TokenVowelCircumflexAccented:
		pattern.Add("(<v>|â|ê|î|ô|û)")
	case tokenizer.TokenVowelDieresisAccented:
		pattern.Add("(<v>|ä|ë|ï|ö|ü)")
	case tokenizer.TokenConsonant:
		pattern.Add("c")
	case tokenizer.TokenTildeN:
		pattern.Add("(<c>|ñ)")
	case tokenizer.TokenCedilla:
		pattern.Add("(<c>|ç)")
	case tokenizer.TokenApostrophe:
		pattern.Add("('|)")
	}

	return pattern
}
