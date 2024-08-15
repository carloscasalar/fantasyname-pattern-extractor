package commands_test

import (
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/commands"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractPatterns_should_return_the_transformed_tokens_patter_uppercased(t *testing.T) {
	const firstTransformerGeneratedPattern = "abc"
	const capitalizedFirstTransformerGeneratedPattern = "!abc"
	firstTokenSpy := new(tokenSpy)
	firstTransformerMock := newMockTransformerReturningPattern(firstTransformerGeneratedPattern, firstTokenSpy)
	secondTokenSpy := new(tokenSpy)
	const secondTransformerGeneratedPattern = "def"
	const capitalizedSecondTransformerGeneratedPattern = "!def"
	secondTransformerMock := newMockTransformerReturningPattern(secondTransformerGeneratedPattern, secondTokenSpy)
	extractPatterns := commands.NewExtractPatterns(firstTransformerMock, secondTransformerMock)

	resultingPatterns, err := extractPatterns.Execute("some name")

	require.NoError(t, err)
	require.Len(t, resultingPatterns, 2)
	assert.Equal(t, capitalizedFirstTransformerGeneratedPattern, resultingPatterns[0])
	assert.Equal(t, tokensFor("some name"), firstTokenSpy.tokensTransformed, "expected first transformer had transformed the tokens using the 'some name' string")
	assert.Equal(t, capitalizedSecondTransformerGeneratedPattern, resultingPatterns[1])
	assert.Equal(t, tokensFor("some name"), secondTokenSpy.tokensTransformed, "expected second transformer had transformed the tokens using the 'some name' string")
}
