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
	case tokenizer.TokenVowelWeak, tokenizer.TokenVowelStrong:
		return sequenceVowel
	case tokenizer.TokenVowelWeakAcuteAccented, tokenizer.TokenVowelStrongAcuteAccented:
		return sequenceVowelAcuteAccented
	case tokenizer.TokenVowelStrongGraveAccented, tokenizer.TokenVowelWeakGraveAccented:
		return sequenceVowelGraveAccented
	case tokenizer.TokenVowelWeakCircumflexAccented, tokenizer.TokenVowelStrongCircumflexAccented:
		return sequenceVowelCircumflexAccented
	case tokenizer.TokenVowelWeakDieresisAccented, tokenizer.TokenVowelStrongDieresisAccented:
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
