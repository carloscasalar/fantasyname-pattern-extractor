package transformer

import "github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"

// ContextAwareTransformer Translates starting consonant or group of constants to <B>, two to three consecutive middle
// consonants to <C>, two to three consecutive vowels to V and the rest using the vowel proximity criteria.
type ContextAwareTransformer struct {
}

// NewContextAwareTransformer creates a new ContextAwareTransformer.
func NewContextAwareTransformer() *ContextAwareTransformer {
	return &ContextAwareTransformer{}
}

func (t *ContextAwareTransformer) Transform(tokenChain tokenizer.TokenChain) Pattern {
	resultingPattern := new(pattern)
	for i, token := range tokenChain.Tokens() {
		isFirstToken := i == 0
		resultingPattern.add(t.toPattern(token, isFirstToken))
	}

	return resultingPattern
}

func (t *ContextAwareTransformer) toPattern(token tokenizer.Token, isFirstToken bool) patternSequence {
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
		if isFirstToken {
			return sequenceConsonantStarting
		}
		return sequenceConsonant
	case tokenizer.TokenTildeN:
		if isFirstToken {
			return sequenceTildeNStarting
		}
		return sequenceTildeN
	case tokenizer.TokenCedilla:
		if isFirstToken {
			return sequenceCedillaStarting
		}
		return sequenceCedilla
	case tokenizer.TokenApostrophe:
		return sequenceApostrophe
	case tokenizer.TokenHyphen:
		return sequenceHyphen
	default:
		return emptySequence
	}
}
