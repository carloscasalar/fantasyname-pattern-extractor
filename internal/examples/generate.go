package examples

import (
	"math/rand"
	"strings"

	"github.com/s0rg/fantasyname"
)

const emptyNames = ""

// Generate as many examples as numberOfExamples from the given pattern as a comma-separated string
func Generate(pattern string, numberOfExamples uint) (string, error) {
	if pattern == "" || numberOfExamples == 0 {
		return emptyNames, nil
	}
	examples := make([]string, numberOfExamples)
	for i := 0; i < int(numberOfExamples); i++ {
		gen, err := fantasyname.Compile(pattern, fantasyname.Collapse(true), fantasyname.RandFn(rand.Intn))
		if err != nil {
			return emptyNames, err
		}
		examples[i] = gen.String()
	}
	return commaSeparated(examples), nil
}

func commaSeparated(examples []string) string {
	if len(examples) == 0 {
		return ""
	}

	joinedString := new(strings.Builder)
	for i, example := range examples {
		joinedString.WriteString(example)
		if i < len(examples)-1 {
			joinedString.WriteString(", ")
		}
	}

	return joinedString.String()
}
