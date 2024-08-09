package examples_test

import (
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/examples"

	"github.com/stretchr/testify/assert"
)

func TestGenerateExamples_(t *testing.T) {
	tests := map[string]struct {
		pattern          string
		numberOfExamples uint
		expected         string
	}{
		"Empty pattern should return no examples": {
			pattern:          "",
			numberOfExamples: 100,
			expected:         "",
		},
		"Zero examples should return no examples": {
			pattern:          "vcv",
			numberOfExamples: 0,
			expected:         "",
		},
		"Valid pattern and non-zero number should generate the expected number of examples": {
			pattern:          "(some pattern)",
			numberOfExamples: 3,
			expected:         "some pattern, some pattern, some pattern",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := examples.Generate(tt.pattern, tt.numberOfExamples)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestGenerateExamples_ErrorCases(t *testing.T) {
	tests := map[string]struct {
		pattern          string
		numberOfExamples uint
	}{
		"Invalid pattern should result in error": {
			pattern:          "a>",
			numberOfExamples: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := examples.Generate(tt.pattern, tt.numberOfExamples)
			assert.Error(t, err)
			assert.Equal(t, "", got)
		})
	}
}
