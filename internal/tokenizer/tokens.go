package tokenizer

var (
	TokenConsonant                     Token = newToken(withRoot(rootConsonant))
	TokenTildeN                        Token = newToken(withRoot(rootTildeN))
	TokenCedilla                       Token = newToken(withRoot(rootCedilla))
	TokenVowelStrong                   Token = newVowelToken(withStrength(strengthStrong))
	TokenVowelWeak                     Token = newVowelToken(withStrength(strengthWeak))
	TokenVowelStrongAcuteAccented      Token = newVowelToken(withAccent(accentAcute), withStrength(strengthStrong))
	TokenVowelWeakAcuteAccented        Token = newVowelToken(withAccent(accentAcute), withStrength(strengthWeak))
	TokenVowelStrongGraveAccented      Token = newVowelToken(withAccent(accentGrave), withStrength(strengthStrong))
	TokenVowelWeakGraveAccented        Token = newVowelToken(withAccent(accentGrave), withStrength(strengthWeak))
	TokenVowelStrongCircumflexAccented Token = newVowelToken(withAccent(accentCircumflex), withStrength(strengthStrong))
	TokenVowelWeakCircumflexAccented   Token = newVowelToken(withAccent(accentCircumflex), withStrength(strengthWeak))
	TokenVowelStrongDieresisAccented   Token = newVowelToken(withAccent(accentDieresis), withStrength(strengthStrong))
	TokenVowelWeakDieresisAccented     Token = newVowelToken(withAccent(accentDieresis), withStrength(strengthWeak))
	TokenApostrophe                    Token = newToken(withRoot(rootApostrophe))
	TokenHyphen                        Token = newToken(withRoot(rootHyphen))
)

type Token interface {
	String() string
}

type token struct {
	root tokenRoot
}

func (t token) String() string {
	return t.root.String()
}

func (t token) IsVowel() bool {
	return t.root == rootVowel
}

type TokenVowel struct {
	token

	accent   accent
	strength strength
}

func (t TokenVowel) String() string {
	return t.root.String() + t.strength.String() + t.accent.String()
}

func (t TokenVowel) IsStrong() bool {
	return t.strength == strengthStrong
}

func (t TokenVowel) IsWeak() bool {
	return t.strength == strengthWeak
}

func (t TokenVowel) IsAcuteAccented() bool {
	return t.accent == accentAcute
}

func (t TokenVowel) IsGraveAccented() bool {
	return t.accent == accentGrave
}

func (t TokenVowel) IsCircumflexAccented() bool {
	return t.accent == accentCircumflex
}

func (t TokenVowel) IsDieresisAccented() bool {
	return t.accent == accentDieresis
}

type tokenOption func(*token)

func newToken(opts ...tokenOption) token {
	t := token{}
	for _, opt := range opts {
		opt(&t)
	}
	return t
}

func withRoot(r tokenRoot) tokenOption {
	return func(t *token) {
		t.root = r
	}
}

type tokenVowelOption func(*TokenVowel)

func newVowelToken(opts ...tokenVowelOption) TokenVowel {
	t := TokenVowel{
		token: newToken(withRoot(rootVowel)),
	}
	for _, opt := range opts {
		opt(&t)
	}
	return t
}

func withAccent(a accent) tokenVowelOption {
	return func(t *TokenVowel) {
		t.accent = a
	}
}

func withStrength(s strength) tokenVowelOption {
	return func(t *TokenVowel) {
		t.strength = s
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

type strength int

const (
	strengthNone strength = iota
	strengthWeak
	strengthStrong
)

func (s strength) String() string {
	strengthStrings := map[strength]string{
		strengthNone:   "",
		strengthWeak:   "Weak",
		strengthStrong: "Strong",
	}

	if str, found := strengthStrings[s]; found {
		return str
	}
	return "UndefinedStrength"
}
