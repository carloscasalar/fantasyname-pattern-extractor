package tokenizer

type Token struct {
	root tokenRoot
}

type tokenOption func(*Token)

func newToken(opts ...tokenOption) Token {
	t := Token{}
	for _, opt := range opts {
		opt(&t)
	}
	return t
}

func withRoot(r tokenRoot) tokenOption {
	return func(t *Token) {
		t.root = r
	}
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
	TokenConsonant               = newToken(withRoot(rootConsonant))
	TokenTildeN                  = newToken(withRoot(rootTildeN))
	TokenCedilla                 = newToken(withRoot(rootCedilla))
	TokenVowel                   = newToken(withRoot(rootVowel))
	TokenVowelAcuteAccented      = newToken(withRoot(rootVowelAcuteAccented))
	TokenVowelGraveAccented      = newToken(withRoot(rootVowelGraveAccented))
	TokenVowelCircumflexAccented = newToken(withRoot(rootVowelCircumflexAccented))
	TokenVowelDieresisAccented   = newToken(withRoot(rootVowelDieresisAccented))
	TokenApostrophe              = newToken(withRoot(rootApostrophe))
	TokenHyphen                  = newToken(withRoot(rootHyphen))
)

func (t Token) String() string {
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

	if str, found := tokenStrings[t.root]; found {
		return str
	}
	return "UndefinedToken"
}
