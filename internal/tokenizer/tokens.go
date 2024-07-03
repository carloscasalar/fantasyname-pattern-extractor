package tokenizer

type Token int

const (
	NoToken Token = iota
	TokenConsonant
	TokenVowel
	TokenVowelAcuteAccented
	TokenVowelGraveAccented
	TokenApostrophe
	TokenHyphen
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
		NoToken:                 "NoToken",
		TokenConsonant:          "Consonant",
		TokenVowel:              "Vowel",
		TokenVowelAcuteAccented: "VowelAcuteAccented",
		TokenVowelGraveAccented: "VowelGraveAccented",
		TokenApostrophe:         "Apostrophe",
		TokenHyphen:             "Hyphen",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}
