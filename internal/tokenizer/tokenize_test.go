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
		t.Run(fmt.Sprintf("%v, first token should be consonant", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have only one token")
			assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[0].String())
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
			assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0].String())
		})
	}
}

func Test_given_string_starting_with_a_consonant_and_a_vowel_like(t *testing.T) {
	testCases := []string{
		"Ba", "Pe", "Wi", "Qo", "Zu",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be consonant and second token a vowel", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[0].String())
			assert.Equal(t, tokenizer.TokenVowel.String(), tokens[1].String())
		})
	}
}

func Test_given_string_starting_with_a_vowel_and_a_consonant_like(t *testing.T) {
	testCases := []string{
		"Ab", "Ep", "Iw", "Oq", "Uz",
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be a vowel and the second token a consonant", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0].String())
			assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[1].String())
		})
	}
}

func Test_apostrophe_should_be_translated_to_apostrophe_token(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("'")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenApostrophe.String(), tokens[0].String())
}

func Test_hyphen_should_be_translated_to_hyphen_token(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("-")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenHyphen.String(), tokens[0].String())
}

func Test_apostrophe_should_be_translated_to_apostrophe_token_also_in_the_middle_of_a_string(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("A'B")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 3, "should have three tokens but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0].String())
	assert.Equal(t, tokenizer.TokenApostrophe.String(), tokens[1].String())
	assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[2].String())
}

func Test_acute_accented_vowel_should_be_translated_acute_accented_vowel_token(t *testing.T) {
	testVowels := []string{"á", "é", "í", "ó", "ú", "Á", "É", "Í", "Ó", "Ú"}

	for _, input := range testVowels {
		t.Run(fmt.Sprintf("%v, should be translated to acute accented vowel token", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowelAcuteAccented.String(), tokens[0].String())
		})
	}
}

func Test_acute_accented_vowel_should_be_translated_acute_accented_vowel_token_also_in_the_middle_of_a_string(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("AáB")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 3, "should have three tokens but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenVowel.String(), tokens[0].String())
	assert.Equal(t, tokenizer.TokenVowelAcuteAccented.String(), tokens[1].String())
	assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[2].String())
}

func Test_grave_accented_vowel_should_be_translated_grave_accented_vowel_token(t *testing.T) {
	testVowels := []string{"à", "è", "ì", "ò", "ù", "À", "È", "Ì", "Ò", "Ù"}

	for _, input := range testVowels {
		t.Run(fmt.Sprintf("%v, should be translated to grave accented vowel token", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowelGraveAccented.String(), tokens[0].String())
		})
	}
}

func Test_circumflex_accented_vowel_should_be_translated_grave_accented_vowel_token(t *testing.T) {
	testVowels := []string{"â", "ê", "î", "ô", "û", "Â", "Ê", "Î", "Ô", "Û"}

	for _, input := range testVowels {
		t.Run(fmt.Sprintf("%v, should be translated to circumflex accented vowel token", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowelCircumflexAccented.String(), tokens[0].String())
		})
	}
}

func Test_vowel_with_dieresis_should_be_translated_to_dieresis_vowel_token(t *testing.T) {
	testVowels := []string{"ä", "ë", "ï", "ö", "ü", "Ä", "Ë", "Ï", "Ö", "Ü"}

	for _, input := range testVowels {
		t.Run(fmt.Sprintf("%v, should be translated to dieresis vowel token", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenVowelDieresisAccented.String(), tokens[0].String())
		})
	}
}

func Test_tilde_with_n_on_top_should_be_translated_to_tilde_n_token(t *testing.T) {
	testCases := []string{"ñ", "Ñ"}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, should be translated to consonant", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenTildeN.String(), tokens[0].String())
		})
	}
}

func Test_cedilla_should_be_translated_to_cedilla_token(t *testing.T) {
	testCases := []string{"ç", "Ç"}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("%v, should be translated to cedilla token", input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have one token but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tokenizer.TokenCedilla.String(), tokens[0].String())
		})
	}
}
