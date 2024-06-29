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
		chain, err = chain.AddChar(char)
		if err != nil {
			return nil, err
		}
	}

	return []TokenChain{chain}, nil
}
