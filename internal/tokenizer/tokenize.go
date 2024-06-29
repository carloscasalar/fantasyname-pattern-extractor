package tokenizer

import (
	"strings"
)

func Tokenize(value string) ([][]Token, error) {
	value = strings.ToLower(value)
	tokenChain := make([]Token, len(value))
	if isConsonant(value[0]) {
		tokenChain[0] = TokenInitialConsonant
	}

	if isVowel(value[0]) {
		tokenChain[0] = TokenVowel
	}

	if len(value) > 1 && isVowel(value[1]) {
		tokenChain[1] = TokenVowel
	}

	return [][]Token{tokenChain}, nil
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
