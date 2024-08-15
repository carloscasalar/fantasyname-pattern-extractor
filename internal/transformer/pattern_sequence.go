package transformer

type patternSequence string

const (
	emptySequence patternSequence = ""

	sequenceVowel                         patternSequence = "v"
	sequenceVowelAcuteAccented            patternSequence = "(<v>|(á|é|í|ó|ú))"
	sequenceVowelGraveAccented            patternSequence = "(<v>|(à|è|ì|ò|ù))"
	sequenceVowelCircumflexAccented       patternSequence = "(<v>|(â|ê|î|ô|û))"
	sequenceVowelDieresisAccented         patternSequence = "(<v>|(ä|ë|ï|ö|ü))"
	sequenceVowelStrong                   patternSequence = "(<v>|(a|e|o)|(a|e|o))"
	sequenceVowelStrongAcuteAccented      patternSequence = "(<v>|(á|é|ó)|(á|é|ó))"
	sequenceVowelStrongGraveAccented      patternSequence = "(<v>|(à|è|ò)|(à|è|ò))"
	sequenceVowelStrongCircumflexAccented patternSequence = "(<v>|(â|ê|ô)|(â|ê|ô))"
	sequenceVowelStrongDieresisAccented   patternSequence = "(<v>|(ä|ë|ö)|(ä|ë|ö))"
	sequenceVowelWeak                     patternSequence = "(<v>|(i|u)|(i|u))"
	sequenceVowelWeakAcuteAccented        patternSequence = "(<v>|(í|ú)|(í|ú))"
	sequenceVowelWeakGraveAccented        patternSequence = "(<v>|(ì|ù)|(ì|ù))"
	sequenceVowelWeakCircumflexAccented   patternSequence = "(<v>|(î|û)|(î|û))"
	sequenceVowelWeakDieresisAccented     patternSequence = "(<v>|(ï|ü)|(ï|ü))"

	sequenceConsonant  patternSequence = "c"
	sequenceTildeN     patternSequence = "(<c>|ñ)"
	sequenceCedilla    patternSequence = "(<c>|ç)"
	sequenceApostrophe patternSequence = "('|)"
	sequenceHyphen     patternSequence = "(-|)"
)
