package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
)

func renderCard(ex manifest.Exercise, isFocused bool, isPassed bool, isFailed bool) string {
	var style lipgloss.Style
	statusSymbol := "[ ]"

	if isPassed {
		statusSymbol = "[✓]"
		style = CardPassedStyle
	} else if isFailed {
		statusSymbol = "[✗]"
		style = CardFailedStyle
	} else {
		style = CardPendingStyle
	}

	if isFocused {
		style = CardFocusedStyle
	}

	// Shorten Category Title if needed
	catShort := ex.CategoryTitle
	if len(catShort) > 13 {
		catShort = catShort[:12] + "…"
	}

	// Shorten Topic Title and append Question Number (1 to 3)
	topShort := ex.TopicTitle
	if len(topShort) > 10 {
		topShort = topShort[:9] + "…"
	}
	topWithNum := fmt.Sprintf("%s %d", topShort, ex.TopicExerciseNum)

	content := fmt.Sprintf("%s\n%s\n%s", catShort, topWithNum, statusSymbol)
	return style.Render(content)
}

func (m Model) RenderLeftGrid(panelHeight int) string {
	if len(m.Exercises) == 0 {
		return LeftPanelStyle.Height(panelHeight).Render("No exercises found.")
	}

	startIdx := m.lessonStartIdx()
	endIdx := m.lessonEndIdx()
	lessonExercises := m.Exercises[startIdx:endIdx]
	if len(lessonExercises) == 0 {
		return LeftPanelStyle.Height(panelHeight).Render("Select a lesson.")
	}

	pageSize := m.GridCols * m.GridRows
	if pageSize <= 0 {
		pageSize = 12
	}

	currentPage := (m.FocusedIdx - startIdx) / pageSize
	totalPages := (len(lessonExercises) + pageSize - 1) / pageSize
	pageStart := startIdx + currentPage*pageSize
	pageEnd := pageStart + pageSize
	if pageEnd > endIdx {
		pageEnd = endIdx
	}

	pageExercises := m.Exercises[pageStart:pageEnd]

	var rows []string
	var currentRowCards []string

	for i, ex := range pageExercises {
		actualIdx := pageStart + i
		isFocused := actualIdx == m.FocusedIdx
		isPassed := m.Progress.Passed[ex.ID]
		isFailed := !isPassed && m.TestOutput != "" && isFocused

		cardStr := renderCard(ex, isFocused, isPassed, isFailed)
		currentRowCards = append(currentRowCards, cardStr)

		if len(currentRowCards) == m.GridCols || i == len(pageExercises)-1 {
			rowStr := lipgloss.JoinHorizontal(lipgloss.Top, currentRowCards...)
			rows = append(rows, rowStr)
			currentRowCards = nil
		}
	}

	gridStr := lipgloss.JoinVertical(lipgloss.Left, rows...)

	// Pagination Footer inside Left Panel
	pagStr := fmt.Sprintf("\n  ← Page %d of %d →  ", currentPage+1, totalPages)
	pagStyled := lipgloss.NewStyle().Foreground(lipgloss.Color("244")).Render(pagStr)

	fullLeftContent := lipgloss.JoinVertical(lipgloss.Left, gridStr, pagStyled)

	style := LeftPanelStyle
	if m.FocusedPanel == PanelTasks {
		style = style.BorderForeground(ColorHighlight)
	}
	if panelHeight > 0 {
		style = style.Height(panelHeight)
	}
	return style.Render(fullLeftContent)
}
