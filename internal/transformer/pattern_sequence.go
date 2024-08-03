package transformer

type patternSequence string

const (
	emptySequence patternSequence = ""

	sequenceVowel                   patternSequence = "v"
	sequenceVowelAcuteAccented      patternSequence = "(<v>|(á|é|í|ó|ú))"
	sequenceVowelGraveAccented      patternSequence = "(<v>|(à|è|ì|ò|ù))"
	sequenceVowelCircumflexAccented patternSequence = "(<v>|(â|ê|î|ô|û))"
	sequenceVowelDieresisAccented   patternSequence = "(<v>|(ä|ë|ï|ö|ü))"

	sequenceConsonant  patternSequence = "c"
	sequenceTildeN     patternSequence = "(<c>|ñ)"
	sequenceCedilla    patternSequence = "(<c>|ç)"
	sequenceApostrophe patternSequence = "('|)"
	sequenceHyphen     patternSequence = "(-|)"
)
