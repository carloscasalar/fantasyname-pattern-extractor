package tokenizer

type Token struct {
	root tokenRoot
}

type tokenRoot int

const (
	rootConsonant tokenRoot = iota
	rootTildeN
	rootCedilla
	rootVowel
	rootVowelAcuteAccented
	rootVowelGraveAccented
	rootVowelCircumflexAccented
	rootVowelDieresisAccented
	rootApostrophe
	rootHyphen
)

func (t tokenRoot) String() string {
	tokenStrings := map[tokenRoot]string{
		rootConsonant:               "Consonant",
		rootTildeN:                  "TildeN",
		rootCedilla:                 "Cedilla",
		rootVowel:                   "Vowel",
		rootVowelAcuteAccented:      "VowelAcuteAccented",
		rootVowelGraveAccented:      "VowelGraveAccented",
		rootVowelCircumflexAccented: "VowelCircumflexAccented",
		rootVowelDieresisAccented:   "VowelDieresisAccented",
		rootApostrophe:              "Apostrophe",
		rootHyphen:                  "Hyphen",
	}

	if str, found := tokenStrings[t]; found {
		return str
	}
	return "UndefinedToken"
}

var (
	TokenConsonant               = Token{root: rootConsonant}
	TokenTildeN                  = Token{root: rootTildeN}
	TokenCedilla                 = Token{root: rootCedilla}
	TokenVowel                   = Token{root: rootVowel}
	TokenVowelAcuteAccented      = Token{root: rootVowelAcuteAccented}
	TokenVowelGraveAccented      = Token{root: rootVowelGraveAccented}
	TokenVowelCircumflexAccented = Token{root: rootVowelCircumflexAccented}
	TokenVowelDieresisAccented   = Token{root: rootVowelDieresisAccented}
	TokenApostrophe              = Token{root: rootApostrophe}
	TokenHyphen                  = Token{root: rootHyphen}
)

func (t Token) String() string {
	tokenStrings := map[Token]string{
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
