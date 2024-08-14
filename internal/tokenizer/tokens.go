package tokenizer

type Token struct {
	root   tokenRoot
	accent accent
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

func withAccent(a accent) tokenOption {
	return func(t *Token) {
		t.accent = a
	}
}

type tokenRoot int

const (
	rootUndefined tokenRoot = iota
	rootConsonant
	rootTildeN
	rootCedilla
	rootVowel
	rootApostrophe
	rootHyphen
)

func (t tokenRoot) String() string {
	tokenStrings := map[tokenRoot]string{
		rootUndefined:  "Undefined",
		rootConsonant:  "Consonant",
		rootTildeN:     "TildeN",
		rootCedilla:    "Cedilla",
		rootVowel:      "Vowel",
		rootApostrophe: "Apostrophe",
		rootHyphen:     "Hyphen",
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
	TokenVowelAcuteAccented      = newToken(withRoot(rootVowel), withAccent(accentAcute))
	TokenVowelGraveAccented      = newToken(withRoot(rootVowel), withAccent(accentGrave))
	TokenVowelCircumflexAccented = newToken(withRoot(rootVowel), withAccent(accentCircumflex))
	TokenVowelDieresisAccented   = newToken(withRoot(rootVowel), withAccent(accentDieresis))
	TokenApostrophe              = newToken(withRoot(rootApostrophe))
	TokenHyphen                  = newToken(withRoot(rootHyphen))
)

func (t Token) String() string {
	tokenStrings := map[tokenRoot]string{
		rootConsonant:  "Consonant",
		rootTildeN:     "TildeN",
		rootCedilla:    "Cedilla",
		rootVowel:      "Vowel",
		rootApostrophe: "Apostrophe",
		rootHyphen:     "Hyphen",
	}

	rootString, found := tokenStrings[t.root]
	if !found {
		return "UndefinedToken"
	}
	return rootString + t.accent.String()
}

type accent int

const (
	accentNone accent = iota
	accentAcute
	accentGrave
	accentCircumflex
	accentDieresis
)

func (a accent) String() string {
	accentStrings := map[accent]string{
		accentNone:       "",
		accentAcute:      "AcuteAccented",
		accentGrave:      "GraveAccented",
		accentCircumflex: "CircumflexAccented",
		accentDieresis:   "DieresisAccented",
	}

	if str, found := accentStrings[a]; found {
		return str
	}
	return "UndefinedAccent"
}
