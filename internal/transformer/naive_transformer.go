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
		pattern.add("v")
	case tokenizer.TokenVowelAcuteAccented:
		pattern.add("(<v>|á|é|í|ó|ú)")
	case tokenizer.TokenVowelGraveAccented:
		pattern.add("(<v>|à|è|ì|ò|ù)")
	case tokenizer.TokenVowelCircumflexAccented:
		pattern.add("(<v>|â|ê|î|ô|û)")
	case tokenizer.TokenVowelDieresisAccented:
		pattern.add("(<v>|ä|ë|ï|ö|ü)")
	case tokenizer.TokenConsonant:
		pattern.add("c")
	case tokenizer.TokenTildeN:
		pattern.add("(<c>|ñ)")
	case tokenizer.TokenCedilla:
		pattern.add("(<c>|ç)")
	case tokenizer.TokenApostrophe:
		pattern.add("('|)")
	case tokenizer.TokenHyphen:
		pattern.add("(-|)")
	default:
	}

	return pattern
}
