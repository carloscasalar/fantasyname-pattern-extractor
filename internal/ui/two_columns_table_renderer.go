package ui

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type TwoColumnsTableRenderer struct {
	styles            *lipGlossStyle
	maxWidth          int
	firstColumnTitle  string
	secondColumnTitle string
	rows              []twoColumnsTableRow
}

func (r *TwoColumnsTableRenderer) Render() string {
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Foreground(r.styles.titleColor).Bold(true)
	firstColumnStyle := baseStyle.Foreground(r.styles.valueColor)
	secondColumnStyle := firstColumnStyle.
		Width(r.getSecondColumnWith()).
		Italic(true)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(r.styles.borderColor)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}
			if col == 0 {
				return firstColumnStyle
			}
			return secondColumnStyle
		}).
		Headers(r.firstColumnTitle, r.secondColumnTitle)

	for _, row := range r.rows {
		t.Row(row.firstColumnValue(), row.secondColumnValue())
	}
	return t.Render()
}

func (r *TwoColumnsTableRenderer) getSecondColumnWith() int {
	const paddingsAndColumnsBorderLength = 5
	maxFirstColumnLength, maxSecondColumnLength := r.getMaxColumnLengths()
	secondColumnMaxWidth := r.maxWidth - maxFirstColumnLength - paddingsAndColumnsBorderLength
	secondColumnWidth := min(secondColumnMaxWidth, maxSecondColumnLength)
	return secondColumnWidth
}

func (r *TwoColumnsTableRenderer) getMaxColumnLengths() (int, int) {
	maxFirstColumnLength := len(r.firstColumnTitle)
	maxSecondColumnLength := len(r.secondColumnTitle)

	for _, row := range r.rows {
		maxFirstColumnLength = max(maxFirstColumnLength, len(row.firstColumnValue()))
		maxSecondColumnLength = max(maxSecondColumnLength, len(row.secondColumnValue()))
	}

	return maxFirstColumnLength, maxSecondColumnLength
}

type TwoColumnsTableBuilder struct {
	styles            *lipGlossStyle
	maxWidth          *int
	firstColumnTitle  *string
	secondColumnTitle *string
	rows              []twoColumnsTableRow
}

func (b *TwoColumnsTableBuilder) WithMaxWidth(maxWidth int) *TwoColumnsTableBuilder {
	b.maxWidth = &maxWidth
	return b
}

func (b *TwoColumnsTableBuilder) WithTitles(firstColumnTitle, secondColumnTitle string) *TwoColumnsTableBuilder {
	b.firstColumnTitle = &firstColumnTitle
	b.secondColumnTitle = &secondColumnTitle
	return b
}

func (b *TwoColumnsTableBuilder) WithRow(firstColumnValue, secondColumnValue string) *TwoColumnsTableBuilder {
	b.rows = append(b.rows, twoColumnsTableRow{firstColumnValue, secondColumnValue})
	return b
}

func (b *TwoColumnsTableBuilder) Build() *TwoColumnsTableRenderer {
	return &TwoColumnsTableRenderer{
		styles:            newLipGlossDefaultStyle(),
		maxWidth:          nvl(b.maxWidth, 80),
		firstColumnTitle:  nvl(b.firstColumnTitle, "First Column"),
		secondColumnTitle: nvl(b.secondColumnTitle, "Second Column"),
		rows:              b.rows,
	}
}

func nvl[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}
