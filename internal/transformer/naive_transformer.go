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
		pattern.add(sequenceVowel)
	case tokenizer.TokenVowelAcuteAccented:
		pattern.add(sequenceVowelAcuteAccented)
	case tokenizer.TokenVowelGraveAccented:
		pattern.add(sequenceVowelGraveAccented)
	case tokenizer.TokenVowelCircumflexAccented:
		pattern.add(sequenceVowelCircumflexAccented)
	case tokenizer.TokenVowelDieresisAccented:
		pattern.add(sequenceVowelDieresisAccented)
	case tokenizer.TokenConsonant:
		pattern.add(sequenceConsonant)
	case tokenizer.TokenTildeN:
		pattern.add(sequenceTildeN)
	case tokenizer.TokenCedilla:
		pattern.add(sequenceCedilla)
	case tokenizer.TokenApostrophe:
		pattern.add(sequenceApostrophe)
	case tokenizer.TokenHyphen:
		pattern.add(sequenceHyphen)
	default:
	}

	return pattern
}
