package tokenizer

type Token int

const (
	NoToken Token = iota
	TokenConsonant
	TokenVowel
	TokenApostrophe
	TokenHyphen
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
		NoToken:         "NoToken",
		TokenConsonant:  "Consonant",
		TokenVowel:      "Vowel",
		TokenApostrophe: "Apostrophe",
		TokenHyphen:     "Hyphen",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}
