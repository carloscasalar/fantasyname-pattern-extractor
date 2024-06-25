package tokenizer_test

import (
	"github.com/carloscasalar/fantasy-pattern-inferrer/internal/tokenizer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_given_string_starting_with_consonant_token_should_be_starting_consonant(t *testing.T) {
	input := "T"

	tokens, err := tokenizer.Tokenize(input)

	require.NoError(t, err)
	require.Len(t, tokens, 1)
	assert.Equal(t, tokenizer.TokenInitialConsonant, tokens[0][0])
}
