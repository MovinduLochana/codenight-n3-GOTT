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