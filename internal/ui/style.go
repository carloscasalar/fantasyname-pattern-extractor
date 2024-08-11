package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type lipGlossStyle struct {
	titleColor  lipgloss.AdaptiveColor
	valueColor  lipgloss.AdaptiveColor
	borderColor lipgloss.AdaptiveColor
}

func newLipGlossDefaultStyle() *lipGlossStyle {
	return &lipGlossStyle{
		titleColor:  lipgloss.AdaptiveColor{Light: "202", Dark: "252"},
		valueColor:  lipgloss.AdaptiveColor{Light: "#3C3C3C", Dark: "#04B575"},
		borderColor: lipgloss.AdaptiveColor{Light: "99", Dark: "99"},
	}
}
