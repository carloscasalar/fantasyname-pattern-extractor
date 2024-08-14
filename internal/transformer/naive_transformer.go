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
	if token.IsVowel() {
		if token.IsAcuteAccented() {
			return sequenceVowelAcuteAccented
		}
		if token.IsGraveAccented() {
			return sequenceVowelGraveAccented
		}
		if token.IsCircumflexAccented() {
			return sequenceVowelCircumflexAccented
		}
		if token.IsDieresisAccented() {
			return sequenceVowelDieresisAccented
		}
		return sequenceVowel
	}

	if token.IsConsonant() {
		if token.IsCedilla() {
			return sequenceCedilla
		}

		if token.IsTildeN() {
			return sequenceTildeN
		}
		return sequenceConsonant
	}

	if token.IsApostrophe() {
		return sequenceApostrophe
	}

	if token.IsHyphen() {
		return sequenceHyphen
	}

	return emptySequence
}
