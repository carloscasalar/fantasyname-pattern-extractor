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
	resultingPattern := new(pattern)
	for _, token := range tokenChain.Tokens() {
		resultingPattern.add(t.toPattern(token))
	}

	return resultingPattern
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
