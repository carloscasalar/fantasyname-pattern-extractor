package tokenizer_test

import (
	"fmt"
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/stretchr/testify/require"
)

func Test_given_strong_vowel(t *testing.T) {
	testCases := []tokenizer.Token{
		tokenizer.TokenVowelStrong,
		tokenizer.TokenVowelStrongAcuteAccented,
		tokenizer.TokenVowelStrongGraveAccented,
		tokenizer.TokenVowelStrongCircumflexAccented,
		tokenizer.TokenVowelStrongDieresisAccented,
	}

	for _, token := range testCases {
		t.Run(fmt.Sprintf("%v token should be strong", token), func(t *testing.T) {
			require.IsType(t, tokenizer.TokenVowel{}, token)

			vowelToken := token.(tokenizer.TokenVowel)
			require.True(t, vowelToken.IsStrong())
			require.False(t, vowelToken.IsWeak())
		})
	}
}

func Test_given_weak_vowel(t *testing.T) {
	testCases := []tokenizer.Token{
		tokenizer.TokenVowelWeak,
		tokenizer.TokenVowelWeakAcuteAccented,
		tokenizer.TokenVowelWeakGraveAccented,
		tokenizer.TokenVowelWeakCircumflexAccented,
		tokenizer.TokenVowelWeakDieresisAccented,
	}

	for _, token := range testCases {
		t.Run(fmt.Sprintf("%v token should be weak", token), func(t *testing.T) {
			require.IsType(t, tokenizer.TokenVowel{}, token)
			// cast to TokenVowel:
			vowelToken := token.(tokenizer.TokenVowel)
			require.True(t, vowelToken.IsWeak())
			require.False(t, vowelToken.IsStrong())
		})
	}
}
