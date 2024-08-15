package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMainFunctionSnapshot tests the main function and compares the output to a snapshot. Kind of smoke test.
func TestMainFunctionSnapshot(t *testing.T) {
	inputsAndExpectedOutputs := []struct {
		options        []string
		expectedOutput string
	}{
		{[]string{"-n", "Tanis"}, " PATTERNS: !cvcvc, !c(<v>|(a|e|o)|(a|e|o))c(<v>|(i|u)|(i|u))c\n"},
		{[]string{"-n", "Iöwen"}, " PATTERNS: !v(<v>|(ä|ë|ï|ö|ü))cvc, !(<v>|(i|u)|(i|u))(<v>|(ä|ë|ö)|(ä|ë|ö))c(<v>|(a|e|o)|(a|e|o))c\n"},
	}

	for _, inputAndExpectedOutput := range inputsAndExpectedOutputs {
		t.Run("expected pattern for "+inputAndExpectedOutput.options[1], func(t *testing.T) {
			// Set up any necessary environment variables or arguments
			os.Args = append([]string{"cmd"}, inputAndExpectedOutput.options...)

			// Capture the output
			r, w, _ := os.Pipe()
			old := os.Stdout
			os.Stdout = w

			// Call the main function
			main()

			// Restore the original stdout
			w.Close()
			os.Stdout = old

			// Read the captured output
			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := buf.String()

			// Compare the output to the snapshot
			assert.Equal(t, inputAndExpectedOutput.expectedOutput, output)
		})
	}
}
