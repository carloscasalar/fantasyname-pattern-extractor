package transformer_test

import (
	"fmt"
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
	"github.com/stretchr/testify/assert"
)

func TestVowelProximityTransformer_returns_a_pattern_for_each_token(t *testing.T) {
	testCases := map[string]struct {
		tokenChain      tokenizer.Token
		expectedPattern string
	}{
		"strong vowel should translate to any vowel 1/3 or strong vowel 2/3": {
			tokenizer.TokenVowelStrong,
			"(<v>|(a|e|o)|(a|e|o))",
		},
		"weak vowel should translate to any vowel 1/3 or weak vowel 2/3": {
			tokenizer.TokenVowelWeak,
			"(<v>|(i|u)|(i|u))",
		},
		"strong vowel acute accented should translate to any vowel 1/3 or strong acute accented vowel 2/3": {
			tokenizer.TokenVowelStrongAcuteAccented,
			"(<v>|(á|é|ó)|(á|é|ó))",
		},
		"weak vowel acute accented should translate to any vowel 1/3 or weak acute accented vowel 2/3": {
			tokenizer.TokenVowelWeakAcuteAccented,
			"(<v>|(í|ú)|(í|ú))",
		},
		"strong vowel grave accented should translate to any vowel 1/3 or to any strong grave accented vowel": {
			tokenizer.TokenVowelStrongGraveAccented,
			"(<v>|(à|è|ò)|(à|è|ò))",
		},
		"weak vowel grave accented should translate to any vowel 1/3 or to any weak grave accented vowel": {
			tokenizer.TokenVowelWeakGraveAccented,
			"(<v>|(ì|ù)|(ì|ù))",
		},
		"strong vowel circumflex accented should translate to any vowel 1/3 or to any strong circumflex accented vowel": {
			tokenizer.TokenVowelStrongCircumflexAccented,
			"(<v>|(â|ê|ô)|(â|ê|ô))",
		},
		"weak vowel circumflex accented should translate to any vowel 1/3 or to any weak circumflex accented vowel": {
			tokenizer.TokenVowelWeakCircumflexAccented,
			"(<v>|(î|û)|(î|û))",
		},
		"strong vowel dieresis accented should translate to any vowel 1/3 or to any strong dieresis accented vowel": {
			tokenizer.TokenVowelStrongDieresisAccented,
			"(<v>|(ä|ë|ö)|(ä|ë|ö))",
		},
		"weak vowel dieresis accented should translate to any vowel 1/3 or to any weak dieresis accented vowel": {
			tokenizer.TokenVowelWeakDieresisAccented,
			"(<v>|(ï|ü)|(ï|ü))",
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
			naiveTransformer := transformer.NewVowelProximityTransformer()

			pattern := naiveTransformer.Transform(*tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern.String())
		})
	}
}

func TestVowelProximityTransformer_translates_several_tokens_to_a_sequence_of_naive_patterns(t *testing.T) {
	testCases := []struct {
		tokens          []tokenizer.Token
		expectedPattern string
	}{
		{
			[]tokenizer.Token{tokenizer.TokenVowelStrong, tokenizer.TokenConsonant},
			"(<v>|(a|e|o)|(a|e|o))c",
		},
		{
			[]tokenizer.Token{tokenizer.TokenVowelWeak, tokenizer.TokenConsonant},
			"(<v>|(i|u)|(i|u))c",
		},
		{
			[]tokenizer.Token{tokenizer.TokenCedilla, tokenizer.TokenVowelStrong, tokenizer.TokenApostrophe, tokenizer.TokenVowelWeakAcuteAccented},
			"(<c>|ç)(<v>|(a|e|o)|(a|e|o))('|)(<v>|(í|ú)|(í|ú))",
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v should translate to %v", commaSeparatedStr(tc.tokens), tc.expectedPattern), func(t *testing.T) {
			naiveTransformer := transformer.NewVowelProximityTransformer()
			tokenChain := tokenizer.NewTokenChain(tc.tokens...)

			pattern := naiveTransformer.Transform(*tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern.String())
		})
	}
}

func commaSeparatedStr(tokens []tokenizer.Token) string {
	commaSeparatedTokens := ""
	for i, token := range tokens {
		commaSeparatedTokens += token.String()
		if i < len(tokens)-1 {
			commaSeparatedTokens += ", "
		}
	}

	return commaSeparatedTokens
}
