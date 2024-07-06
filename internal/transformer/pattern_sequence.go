package transformer

type patternSequence string

const (
	sequenceVowel                   = "v"
	sequenceVowelAcuteAccented      = "(<v>|á|é|í|ó|ú)"
	sequenceVowelGraveAccented      = "(<v>|à|è|ì|ò|ù)"
	sequenceVowelCircumflexAccented = "(<v>|â|ê|î|ô|û)"
	sequenceVowelDieresisAccented   = "(<v>|ä|ë|ï|ö|ü)"
	sequenceConsonant               = "c"
	sequenceTildeN                  = "(<c>|ñ)"
	sequenceCedilla                 = "(<c>|ç)"
	sequenceApostrophe              = "('|)"
	sequenceHyphen                  = "(-|)"
)
