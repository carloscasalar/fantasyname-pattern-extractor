package tokenizer

type Token int

const (
	NoToken Token = iota
	TokenInitialConsonant
	TokenMiddleConsonant
	TokenVowel
	TokenApostrophe
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
		NoToken:               "NoToken",
		TokenInitialConsonant: "InitialConsonant",
		TokenMiddleConsonant:  "MiddleConsonant",
		TokenVowel:            "Vowel",
		TokenApostrophe:       "Apostrophe",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}
