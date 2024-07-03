package tokenizer

func Tokenize(value string) (*TokenChain, error) {
	chain := NewEmptyTokenChain()

	for _, r := range value {
		var err error
		chain, err = chain.AddChar(r)
		if err != nil {
			return nil, err
		}
	}

	return chain, nil
}
