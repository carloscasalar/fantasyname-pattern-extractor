package tokenizer

type TokenChain interface {
	Tokens() []Token
	IsEmpty() bool
	Add(Token) TokenChain
}

type SingleTokenChain struct {
	token Token
}

func (s SingleTokenChain) Tokens() []Token {
	return []Token{s.token}
}

func (s SingleTokenChain) IsEmpty() bool {
	return false
}

func (s SingleTokenChain) Add(token Token) TokenChain {
	return &SeveralTokenChain{tokens: []Token{s.token, token}}
}

type SeveralTokenChain struct {
	tokens []Token
}

func (s SeveralTokenChain) Tokens() []Token {
	return s.tokens
}

func (s SeveralTokenChain) IsEmpty() bool {
	return false
}

func (s SeveralTokenChain) Add(token Token) TokenChain {
	s.tokens = append(s.tokens, token)
	return &s
}

type EmptyTokenChain struct{}

func NewEmptyTokenChain() *EmptyTokenChain {
	return &EmptyTokenChain{}
}

func (e EmptyTokenChain) Tokens() []Token {
	return []Token{}
}

func (e EmptyTokenChain) IsEmpty() bool {
	return true
}

func (e EmptyTokenChain) Add(token Token) TokenChain {
	return &SingleTokenChain{token: token}
}
