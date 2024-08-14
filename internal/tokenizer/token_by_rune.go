package tokenizer

var tokenByRune = map[rune]Token{
	'b': TokenConsonant,
	'c': TokenConsonant,
	'd': TokenConsonant,
	'f': TokenConsonant,
	'g': TokenConsonant,
	'h': TokenConsonant,
	'j': TokenConsonant,
	'k': TokenConsonant,
	'l': TokenConsonant,
	'm': TokenConsonant,
	'n': TokenConsonant,
	'p': TokenConsonant,
	'q': TokenConsonant,
	'r': TokenConsonant,
	's': TokenConsonant,
	't': TokenConsonant,
	'v': TokenConsonant,
	'w': TokenConsonant,
	'x': TokenConsonant,
	'y': TokenConsonant,
	'z': TokenConsonant,

	'a': TokenVowelStrong,
	'e': TokenVowelStrong,
	'i': TokenVowelWeak,
	'o': TokenVowelStrong,
	'u': TokenVowelWeak,

	'á': TokenVowelStrongAcuteAccented,
	'é': TokenVowelStrongAcuteAccented,
	'í': TokenVowelWeakAcuteAccented,
	'ó': TokenVowelStrongAcuteAccented,
	'ú': TokenVowelWeakAcuteAccented,

	'à': TokenVowelStrongGraveAccented,
	'è': TokenVowelStrongGraveAccented,
	'ì': TokenVowelWeakGraveAccented,
	'ò': TokenVowelStrongGraveAccented,
	'ù': TokenVowelWeakGraveAccented,

	'â': TokenVowelStrongCircumflexAccented,
	'ê': TokenVowelStrongCircumflexAccented,
	'î': TokenVowelWeakCircumflexAccented,
	'ô': TokenVowelStrongCircumflexAccented,
	'û': TokenVowelWeakCircumflexAccented,

	'ä': TokenVowelStrongDieresisAccented,
	'ë': TokenVowelStrongDieresisAccented,
	'ï': TokenVowelWeakDieresisAccented,
	'ö': TokenVowelStrongDieresisAccented,
	'ü': TokenVowelWeakDieresisAccented,

	'ç': TokenCedilla,
	'ñ': TokenTildeN,

	'\'': TokenApostrophe,
	'-':  TokenHyphen,
}
