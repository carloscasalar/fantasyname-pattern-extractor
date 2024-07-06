package transformer_test

import (
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
	"github.com/stretchr/testify/assert"
)

func TestNaiveTransformer_returns_a_symbol_for_each_token(t *testing.T) {
	testCases := map[string]struct {
		tokenChain      tokenizer.Token
		expectedPattern string
	}{
		"vowel should translate to v": {
			tokenizer.TokenVowel,
			"v",
		},
		"vowel should acute accented to any acute accented vowel": {
			tokenizer.TokenVowelAcuteAccented,
			"(<v>|á|é|í|ó|ú)",
		},
		"vowel should grave accented to any acute accented vowel": {
			tokenizer.TokenVowelGraveAccented,
			"(<v>|à|è|ì|ò|ù)",
		},
		"vowel should circumflex accented to any acute accented vowel": {
			tokenizer.TokenVowelCircumflexAccented,
			"(<v>|â|ê|î|ô|û)",
		},
		"vowel should  dieresis accented to any acute accented vowel": {
			tokenizer.TokenVowelDieresisAccented,
			"(<v>|ä|ë|ï|ö|ü)",
		},
		"constant at the beginning should translate to c": {
			tokenizer.TokenConsonant,
			"c",
		},
		"tilde n at the beginning should translate to (<c>|ñ)": {
			tokenizer.TokenTildeN,
			"(<c>|ñ)",
		},
		"cedilla at the beginning should translate to (<c>|ç)": {
			tokenizer.TokenCedilla,
			"(<c>|ç)",
		},
		"apostrophe should translate to apostrophe or nothing('|)": {
			tokenizer.TokenApostrophe,
			"('|)",
		},
		"hyphen should translate to hyphen or nothing(-|)": {
			tokenizer.TokenHyphen,
			"(-|)",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tokenChain := tokenizer.NewTokenChain(tc.tokenChain)
			naiveTransformer := transformer.NewNaiveTransformer()

			pattern := naiveTransformer.Transform(*tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern.String())
		})
	}
}
