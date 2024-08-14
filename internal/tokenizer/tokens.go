package tokenizer

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

type Token struct {
	root   tokenRoot
	accent accent
}

func (t Token) IsVowel() bool {
	return t.root == rootVowel
}

func (t Token) IsAcuteAccented() bool {
	return t.accent == accentAcute
}

func (t Token) IsGraveAccented() bool {
	return t.accent == accentGrave
}

func (t Token) IsCircumflexAccented() bool {
	return t.accent == accentCircumflex
}

func (t Token) IsDieresisAccented() bool {
	return t.accent == accentDieresis
}

func (t Token) IsConsonant() bool {
	return t.root == rootConsonant || t.root == rootTildeN || t.root == rootCedilla
}

func (t Token) IsTildeN() bool {
	return t.root == rootTildeN
}

func (t Token) IsCedilla() bool {
	return t.root == rootCedilla
}

func (t Token) IsApostrophe() bool {
	return t.root == rootApostrophe
}

func (t Token) IsHyphen() bool {
	return t.root == rootHyphen
}

func (t Token) String() string {
	return t.root.String() + t.accent.String()
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
