package tokenizer_test

import (
	"fmt"
	"github.com/carloscasalar/fantasy-pattern-inferrer/internal/tokenizer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_given_string_starting_with_the_consonant(t *testing.T) {
	testCases := []string{
		"B", "C", "D", "F", "G", "H", "J", "K", "L", "M",
		"N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z",
		"b", "c", "d", "f", "g", "h", "j", "k", "l", "m",
		"n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be starting consonant", input), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			require.Len(t, tokens, 1)
			assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0][0])
		})
	}
}

func Test_given_string_starting_with_the_vowel(t *testing.T) {
	testCases := []string{
		"A", "E", "I", "O", "U", "a", "e", "i", "o", "u",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be single vowel", input), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			require.Len(t, tokens, 1)
			assert.Equal(t, tokenizer.TokenVowel, tokens[0][0])
		})
	}
}
