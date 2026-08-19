package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) RenderSidebar(panelHeight int) string {
	chapters := m.renderChapterColumn()
	lessons := m.renderLessonColumn()
	joined := lipgloss.JoinHorizontal(lipgloss.Top, chapters, lessons)
	style := SidebarStyle
	if m.FocusedPanel == PanelChapters || m.FocusedPanel == PanelLessons {
		style = SidebarFocusedStyle
	}
	if panelHeight > 0 {
		style = style.Height(panelHeight)
	}
	return style.Render(joined)
}

func (m Model) renderChapterColumn() string {
	if m.Manifest == nil || len(m.Manifest.Categories) == 0 {
		return ColumnTitleStyle.Render("CHAPTERS")
	}

	title := ColumnTitleStyle.Render("CHAPTERS")
	lines := []string{title}

	for i, cat := range m.Manifest.Categories {
		var passed, total int
		for _, top := range cat.Topics {
			for _, ex := range top.Exercises {
				total++
				if m.Progress.Passed[ex.ID] {
					passed++
				}
			}
		}

		label := fmt.Sprintf("%d. %s", i+1, truncateRunes(cat.Title, 9))
		row := fmt.Sprintf("%-12s %d/%d", label, passed, total)

		switch {
		case i == m.SelectedChapter:
			lines = append(lines, SidebarSelectedStyle.Width(ChapterColWidth).Render(row))
		case m.FocusedPanel == PanelChapters:
			lines = append(lines, SidebarItemFocusStyle.Width(ChapterColWidth).Render(row))
		default:
			lines = append(lines, SidebarItemStyle.Width(ChapterColWidth).Render(row))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m Model) renderLessonColumn() string {
	cat := m.currentCategory()
	title := ColumnTitleStyle.Render(fmt.Sprintf("%s — LESSONS", truncateRunes(cat.Title, 12)))
	lines := []string{title}

	if len(cat.Topics) == 0 {
		lines = append(lines, SidebarItemStyle.Width(LessonColWidth).Render("No lessons."))
	}

	for i, top := range cat.Topics {
		var passed, total int
		for _, ex := range top.Exercises {
			total++
			if m.Progress.Passed[ex.ID] {
				passed++
			}
		}

		label := fmt.Sprintf("%d. %s", i+1, truncateRunes(top.Title, 19))
		row := fmt.Sprintf("%-22s %d/%d", label, passed, total)

		switch {
		case i == m.SelectedLesson:
			lines = append(lines, SidebarSelectedStyle.Width(LessonColWidth).Render(row))
		case m.FocusedPanel == PanelLessons:
			lines = append(lines, SidebarItemFocusStyle.Width(LessonColWidth).Render(row))
		default:
			lines = append(lines, SidebarItemStyle.Width(LessonColWidth).Render(row))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func truncateRunes(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n-1]) + "…"
}