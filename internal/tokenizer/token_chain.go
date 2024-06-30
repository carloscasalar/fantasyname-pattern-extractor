package tokenizer

func NewEmptyTokenChain() *TokenChain {
	return &TokenChain{}
}

type TokenChain struct {
	tokens []Token
}

func (s TokenChain) Tokens() []Token {
	return s.tokens
}

func (s TokenChain) IsEmpty() bool {
	return false
}

func (s TokenChain) AddChar(char uint8) (*TokenChain, error) {
	if isVowel(char) {
		return s.add(TokenVowel), nil
	}

	if isAcuteAccentedVowel(char) {
		return s.add(TokenVowelAcuteAccented), nil
	}

	if isConsonant(char) {
		return s.add(TokenConsonant), nil
	}

	if token, hasTokenTranslation := symbolToken(char); hasTokenTranslation {
		return s.add(token), nil
	}

	return &s, nil
}

func (s TokenChain) add(token Token) *TokenChain {
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

func isAcuteAccentedVowel(char uint8) bool {
	return char == "á"[0] ||
		char == "é"[0] ||
		char == "í"[0] ||
		char == "ó"[0] ||
		char == "ú"[0]
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

func symbolToken(char uint8) (Token, bool) {
	symbolTokens := map[uint8]Token{
		'\'': TokenApostrophe,
		'-':  TokenHyphen,
	}

	token, hasToken := symbolTokens[char]
	return token, hasToken
}
