package tokenizer_test

import (
	"fmt"
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			require.Len(t, tokens, 1, "should have only one token")
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
			require.Len(t, tokens, 1, "should have only one token")
			assert.Equal(t, tokenizer.TokenVowel, tokens[0][0])
		})
	}
}

func Test_given_string_starting_with_a_consonant_and_a_vowel_like(t *testing.T) {
	testCases := []string{
		"Ba", "Pe", "Wi", "Qo", "Zu",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be starting consonant and second token a vowel", input), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			require.Len(t, tokens, 1, "should have only one combination of tokens but has %v", len(tokens))
			require.Len(t, tokens[0], 2, "should have two tokens but has %v", len(tokens[0]))
			assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0][0])
			assert.Equal(t, tokenizer.TokenVowel, tokens[0][1])
		})
	}
}

func Test_given_string_starting_with_a_vowel_and_a_consonant_like(t *testing.T) {
	testCases := []string{
		"Ab", "Ep", "Iw", "Oq", "Uz",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be a vowel and the second token a middle consonant", input), func(t *testing.T) {
			tokens, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			require.Len(t, tokens, 1, "should have only one combination of tokens but has %v", len(tokens))
			require.Len(t, tokens[0], 2, "should have two tokens but has %v", len(tokens[0]))
			assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0][0].String())
			assert.Equal(t, tokenizer.TokenMiddleConsonant.String(), tokens[0][1].String())
		})
	}
}
