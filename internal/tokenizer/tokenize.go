package tokenizer

import (
	"strings"
)

func Tokenize(value string) ([]TokenChain, error) {
	value = strings.ToLower(value)
	var chain TokenChain
	chain = NewEmptyTokenChain()

	for i := 0; i < len(value); i++ {
		var err error
		char := value[i]
		chain, err = appendToken(chain, char)
		if err != nil {
			return nil, err
		}
	}

	return []TokenChain{chain}, nil
}

func appendToken(contextChain TokenChain, nextChar uint8) (TokenChain, error) {
	if isVowel(nextChar) {
		return contextChain.Add(TokenVowel), nil
	}
	if isConsonant(nextChar) {
		if contextChain.IsEmpty() {
			return contextChain.Add(TokenInitialConsonant), nil
		}
		return contextChain.Add(TokenMiddleConsonant), nil
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
