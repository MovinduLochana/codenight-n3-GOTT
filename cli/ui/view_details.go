package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/fox"
)

func renderMarkdown(mdContent string, wordWrap int) string {
	// Use a fixed dark style instead of WithAutoStyle: auto style queries the
	// terminal's background color (OSC 11) on every renderer construction and
	// blocks up to 5s waiting for a reply, freezing the TUI. The app UI is
	// dark-themed, so the dark style matches it regardless of the terminal.
	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(wordWrap),
	)
	if err != nil {
		return mdContent
	}
	out, err := r.Render(mdContent)
	if err != nil {
		return mdContent
	}
	return strings.TrimSpace(out)
}

func (m Model) RenderRightDetails(rightWidth int, panelHeight int) string {
	if len(m.Exercises) == 0 {
		return RightPanelStyle.Width(rightWidth).Height(panelHeight).Render("Select an exercise.")
	}

	curEx := m.Exercises[m.FocusedIdx]
	isPassed := m.Progress.Passed[curEx.ID]

	contentWidth := rightWidth - 4
	if contentWidth < 30 {
		contentWidth = 30
	}
	wrapStyle := lipgloss.NewStyle().Width(contentWidth)

	// Header details
	totalTasks := len(m.currentTopic().Exercises)
	if totalTasks < 1 {
		totalTasks = 1
	}
	titleText := fmt.Sprintf("%s (Question %d/%d)", curEx.Title, curEx.TopicExerciseNum, totalTasks)
	title := SectionTitleStyle.Render(titleText)
	meta := wrapStyle.Render(fmt.Sprintf("Chapter:  %s (%d)\nLesson:   %s\nFile:     %s",
		curEx.CategoryTitle, m.chapterNo(), curEx.TopicTitle, curEx.FilePath))

	// Level badge
	var levelBadge string
	switch strings.ToLower(curEx.Level) {
	case "beginner":
		levelBadge = BadgeBeginnerStyle.Render("[BEGINNER]")
	case "intermediate":
		levelBadge = BadgeIntermediateStyle.Render("[INTERMEDIATE]")
	case "advanced":
		levelBadge = BadgeAdvancedStyle.Render("[ADVANCED]")
	default:
		levelBadge = fmt.Sprintf("[%s]", strings.ToUpper(curEx.Level))
	}

	// Status string
	statusStr := lipgloss.NewStyle().Foreground(ColorPending).Render("STATUS: PENDING [ ]")
	if isPassed {
		statusStr = lipgloss.NewStyle().Foreground(ColorSuccess).Bold(true).Render("STATUS: PASSED [✓]")
	}

	headerBlock := fmt.Sprintf("%s  %s\n%s\n%s\n", title, levelBadge, meta, statusStr)

	// Content area: Fox Animation vs Test Output vs Markdown Hint
	var body string

	if m.IsTesting {
		foxAscii := fox.GetFrame(m.FoxFrame)
		body = FoxBoxStyle.Render(foxAscii)
	} else if m.ShowHint {
		hintPath := filepath.Join(m.RootDir, curEx.DocPath)
		data, err := os.ReadFile(hintPath)
		if err != nil {
			body = lipgloss.NewStyle().Foreground(ColorFailure).Render("No hint file found.")
		} else {
			renderedMd := renderMarkdown(string(data), contentWidth)
			body = fmt.Sprintf("%s\n%s",
				lipgloss.NewStyle().Foreground(ColorHighlight).Bold(true).Render("=== TASK EXPLANATION & HINTS ==="),
				renderedMd)
		}
	} else if m.TestOutput != "" {
		outputStyle := lipgloss.NewStyle().Foreground(ColorFailure).Width(contentWidth)
		if isPassed {
			outputStyle = lipgloss.NewStyle().Foreground(ColorSuccess).Width(contentWidth)
		}
		body = fmt.Sprintf("%s\n%s",
			lipgloss.NewStyle().Foreground(ColorHighlight).Bold(true).Render("=== LATEST TEST OUTPUT ==="),
			outputStyle.Render(m.TestOutput))
	} else if m.StatusMsg != "" {
		body = lipgloss.NewStyle().Foreground(ColorHighlight).Width(contentWidth).Render(m.StatusMsg)
	} else {
		body = wrapStyle.Foreground(lipgloss.Color("244")).Render(
			"Press [r] or [Enter] to run tests for this exercise.\nPress [h] to view task instructions and hints.\nPress [o] to open this task in your IDE.")
	}

	rightContent := lipgloss.JoinVertical(lipgloss.Left, headerBlock, body)

	style := RightPanelStyle.Width(rightWidth)
	if m.FocusedPanel == PanelDetails {
		style = style.BorderForeground(ColorHighlight)
	}
	if panelHeight > 0 {
		style = style.Height(panelHeight)
	}
	return style.Render(rightContent)
}

func (m Model) View() string {
	if len(m.Exercises) == 0 {
		return "Loading exercises..."
	}

	// Compute outer panel height matching inner content + borders
	panelHeight := (m.GridRows * 5) + 4

	rightWidth := 52
	if m.WindowWidth > 0 {
		leftWidth := sidebarWidth + (m.GridCols * 18) + 4
		availW := m.WindowWidth - leftWidth - 6
		if availW > 35 {
			rightWidth = availW
		}
	}

	// 1. Header Bar
	passedCount := 0
	for _, ex := range m.Exercises {
		if m.Progress.Passed[ex.ID] {
			passedCount++
		}
	}
	total := len(m.Exercises)
	pct := 0
	if total > 0 {
		pct = (passedCount * 100) / total
	}

	headerText := fmt.Sprintf(" GOSTLINGS — Mozilla Campus Club of SLIIT  |  Progress: %d/%d Passed (%d%%) ", passedCount, total, pct)
	header := TitleStyle.Render(headerText)

	// 2. Main Panels (Sidebar + Task Grid + Details)
	sidebar := m.RenderSidebar(panelHeight)
	leftGrid := m.RenderLeftGrid(panelHeight)
	rightDetails := m.RenderRightDetails(rightWidth, panelHeight)
	mainPanels := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, leftGrid, rightDetails)

	// 3. Controls / Footer Bar
	footer := renderFooter()

	view := lipgloss.JoinVertical(lipgloss.Left, header, mainPanels, footer)

	if m.ShowIdePicker {
		return m.RenderIdePicker(m.WindowWidth, m.WindowHeight)
	}

	return view
}
