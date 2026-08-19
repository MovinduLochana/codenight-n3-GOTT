package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/ide"
)

func TestSmokeRender(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()
	m.WindowWidth = 140
	m.WindowHeight = 40
	m.recalculateLayout()

	view := m.View()
	if view == "" || strings.Contains(view, "Loading") {
		t.Fatalf("unexpected view:\n%s", view)
	}
	t.Logf("DEFAULT VIEW\n%s", view)

	// Move to lessons panel and select lesson 1
	m.FocusedPanel = PanelChapters
	m.SelectedChapter = 1
	m.focusLesson(1, 0)
	m.FocusedPanel = PanelTasks
	view2 := m.View()
	t.Logf("CHAPTER 2 VIEW\n%s", view2)
}

func TestCompletionTicks(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()

	press := func(s string) {
		nm, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)})
		m = nm.(Model)
	}

	// Mark all tasks of chapter 0, lesson 0 as complete.
	m.SelectedChapter = 0
	m.SelectedLesson = 0
	start, end := m.lessonStartIdx(), m.lessonEndIdx()
	for i := start; i < end; i++ {
		m.Progress.Passed[m.Exercises[i].ID] = true
	}

	if got := m.renderLessonColumn(); !strings.Contains(got, "✓") {
		t.Fatalf("completed lesson should show green tick:\n%s", got)
	}

	// Un-marking one task removes the lesson tick.
	m.Progress.Passed[m.Exercises[start].ID] = false
	if got := m.renderLessonColumn(); strings.Contains(got, "✓") {
		t.Fatalf("incomplete lesson should not show tick:\n%s", got)
	}
	m.Progress.Passed[m.Exercises[start].ID] = true

	// u/m key toggles the focused task.
	id := m.Exercises[m.FocusedIdx].ID
	before := m.Progress.Passed[id]
	press("u")
	if m.Progress.Passed[id] == before {
		t.Fatalf("u should toggle task completion")
	}
	press("m")
	if m.Progress.Passed[id] != before {
		t.Fatalf("m should toggle task back")
	}

	// Mark every task in the whole chapter -> chapter gets a tick.
	for _, ex := range m.Exercises[:m.lessonEndIdx()+100] {
		if m.chapterNo() == 1 {
			m.Progress.Passed[ex.ID] = true
		}
	}
	for i, ex := range m.Exercises {
		if i < m.lessonStartIdx() {
			m.Progress.Passed[ex.ID] = true
		}
	}
	if got := m.renderChapterColumn(); !strings.Contains(got, "✓") {
		t.Fatalf("completed chapter should show green tick:\n%s", got)
	}
}

func TestSwitchIDEKey(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()

	// Even with a saved IDE, s re-opens the picker.
	if len(m.DetectedIDEs) == 0 {
		t.Skip("no IDEs detected on this machine")
	}
	m.Progress.PreferredIDE = m.DetectedIDEs[0].Key
	nm, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("s")})
	m = nm.(Model)
	if !m.ShowIdePicker {
		t.Fatal("s should open the IDE picker")
	}
	if m.IdePickerIdx != 0 || m.DetectedIDEs[m.IdePickerIdx].Key != m.Progress.PreferredIDE {
		t.Fatalf("picker should preselect current IDE, idx=%d", m.IdePickerIdx)
	}

	// Selecting a different IDE updates and persists the preference.
	target := 1 % len(m.DetectedIDEs)
	if target == 0 {
		target = len(m.DetectedIDEs) - 1
	}
	m.IdePickerIdx = target
	nm, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = nm.(Model)
	if m.ShowIdePicker {
		t.Fatal("picker should close after selecting")
	}
	if m.Progress.PreferredIDE != m.DetectedIDEs[target].Key {
		t.Fatalf("preferred IDE not updated: %q", m.Progress.PreferredIDE)
	}
	if cmd == nil {
		t.Fatal("expected launch cmd after switching")
	}
}

func TestNavigation(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()

	key := func(s string) tea.Msg { return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(s)} }
	press := func(s string) {
		nm, _ := m.Update(key(s))
		m = nm.(Model)
	}

	// Start in chapters panel, move down → chapter 2, lesson reset to first.
	m.FocusedPanel = PanelChapters
	m.SelectedChapter = 0
	m.SelectedLesson = 0
	press("down")
	if m.SelectedChapter != 1 || m.SelectedLesson != 0 {
		t.Fatalf("after down: chapter=%d lesson=%d", m.SelectedChapter, m.SelectedLesson)
	}

	// Right → lessons panel; down → second lesson; FocusedIdx points into lesson.
	press("right")
	if m.FocusedPanel != PanelLessons {
		t.Fatalf("panel=%v", m.FocusedPanel)
	}
	startBefore := m.lessonStartIdx()
	press("down")
	if m.SelectedLesson != 1 || m.lessonStartIdx() <= startBefore {
		t.Fatalf("lesson=%d start %d -> %d", m.SelectedLesson, startBefore, m.lessonStartIdx())
	}

	// Right → tasks panel; n/p bounded within the lesson.
	press("right")
	if m.FocusedPanel != PanelTasks {
		t.Fatalf("panel=%v", m.FocusedPanel)
	}
	start, end := m.lessonStartIdx(), m.lessonEndIdx()
	if m.FocusedIdx != start {
		t.Fatalf("focus should be lesson start %d, got %d", start, m.FocusedIdx)
	}
	press("n")
	press("n")
	if m.FocusedIdx >= end {
		t.Fatalf("n overran lesson end %d: %d", end, m.FocusedIdx)
	}
	press("p")
	press("p")
	if m.FocusedIdx < start {
		t.Fatalf("p underran lesson start %d: %d", start, m.FocusedIdx)
	}

	// Left at task edge switches back to lessons panel.
	m.FocusedIdx = start
	press("left")
	if m.FocusedPanel != PanelLessons {
		t.Fatalf("left from task edge should go to lessons, got %v", m.FocusedPanel)
	}
}

func TestIdePicker(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()
	m.Progress.PreferredIDE = ""

	if len(m.DetectedIDEs) > 0 {
		// o with no preferred IDE should open the picker.
		nm, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("o")})
		m = nm.(Model)
		if !m.ShowIdePicker {
			t.Fatal("expected picker to open on o")
		}
		view := m.RenderIdePicker(120, 30)
		if !strings.Contains(view, "Open With") {
			t.Fatalf("picker view missing title:\n%s", view)
		}

		// esc closes the picker.
		nm, _ = m.Update(tea.KeyMsg{Type: tea.KeyEscape})
		m = nm.(Model)
		if m.ShowIdePicker {
			t.Fatal("esc should close picker")
		}
	}

	// With no detected IDEs, picker opens in custom-command mode.
	m.DetectedIDEs = nil
	m.Progress.PreferredIDE = ""
	nm, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("o")})
	m = nm.(Model)
	if !m.ShowIdePicker || !m.IdeCustom {
		t.Fatalf("expected custom picker, show=%v custom=%v", m.ShowIdePicker, m.IdeCustom)
	}
	if cmd == nil {
		t.Fatal("expected blink cmd")
	}
	view := m.RenderIdePicker(120, 30)
	if !strings.Contains(view, "No IDE detected") {
		t.Fatalf("expected no-IDE message:\n%s", view)
	}

	// Type a command and submit.
	m.ideInput.SetValue("/bin/true")
	nm, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = nm.(Model)
	if m.ShowIdePicker || m.IdeCustom {
		t.Fatalf("picker should close, show=%v custom=%v", m.ShowIdePicker, m.IdeCustom)
	}
	if m.Progress.PreferredIDE != "/bin/true" {
		t.Fatalf("preferred IDE not saved: %q", m.Progress.PreferredIDE)
	}
	if cmd == nil {
		t.Fatal("expected open cmd after submit")
	}
	res := cmd().(ide.ResultMsg)
	if res.Err != nil {
		t.Fatalf("open failed: %v", res.Err)
	}
}

func TestOpenInIDECmd(t *testing.T) {
	m, err := NewModel()
	if err != nil {
		t.Fatalf("NewModel: %v", err)
	}
	m.RootDir = t.TempDir()
	if len(m.Exercises) == 0 {
		t.Fatal("no exercises")
	}
	m.Progress.PreferredIDE = "/bin/true"
	m.ShowIdePicker = false

	model, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("o")})
	if cmd == nil {
		t.Fatal("expected cmd from o key")
	}
	msg := cmd()
	res, ok := msg.(ide.ResultMsg)
	if !ok {
		t.Fatalf("expected ide.ResultMsg, got %T", msg)
	}
	if res.Err != nil {
		t.Fatalf("openInIDE error: %v", res.Err)
	}

	model, _ = model.Update(res)
	if !strings.Contains(model.View(), "IDE opened") {
		t.Fatalf("status not shown after open:\n%s", model.View())
	}
	t.Logf("OPENED: %+v", res)
}
