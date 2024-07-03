package tokenizer

type Token int

const (
	NoToken Token = iota
	TokenConsonant
	TokenTildeN
	TokenCedilla
	TokenVowel
	TokenVowelAcuteAccented
	TokenVowelGraveAccented
	TokenVowelCircumflexAccented
	TokenVowelDieresisAccented
	TokenApostrophe
	TokenHyphen
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
		NoToken:                      "NoToken",
		TokenConsonant:               "Consonant",
		TokenTildeN:                  "TildeN",
		TokenCedilla:                 "Cedilla",
		TokenVowel:                   "Vowel",
		TokenVowelAcuteAccented:      "VowelAcuteAccented",
		TokenVowelGraveAccented:      "VowelGraveAccented",
		TokenVowelCircumflexAccented: "VowelCircumflexAccented",
		TokenVowelDieresisAccented:   "VowelDieresisAccented",
		TokenApostrophe:              "Apostrophe",
		TokenHyphen:                  "Hyphen",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}
