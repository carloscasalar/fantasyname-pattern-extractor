package transformer

import "github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"

// VowelProximityTransformer is a transformer that translates to a character for every token in the chain but tries to generate a pattern that groups vowels by strength and accent.
type VowelProximityTransformer struct {
}

// NewVowelProximityTransformer creates a new VowelProximityTransformer.
func NewVowelProximityTransformer() *VowelProximityTransformer {
	return &VowelProximityTransformer{}
}

func (t *VowelProximityTransformer) Transform(tokenChain tokenizer.TokenChain) Pattern {
	resultingPattern := new(pattern)
	for _, token := range tokenChain.Tokens() {
		resultingPattern.add(t.toPattern(token))
	}

	return resultingPattern
}

func (t *VowelProximityTransformer) toPattern(token tokenizer.Token) patternSequence {
	switch token {
	case tokenizer.TokenVowelStrong:
		return sequenceVowelStrong
	case tokenizer.TokenVowelWeak:
		return sequenceVowelWeak
	case tokenizer.TokenVowelStrongAcuteAccented:
		return sequenceVowelStrongAcuteAccented
	case tokenizer.TokenVowelWeakAcuteAccented:
		return sequenceVowelWeakAcuteAccented
	case tokenizer.TokenVowelStrongGraveAccented:
		return sequenceVowelStrongGraveAccented
	case tokenizer.TokenVowelWeakGraveAccented:
		return sequenceVowelWeakGraveAccented
	case tokenizer.TokenVowelStrongCircumflexAccented:
		return sequenceVowelStrongCircumflexAccented
	case tokenizer.TokenVowelWeakCircumflexAccented:
		return sequenceVowelWeakCircumflexAccented
	case tokenizer.TokenVowelStrongDieresisAccented:
		return sequenceVowelStrongDieresisAccented
	case tokenizer.TokenVowelWeakDieresisAccented:
		return sequenceVowelWeakDieresisAccented
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
