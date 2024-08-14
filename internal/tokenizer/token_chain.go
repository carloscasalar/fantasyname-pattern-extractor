package tokenizer

import "unicode"

type TokenChain struct {
	tokens []Token
}

func NewEmptyTokenChain() *TokenChain {
	return &TokenChain{}
}

func NewTokenChain(tokens ...Token) *TokenChain {
	return &TokenChain{tokens: tokens}
}

func (s TokenChain) Tokens() []Token {
	return s.tokens
}

func (s TokenChain) AddChar(r rune) (*TokenChain, error) {
	if token, hasTokenTranslation := tokenByRune[unicode.ToLower(r)]; hasTokenTranslation {
		return s.add(token), nil
	}

	return &s, nil
}

func (s TokenChain) add(token Token) *TokenChain {
	s.tokens = append(s.tokens, token)
	return &s
}
