package ui

import "github.com/charmbracelet/lipgloss"

type TitleValueRenderer struct {
	style *lipGlossStyle
	title string
	value string
}

func NewTitleValueRenderer(title, value string) *TitleValueRenderer {
	return &TitleValueRenderer{
		style: newLipGlossDefaultStyle(),
		title: title,
		value: value,
	}
}

func (r TitleValueRenderer) Render() string {
	titleBox := lipgloss.NewStyle().
		Bold(true).
		PaddingLeft(1).
		Foreground(r.style.titleColor)
	valueBox := lipgloss.NewStyle().
		MarginLeft(1).
		Foreground(r.style.valueColor)
	return lipgloss.JoinHorizontal(lipgloss.Left, titleBox.Render(r.title), valueBox.Render(r.value))
}
