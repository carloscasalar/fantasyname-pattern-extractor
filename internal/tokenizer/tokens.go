package tokenizer

type Token int

const (
	NoToken Token = iota
	TokenConsonant
	TokenTildeN
	TokenVowel
	TokenVowelAcuteAccented
	TokenVowelGraveAccented
	TokenVowelCircumflexAccented
	TokenApostrophe
	TokenHyphen
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
		NoToken:                      "NoToken",
		TokenConsonant:               "Consonant",
		TokenTildeN:                  "TildeN",
		TokenVowel:                   "Vowel",
		TokenVowelAcuteAccented:      "VowelAcuteAccented",
		TokenVowelGraveAccented:      "VowelGraveAccented",
		TokenVowelCircumflexAccented: "VowelCircumflexAccented",
		TokenApostrophe:              "Apostrophe",
		TokenHyphen:                  "Hyphen",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}
