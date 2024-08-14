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
	testCases := []struct {
		input         string
		expectedToken tokenizer.Token
	}{
		{"A", tokenizer.TokenVowelStrong},
		{"E", tokenizer.TokenVowelStrong},
		{"I", tokenizer.TokenVowelWeak},
		{"O", tokenizer.TokenVowelStrong},
		{"U", tokenizer.TokenVowelWeak},
		{"a", tokenizer.TokenVowelStrong},
		{"e", tokenizer.TokenVowelStrong},
		{"i", tokenizer.TokenVowelWeak},
		{"o", tokenizer.TokenVowelStrong},
		{"u", tokenizer.TokenVowelWeak},

		{"Á", tokenizer.TokenVowelStrongAcuteAccented},
		{"É", tokenizer.TokenVowelStrongAcuteAccented},
		{"Í", tokenizer.TokenVowelWeakAcuteAccented},
		{"Ó", tokenizer.TokenVowelStrongAcuteAccented},
		{"Ú", tokenizer.TokenVowelWeakAcuteAccented},
		{"á", tokenizer.TokenVowelStrongAcuteAccented},
		{"é", tokenizer.TokenVowelStrongAcuteAccented},
		{"í", tokenizer.TokenVowelWeakAcuteAccented},
		{"ó", tokenizer.TokenVowelStrongAcuteAccented},
		{"ú", tokenizer.TokenVowelWeakAcuteAccented},

		{"À", tokenizer.TokenVowelStrongGraveAccented},
		{"È", tokenizer.TokenVowelStrongGraveAccented},
		{"Ì", tokenizer.TokenVowelWeakGraveAccented},
		{"Ò", tokenizer.TokenVowelStrongGraveAccented},
		{"Ù", tokenizer.TokenVowelWeakGraveAccented},
		{"à", tokenizer.TokenVowelStrongGraveAccented},
		{"è", tokenizer.TokenVowelStrongGraveAccented},
		{"ì", tokenizer.TokenVowelWeakGraveAccented},
		{"ò", tokenizer.TokenVowelStrongGraveAccented},
		{"ù", tokenizer.TokenVowelWeakGraveAccented},

		{"Â", tokenizer.TokenVowelStrongCircumflexAccented},
		{"Ê", tokenizer.TokenVowelStrongCircumflexAccented},
		{"Î", tokenizer.TokenVowelWeakCircumflexAccented},
		{"Ô", tokenizer.TokenVowelStrongCircumflexAccented},
		{"Û", tokenizer.TokenVowelWeakCircumflexAccented},
		{"â", tokenizer.TokenVowelStrongCircumflexAccented},
		{"ê", tokenizer.TokenVowelStrongCircumflexAccented},
		{"î", tokenizer.TokenVowelWeakCircumflexAccented},
		{"ô", tokenizer.TokenVowelStrongCircumflexAccented},
		{"û", tokenizer.TokenVowelWeakCircumflexAccented},

		{"Ä", tokenizer.TokenVowelStrongDieresisAccented},
		{"Ë", tokenizer.TokenVowelStrongDieresisAccented},
		{"Ï", tokenizer.TokenVowelWeakDieresisAccented},
		{"Ö", tokenizer.TokenVowelStrongDieresisAccented},
		{"Ü", tokenizer.TokenVowelWeakDieresisAccented},
		{"ä", tokenizer.TokenVowelStrongDieresisAccented},
		{"ë", tokenizer.TokenVowelStrongDieresisAccented},
		{"ï", tokenizer.TokenVowelWeakDieresisAccented},
		{"ö", tokenizer.TokenVowelStrongDieresisAccented},
		{"ü", tokenizer.TokenVowelWeakDieresisAccented},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be ", tc.input), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(tc.input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 1, "should have only one token")
			assert.Equal(t, tc.expectedToken.String(), tokens[0].String())
		})
	}
}

func Test_given_string_starting_with_a_consonant_and_a_vowel_like(t *testing.T) {
	testCases := []struct {
		input               string
		expectedFirstToken  tokenizer.Token
		expectedSecondToken tokenizer.Token
	}{
		{"Ba", tokenizer.TokenConsonant, tokenizer.TokenVowelStrong},
		{"Pe", tokenizer.TokenConsonant, tokenizer.TokenVowelStrong},
		{"Wi", tokenizer.TokenConsonant, tokenizer.TokenVowelWeak},
		{"Qo", tokenizer.TokenConsonant, tokenizer.TokenVowelStrong},
		{"Zu", tokenizer.TokenConsonant, tokenizer.TokenVowelWeak},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be %v and second token %v", tc.input, tc.expectedFirstToken, tc.expectedSecondToken), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(tc.input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tc.expectedFirstToken.String(), tokens[0].String())
			assert.Equal(t, tc.expectedSecondToken.String(), tokens[1].String())
		})
	}
}

func Test_given_string_starting_with_a_vowel_and_a_consonant_like(t *testing.T) {
	testCases := []struct {
		input               string
		expectedFirstToken  tokenizer.Token
		expectedSecondToken tokenizer.Token
	}{
		{"Ab", tokenizer.TokenVowelStrong, tokenizer.TokenConsonant},
		{"Ep", tokenizer.TokenVowelStrong, tokenizer.TokenConsonant},
		{"Iw", tokenizer.TokenVowelWeak, tokenizer.TokenConsonant},
		{"Oq", tokenizer.TokenVowelStrong, tokenizer.TokenConsonant},
		{"Uz", tokenizer.TokenVowelWeak, tokenizer.TokenConsonant},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v, first token should be %v and the second token %v", tc.input, tc.expectedFirstToken, tc.expectedSecondToken), func(t *testing.T) {
			tokenChain, err := tokenizer.Tokenize(tc.input)

			require.NoError(t, err)
			tokens := tokenChain.Tokens()
			require.Len(t, tokens, 2, "should have two tokens but has %v", len(tokenChain.Tokens()))
			assert.Equal(t, tc.expectedFirstToken.String(), tokens[0].String())
			assert.Equal(t, tc.expectedSecondToken.String(), tokens[1].String())
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
	assert.Equal(t, tokenizer.TokenVowelStrong.String(), tokens[0].String())
	assert.Equal(t, tokenizer.TokenApostrophe.String(), tokens[1].String())
	assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[2].String())
}

func Test_acute_accented_vowel_should_be_translated_acute_accented_vowel_token_also_in_the_middle_of_a_string(t *testing.T) {
	tokenChain, err := tokenizer.Tokenize("AáB")

	require.NoError(t, err)
	tokens := tokenChain.Tokens()
	require.Len(t, tokens, 3, "should have three tokens but has %v", len(tokenChain.Tokens()))
	assert.Equal(t, tokenizer.TokenVowelStrong.String(), tokens[0].String())
	assert.Equal(t, tokenizer.TokenVowelStrongAcuteAccented.String(), tokens[1].String())
	assert.Equal(t, tokenizer.TokenConsonant.String(), tokens[2].String())
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
