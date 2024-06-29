package tokenizer

import (
	"strings"
)

func Tokenize(value string) ([][]Token, error) {
	value = strings.ToLower(value)
	tokenChain := make([]Token, 0)

	for i := 0; i < len(value); i++ {
		var err error
		char := value[i]
		tokenChain, err = extractToken(tokenChain, char)
		if err != nil {
			return nil, err
		}
	}

	return [][]Token{tokenChain}, nil
}

func extractToken(contextChain []Token, nextChar uint8) ([]Token, error) {
	if isVowel(nextChar) {
		return append(contextChain, TokenVowel), nil
	}
	if isConsonant(nextChar) {
		if len(contextChain) == 0 {
			return append(contextChain, TokenInitialConsonant), nil
		}
		return append(contextChain, TokenMiddleConsonant), nil
	}

	return contextChain, nil
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
