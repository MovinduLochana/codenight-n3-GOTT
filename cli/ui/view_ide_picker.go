package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) RenderIdePicker(width, height int) string {
	if width == 0 {
		width = 80
	}
	if height == 0 {
		height = 24
	}

	var lines []string
	lines = append(lines, SectionTitleStyle.Render("Open With — choose your preferred IDE"))

	if len(m.DetectedIDEs) == 0 {
		lines = append(lines, lipgloss.NewStyle().Foreground(ColorPending).Render("  No IDE detected — enter a command below."))
	}

	for i, e := range m.DetectedIDEs {
		marker := "  "
		style := lipgloss.NewStyle().Foreground(ColorPending)
		if !m.IdeCustom && i == m.IdePickerIdx {
			marker = "▶ "
			style = lipgloss.NewStyle().Foreground(ColorHighlight).Bold(true)
		}
		lines = append(lines, style.Render(fmt.Sprintf(" %s%s", marker, e.Name)))
	}

	customIdx := len(m.DetectedIDEs)
	if m.IdeCustom {
		lines = append(lines, lipgloss.NewStyle().Foreground(ColorHighlight).Bold(true).Render(" ▶ " + m.ideInput.View()))
	} else {
		marker := "  "
		if customIdx == m.IdePickerIdx {
			marker = "▶ "
		}
		lines = append(lines, lipgloss.NewStyle().Foreground(ColorPending).Render(fmt.Sprintf(" %sCustom command…", marker)))
	}

	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Foreground(ColorPending).Render("  [↑/↓] Choose  [Enter] Open  [Esc] Cancel"))

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorHighlight).
		Padding(1, 2).
		Width(46).
		Render(lipgloss.JoinVertical(lipgloss.Left, lines...))

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}