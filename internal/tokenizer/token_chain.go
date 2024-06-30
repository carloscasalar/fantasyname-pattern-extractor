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

func (e EmptyTokenChain) add(token Token) TokenChain {
	return &SingleTokenChain{token: token}
}

func (e EmptyTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return e.add(TokenVowel), nil
	}
	if isConsonant(char) {
		return e.add(TokenInitialConsonant), nil
	}
	return e, nil
}

func (e EmptyTokenChain) IsEmpty() bool {
	return true
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

func (s SingleTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return s.add(TokenVowel), nil
	}
	if isConsonant(char) {
		return s.add(TokenMiddleConsonant), nil
	}

	return s, nil
}

func (s SingleTokenChain) add(token Token) TokenChain {
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

func (s SeveralTokenChain) AddChar(char uint8) (TokenChain, error) {
	if isVowel(char) {
		return s.add(TokenVowel), nil
	}
	if isConsonant(char) {
		return s.add(TokenMiddleConsonant), nil
	}
	return s, nil
}

func (s SeveralTokenChain) add(token Token) TokenChain {
	s.tokens = append(s.tokens, token)
	return &s
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
