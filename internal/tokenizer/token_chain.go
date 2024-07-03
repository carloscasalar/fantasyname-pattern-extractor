package tokenizer

import "unicode"

type TokenChain struct {
	tokens []Token
}

func NewEmptyTokenChain() *TokenChain {
	return &TokenChain{}
}

func (s TokenChain) Tokens() []Token {
	return s.tokens
}

func (s TokenChain) AddChar(r rune) (*TokenChain, error) {
	if isPlainVowel(r) {
		return s.add(TokenVowel), nil
	}

	if isAcuteVowel(r) {
		return s.add(TokenVowelAcuteAccented), nil
	}

	if isGraveVowel(r) {
		return s.add(TokenVowelGraveAccented), nil
	}

	if isCircumflexVowel(r) {
		return s.add(TokenVowelCircumflexAccented), nil
	}

	if isDiaeresisVowel(r) {
		return s.add(TokenVowelDieresisAccented), nil
	}

	if isTildeN(r) {
		return s.add(TokenTildeN), nil
	}

	if isConsonant(r) {
		return s.add(TokenConsonant), nil
	}

	if token, hasTokenTranslation := symbolToken(r); hasTokenTranslation {
		return s.add(token), nil
	}

	return &s, nil
}

func (s TokenChain) add(token Token) *TokenChain {
	s.tokens = append(s.tokens, token)
	return &s
}

func isConsonant(r rune) bool {
	return unicode.IsLetter(r) && !isVowel(r)
}

func isTildeN(r rune) bool {
	return r == 'ñ' || r == 'Ñ'
}

func isVowel(r rune) bool {
	return isPlainVowel(r) || isAcuteVowel(r) || isGraveVowel(r) || isDiaeresisVowel(r) ||
		isCircumflexVowel(r) || isDieresisVowel(r)
}

func isPlainVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func isAcuteVowel(r rune) bool {
	return r == 'á' || r == 'é' || r == 'í' || r == 'ó' || r == 'ú' ||
		r == 'Á' || r == 'É' || r == 'Í' || r == 'Ó' || r == 'Ú'
}

func isGraveVowel(r rune) bool {
	return r == 'à' || r == 'è' || r == 'ì' || r == 'ò' || r == 'ù' ||
		r == 'À' || r == 'È' || r == 'Ì' || r == 'Ò' || r == 'Ù'
}

func isCircumflexVowel(r rune) bool {
	return r == 'â' || r == 'ê' || r == 'î' || r == 'ô' || r == 'û' ||
		r == 'Â' || r == 'Ê' || r == 'Î' || r == 'Ô' || r == 'Û'
}

func isDiaeresisVowel(r rune) bool {
	return r == 'ä' || r == 'ë' || r == 'ï' || r == 'ö' || r == 'ü' ||
		r == 'Ä' || r == 'Ë' || r == 'Ï' || r == 'Ö' || r == 'Ü'
}

func isDieresisVowel(r rune) bool {
	return r == 'ä' || r == 'ë' || r == 'ï' || r == 'ö' || r == 'ü' ||
		r == 'Ä' || r == 'Ë' || r == 'Ï' || r == 'Ö' || r == 'Ü'
}

func symbolToken(r rune) (Token, bool) {
	symbolTokens := map[rune]Token{
		'\'': TokenApostrophe,
		'-':  TokenHyphen,
	}

	token, hasToken := symbolTokens[r]
	return token, hasToken
}
