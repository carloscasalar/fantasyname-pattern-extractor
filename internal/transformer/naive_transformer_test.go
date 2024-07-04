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
		"vowel should translate to !v": {
			tokenizer.NewTokenChain(tokenizer.TokenVowel),
			"!v",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			naiveTransformer := transformer.NewNaiveTransformer()

			pattern := naiveTransformer.Transform(*tc.tokenChain)

			assert.Equal(t, tc.expectedPattern, pattern)
		})
	}
}
