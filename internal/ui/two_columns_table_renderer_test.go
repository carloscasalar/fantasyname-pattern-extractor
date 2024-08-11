package ui_test

import (
	"testing"

	"github.com/carloscasalar/fantasyname-pattern-extractor/internal/ui"
	"github.com/stretchr/testify/assert"
)

func TestTwoColumnsTableRenderer_Render(t *testing.T) {
	const paddingLen = 1
	const borderLen = 1
	const sideBorderLen = borderLen + paddingLen
	const middleBorderLen = paddingLen + borderLen + paddingLen

	testCases := map[string]struct {
		titles                [2]string
		row                   [2]string
		maxTableWidth         int
		expectedRenderedTable string
	}{
		"should render row without breaks if there is space": {
			[2]string{"title 1", "title 2"},
			[2]string{"value 1", "value 2"},
			80,
			"┌─────────┬─────────┐\n" +
				"│ title 1 │ title 2 │\n" +
				"├─────────┼─────────┤\n" +
				"│ value 1 │ value 2 │\n" +
				"│         │         │\n" +
				"└─────────┴─────────┘",
		},
		"should wrap second column when the row doesn't fit the max width": {
			[2]string{"PATTERN", "EXAMPLES"},
			[2]string{"!cvcc", "Soqr, Vebk, Watm, Wegp, Jetc, Maff, Zikt, Pizb, Loht, Tiws"},
			70,
			"┌─────────┬──────────────────────────────────────────────────────────┐\n" +
				"│ PATTERN │ EXAMPLES                                                 │\n" +
				"├─────────┼──────────────────────────────────────────────────────────┤\n" +
				"│ !cvcc   │ Soqr, Vebk, Watm, Wegp, Jetc, Maff, Zikt, Pizb, Loht,    │\n" +
				"│         │ Tiws                                                     │\n" +
				"└─────────┴──────────────────────────────────────────────────────────┘",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			table := new(ui.TwoColumnsTableBuilder).
				WithTitles(tc.titles[0], tc.titles[1]).
				WithRow(tc.row[0], tc.row[1]).
				WithMaxWidth(tc.maxTableWidth).
				Build()

			renderedTable := table.Render()

			assert.Equal(t, tc.expectedRenderedTable, renderedTable)
		})
	}
}
