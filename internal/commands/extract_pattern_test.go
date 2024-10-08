package commands_test

import (
	"fmt"
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/commands"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/tokenizer"
	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/transformer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractPattern_should_return_the_transformed_tokens_patter_uppercased(t *testing.T) {
	const generatedPattern = "abc"
	const firstLetterUppercasedPattern = "!abc"
	tokenSpy := new(tokenSpy)
	transformerMock := newMockTransformerReturningPattern(generatedPattern, tokenSpy)
	extractPattern := commands.NewExtractPattern(transformerMock)

	resultingPattern, err := extractPattern.Execute("some name")

	require.NoError(t, err)
	assert.Equal(t, firstLetterUppercasedPattern, resultingPattern)
	assert.Equal(t, tokensFor("some name"), tokenSpy.tokensTransformed, "expected first transformer had transformed the tokens using the 'some name' string")

}

func tokensFor(value string) *tokenizer.TokenChain {
	tokens, _ := tokenizer.Tokenize(value)
	return tokens
}

type mockPattern struct {
	pattern string
}

func newMockPattern(pattern string) mockPattern {
	return mockPattern{pattern}
}

func (p mockPattern) String() string {
	return p.pattern
}

func (p mockPattern) Capitalize() string {
	return fmt.Sprintf("!%s", p.pattern)
}

type tokenSpy struct {
	tokensTransformed *tokenizer.TokenChain
}

type mockTransformer struct {
	patternToReturn mockPattern
	tokenSpy        *tokenSpy
}

func newMockTransformerReturningPattern(patternToGenerate string, tspy *tokenSpy) *mockTransformer {
	return &mockTransformer{
		patternToReturn: newMockPattern(patternToGenerate),
		tokenSpy:        tspy,
	}
}

func (m *mockTransformer) Transform(tokens tokenizer.TokenChain) transformer.Pattern {
	m.tokenSpy.tokensTransformed = &tokens
	return m.patternToReturn
}
