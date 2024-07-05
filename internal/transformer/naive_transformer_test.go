package transformer_test

import (
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
	"github.com/stretchr/testify/assert"
)

func Test_naive_transformer_returns_a_symbol_for_each_token(t *testing.T) {
	testCases := map[string]struct {
		tokenChain      *tokenizer.TokenChain
		expectedPattern string
	}{
		"vowel should translate to v": {
			tokenizer.NewTokenChain(tokenizer.TokenVowel),
			"v",
		},
		"vowel should acute accented to any acute accented vowel": {
			tokenizer.NewTokenChain(tokenizer.TokenVowelAcuteAccented),
			"(<v>|á|é|í|ó|ú)",
		},
		"vowel should grave accented to any acute accented vowel": {
			tokenizer.NewTokenChain(tokenizer.TokenVowelGraveAccented),
			"(<v>|à|è|ì|ò|ù)",
		},
		"vowel should circumflex accented to any acute accented vowel": {
			tokenizer.NewTokenChain(tokenizer.TokenVowelCircumflexAccented),
			"(<v>|â|ê|î|ô|û)",
		},
		"vowel should  dieresis accented to any acute accented vowel": {
			tokenizer.NewTokenChain(tokenizer.TokenVowelDieresisAccented),
			"(<v>|ä|ë|ï|ö|ü)",
		},
		"constant at the beginning should translate to c": {
			tokenizer.NewTokenChain(tokenizer.TokenConsonant),
			"c",
		},
		"tilde n at the beginning should translate to (<c>|ñ)": {
			tokenizer.NewTokenChain(tokenizer.TokenTildeN),
			"(<c>|ñ)",
		},
		"cedilla at the beginning should translate to (<c>|ç)": {
			tokenizer.NewTokenChain(tokenizer.TokenCedilla),
			"(<c>|ç)",
		},
		"apostrophe should translate to apostrophe or nothing('|)": {
			tokenizer.NewTokenChain(tokenizer.TokenApostrophe),
			"('|)",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			naiveTransformer := transformer.NewNaiveTransformer()

			pattern := naiveTransformer.Transform(*tc.tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern.String())
		})
	}
}
