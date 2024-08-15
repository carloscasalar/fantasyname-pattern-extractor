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
		"strong vowel should translate to v": {
			tokenizer.TokenVowelStrong,
			"v",
		},
		"weak vowel should translate to v": {
			tokenizer.TokenVowelWeak,
			"v",
		},
		"strong vowel acute accented should translate to any acute accented vowel": {
			tokenizer.TokenVowelStrongAcuteAccented,
			"(<v>|(á|é|í|ó|ú))",
		},
		"weak vowel acute accented should translate to any acute accented vowel": {
			tokenizer.TokenVowelWeakAcuteAccented,
			"(<v>|(á|é|í|ó|ú))",
		},
		"strong vowel grave accented should translate to any grave accented vowel": {
			tokenizer.TokenVowelStrongGraveAccented,
			"(<v>|(à|è|ì|ò|ù))",
		},
		"weak vowel grave accented should translate to any grave accented vowel": {
			tokenizer.TokenVowelWeakGraveAccented,
			"(<v>|(à|è|ì|ò|ù))",
		},
		"strong vowel circumflex accented should translate to any circumflex accented vowel": {
			tokenizer.TokenVowelStrongCircumflexAccented,
			"(<v>|(â|ê|î|ô|û))",
		},
		"weak vowel circumflex accented should translate to any circumflex accented vowel": {
			tokenizer.TokenVowelWeakCircumflexAccented,
			"(<v>|(â|ê|î|ô|û))",
		},
		"strong vowel dieresis accented should translate to any dieresis accented vowel": {
			tokenizer.TokenVowelStrongDieresisAccented,
			"(<v>|(ä|ë|ï|ö|ü))",
		},
		"weak vowel dieresis accented should translate to any dieresis accented vowel": {
			tokenizer.TokenVowelWeakDieresisAccented,
			"(<v>|(ä|ë|ï|ö|ü))",
		},
		"consonant at the beginning should translate to c": {
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

func TestNaiveTransformer_translates_several_tokens_to_a_sequence_of_naive_patterns(t *testing.T) {
	testCases := map[string]struct {
		tokens          []tokenizer.Token
		expectedPattern string
	}{
		"strong vowel and consonant should translate to v and c": {
			[]tokenizer.Token{tokenizer.TokenVowelStrong, tokenizer.TokenConsonant},
			"vc",
		},
		"weak vowel and consonant should translate to v and c": {
			[]tokenizer.Token{tokenizer.TokenVowelWeak, tokenizer.TokenConsonant},
			"vc",
		},
		"cedilla, vowel, apostrophe and acute accented vowel should translate to (<c>|ç)v('|)|(<v>|(á|é|í|ó|ú))": {
			[]tokenizer.Token{tokenizer.TokenCedilla, tokenizer.TokenVowelStrong, tokenizer.TokenApostrophe, tokenizer.TokenVowelWeakAcuteAccented},
			"(<c>|ç)v('|)(<v>|(á|é|í|ó|ú))",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			naiveTransformer := transformer.NewNaiveTransformer()
			tokenChain := tokenizer.NewTokenChain(tc.tokens...)

			pattern := naiveTransformer.Transform(*tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern.String())
		})
	}
}
