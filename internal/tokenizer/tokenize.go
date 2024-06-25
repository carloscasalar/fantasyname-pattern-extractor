package tokenizer

import (
	"errors"
	"strings"
)

func Tokenize(value string) ([][]Token, error) {
	value = strings.ToLower(value)
	if isConsonant(value[0]) {
		return [][]Token{{TokenInitialConsonant}}, nil
	}
	return nil, errors.New("case not implemented yet")
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
