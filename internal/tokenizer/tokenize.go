package tokenizer

import "errors"

func Tokenize(value string) ([][]Token, error) {
	// if the string starts with a consonant, the first token is InitialConsonant
	if isConsonant(value[0]) {
		return [][]Token{{TokenInitialConsonant}}, nil
	}
	return nil, errors.New("case not implemented yet")
}

func isConsonant(value uint8) bool {
	return value == 'B' ||
		value == 'C' ||
		value == 'D' ||
		value == 'F' ||
		value == 'G' ||
		value == 'H' ||
		value == 'J' ||
		value == 'K' ||
		value == 'L' ||
		value == 'M' ||
		value == 'N' ||
		value == 'P' ||
		value == 'Q' ||
		value == 'R' ||
		value == 'S' ||
		value == 'T' ||
		value == 'V' ||
		value == 'W' ||
		value == 'X' ||
		value == 'Y' ||
		value == 'Z'
}
