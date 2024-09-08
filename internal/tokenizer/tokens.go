package tokenizer

var (
	TokenConsonant                     Token = newToken(withRoot(rootConsonant))
	TokenTildeN                        Token = newToken(withRoot(rootTildeN))
	TokenCedilla                       Token = newToken(withRoot(rootCedilla))
	TokenVowelStrong                         = newVowelToken(withStrength(strengthStrong))
	TokenVowelWeak                           = newVowelToken(withStrength(strengthWeak))
	TokenVowelStrongAcuteAccented            = newVowelToken(withAccent(accentAcute), withStrength(strengthStrong))
	TokenVowelWeakAcuteAccented              = newVowelToken(withAccent(accentAcute), withStrength(strengthWeak))
	TokenVowelStrongGraveAccented            = newVowelToken(withAccent(accentGrave), withStrength(strengthStrong))
	TokenVowelWeakGraveAccented              = newVowelToken(withAccent(accentGrave), withStrength(strengthWeak))
	TokenVowelStrongCircumflexAccented       = newVowelToken(withAccent(accentCircumflex), withStrength(strengthStrong))
	TokenVowelWeakCircumflexAccented         = newVowelToken(withAccent(accentCircumflex), withStrength(strengthWeak))
	TokenVowelStrongDieresisAccented         = newVowelToken(withAccent(accentDieresis), withStrength(strengthStrong))
	TokenVowelWeakDieresisAccented           = newVowelToken(withAccent(accentDieresis), withStrength(strengthWeak))
	TokenApostrophe                    Token = newToken(withRoot(rootApostrophe))
	TokenHyphen                        Token = newToken(withRoot(rootHyphen))
)

type Token interface {
	String() string
	IsVowel() (bool, *TokenVowel)
}

type token struct {
	root tokenRoot
}

func (t token) String() string {
	return t.root.String()
}

func (t token) IsVowel() (bool, *TokenVowel) {
	if t.root == rootVowel {
		return true, &TokenVowel{token: t}
	}
	return false, nil
}

type TokenVowel struct {
	token

	accent   accent
	strength strength
}

func (t TokenVowel) String() string {
	return t.root.String() + t.strength.String() + t.accent.String()
}

func (t TokenVowel) IsVowel() (bool, *TokenVowel) {
	return true, &t
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

//go:generate stringer -type=tokenRoot -output=tokens_root_auto.go -trimprefix=root
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

//go:generate stringer -type=accent -output=tokens_accent_auto.go -trimprefix=accent
type accent int

const (
	accentNone accent = iota
	accentAcute
	accentGrave
	accentCircumflex
	accentDieresis
)

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
