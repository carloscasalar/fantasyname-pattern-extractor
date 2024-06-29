package tokenizer

type TokenChain interface {
	Tokens() []Token
	IsEmpty() bool
	AddChar(uint8) (TokenChain, error)
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

func (e EmptyTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return e.Add(TokenVowel), nil
	}
	if isConsonant(char) {
		return e.Add(TokenInitialConsonant), nil
	}
	return e, nil
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

func (s SingleTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return s.Add(TokenVowel), nil
	}
	if isConsonant(char) {
		return s.Add(TokenMiddleConsonant), nil
	}

	return s, nil
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

func (s SeveralTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return s.Add(TokenVowel), nil
	}
	if isConsonant(char) {
		return s.Add(TokenMiddleConsonant), nil
	}
	return s, nil
}

func isVowel(value uint8) bool {
	return value == 'a' ||
		value == 'e' ||
		value == 'i' ||
		value == 'o' ||
		value == 'u'
}

func isConsonant(value uint8) bool {
	return value == 'b' ||
		value == 'c' ||
		value == 'd' ||
		value == 'f' ||
		value == 'g' ||
		value == 'h' ||
		value == 'j' ||
		value == 'k' ||
		value == 'l' ||
		value == 'm' ||
		value == 'n' ||
		value == 'p' ||
		value == 'q' ||
		value == 'r' ||
		value == 's' ||
		value == 't' ||
		value == 'v' ||
		value == 'w' ||
		value == 'x' ||
		value == 'y' ||
		value == 'z'
}
