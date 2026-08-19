package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func renderFooter() string {
	navigation := " [←↑↓→] Move  |  [Tab] Switch Panel  |  [r/Enter] Run Test  |  [n/p] Next/Prev"
	actions := " [u/m] Toggle Done  |  [h] Hint  |  [o] Open IDE  |  [s] Switch IDE  |  [q] Quit"

	sectionNav := FooterStyle.Render(navigation)
	sectionActions := FooterStyle.Render(actions)
	return lipgloss.JoinVertical(lipgloss.Left, sectionNav, sectionActions)
}