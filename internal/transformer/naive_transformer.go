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
	pattern := new(Pattern)
	pattern.add(t.toPattern(token))

	return *pattern
}

func (t *NaiveTransformer) toPattern(token tokenizer.Token) patternSequence {
	switch token {
	case tokenizer.TokenVowel:
		return sequenceVowel
	case tokenizer.TokenVowelAcuteAccented:
		return sequenceVowelAcuteAccented
	case tokenizer.TokenVowelGraveAccented:
		return sequenceVowelGraveAccented
	case tokenizer.TokenVowelCircumflexAccented:
		return sequenceVowelCircumflexAccented
	case tokenizer.TokenVowelDieresisAccented:
		return sequenceVowelDieresisAccented
	case tokenizer.TokenConsonant:
		return sequenceConsonant
	case tokenizer.TokenTildeN:
		return sequenceTildeN
	case tokenizer.TokenCedilla:
		return sequenceCedilla
	case tokenizer.TokenApostrophe:
		return sequenceApostrophe
	case tokenizer.TokenHyphen:
		return sequenceHyphen
	default:
		return emptySequence
	}
}
