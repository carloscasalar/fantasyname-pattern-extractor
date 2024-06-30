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
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have only one token")
			assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0])
		})
	}
}

func Test_given_string_starting_with_the_vowel(t *testing.T) {
	testCases := []string{
		"A", "E", "I", "O", "U", "a", "e", "i", "o", "u",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be single vowel", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have only one token")
			assert.Equal(t, tokenizer.TokenVowel, tokens[0])
		})
	}
}

func Test_given_string_starting_with_a_consonant_and_a_vowel_like(t *testing.T) {
	testCases := []string{
		"Ba", "Pe", "Wi", "Qo", "Zu",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be starting consonant and second token a vowel", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0])
			assert.Equal(t, tokenizer.TokenVowel, tokens[1])
		})
	}
}

func Test_given_string_starting_with_a_vowel_and_a_consonant_like(t *testing.T) {
	testCases := []string{
		"Ab", "Ep", "Iw", "Oq", "Uz",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be a vowel and the second token a middle consonant", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0].String())
			assert.Equal(t, tokenizer.TokenMiddleConsonant.String(), tokens[1].String())
		})
	}
}

func Test_given_string_starting_with_two_consonants_first_token_should_be_starting_consonant_and_second_middle_consonant(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("Kr")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0])
	assert.Equal(t, tokenizer.TokenMiddleConsonant, tokens[1])
}

func Test_apostrophe_should_be_translated_to_apostrophe_token(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("'")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenApostrophe, tokens[0])
}

func Test_apostrophe_should_be_translated_to_apostrophe_token_also_in_the_middle_of_a_string(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("A'B")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 3, "should have three tokens but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenVowel, tokens[0])
	assert.Equal(t, tokenizer.TokenApostrophe, tokens[1])
	assert.Equal(t, tokenizer.TokenMiddleConsonant, tokens[2])
}

func Test_hyphen_should_be_translated_to_hyphen_token(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("-")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenHyphen, tokens[0])
}
