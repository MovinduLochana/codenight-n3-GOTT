package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/fox"
	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/ide"
	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/runner"
)

// sidebarWidth is the combined width of the chapters + lessons columns
// including the sidebar border and padding.
const sidebarWidth = 48

type Panel int

const (
	PanelChapters Panel = iota
	PanelLessons
	PanelTasks
	PanelDetails
)

func (p Panel) String() string {
	switch p {
	case PanelChapters:
		return "Chapters"
	case PanelLessons:
		return "Lessons"
	case PanelTasks:
		return "Tasks"
	case PanelDetails:
		return "Details"
	}
	return "Unknown"
}

type Model struct {
	RootDir       string
	Manifest      *manifest.Manifest
	Exercises     []manifest.Exercise
	Progress      manifest.Progress
	FocusedIdx    int
	PageSize      int
	GridCols      int
	GridRows      int
	IsTesting     bool
	FoxFrame      int
	TestOutput    string
	StatusMsg     string
	ShowHint      bool
	WindowWidth   int
	WindowHeight  int
	Viewport      viewport.Model
	ViewportReady bool

	SelectedChapter int
	SelectedLesson  int
	FocusedPanel    Panel

	DetectedIDEs  []ide.IDE
	ShowIdePicker bool
	IdePickerIdx  int
	IdeCustom     bool
	ideInput      textinput.Model
}

func NewModel() (Model, error) {
	rootDir := manifest.FindRootDir()
	m, exercises, err := manifest.LoadManifest(rootDir)
	if err != nil {
		return Model{}, err
	}
	progress := manifest.LoadProgress(rootDir)

	focusedIdx := 0
	selectedChapter := 0
	selectedLesson := 0

	if progress.LastID != "" {
		flat := 0
		for ci, cat := range m.Categories {
			for li, top := range cat.Topics {
				for _, ex := range top.Exercises {
					if ex.ID == progress.LastID {
						selectedChapter = ci
						selectedLesson = li
						focusedIdx = flat
					}
					flat++
				}
			}
		}
	}

	ideInput := textinput.New()
	ideInput.Placeholder = "e.g. code"
	ideInput.Prompt = "> "

	return Model{
		RootDir:         rootDir,
		Manifest:        m,
		Exercises:       exercises,
		Progress:        progress,
		FocusedIdx:      focusedIdx,
		SelectedChapter: selectedChapter,
		SelectedLesson:  selectedLesson,
		FocusedPanel:    PanelTasks,
		GridCols:        4,
		GridRows:        3,
		PageSize:        12,
		ShowHint:        false,
		DetectedIDEs:    ide.Detect(),
		ideInput:        ideInput,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) currentCategory() manifest.Category {
	if m.Manifest == nil || m.SelectedChapter < 0 || m.SelectedChapter >= len(m.Manifest.Categories) {
		return manifest.Category{}
	}
	return m.Manifest.Categories[m.SelectedChapter]
}

func (m Model) currentTopic() manifest.Topic {
	cat := m.currentCategory()
	if m.SelectedLesson < 0 || m.SelectedLesson >= len(cat.Topics) {
		return manifest.Topic{}
	}
	return cat.Topics[m.SelectedLesson]
}

// lessonStartIdx returns the flattened index of the selected lesson's first exercise.
func (m Model) lessonStartIdx() int {
	if m.Manifest == nil {
		return 0
	}
	idx := 0
	for ci, cat := range m.Manifest.Categories {
		for li, top := range cat.Topics {
			if ci == m.SelectedChapter && li == m.SelectedLesson {
				return idx
			}
			idx += len(top.Exercises)
		}
	}
	return 0
}

// lessonEndIdx returns the flattened index just past the selected lesson's exercises.
func (m Model) lessonEndIdx() int {
	return m.lessonStartIdx() + len(m.currentTopic().Exercises)
}

func (m Model) chapterNo() int {
	return m.SelectedChapter + 1
}

func (m *Model) focusLesson(ci, li int) {
	m.SelectedChapter = ci
	m.SelectedLesson = li
	if start := m.lessonStartIdx(); start < len(m.Exercises) {
		m.FocusedIdx = start
	}
	if len(m.Exercises) == 0 {
		return
	}
	m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
	manifest.SaveProgress(m.RootDir, m.Progress)
}

func (m *Model) focusTask(idx int) {
	if len(m.Exercises) == 0 {
		return
	}
	if idx < 0 {
		idx = 0
	}
	if idx >= len(m.Exercises) {
		idx = len(m.Exercises) - 1
	}
	m.FocusedIdx = idx
	m.Progress.LastID = m.Exercises[idx].ID
	manifest.SaveProgress(m.RootDir, m.Progress)
}

func (m *Model) recalculateLayout() {
	if m.WindowWidth == 0 || m.WindowHeight == 0 {
		return
	}

	availWidth := m.WindowWidth - sidebarWidth - 6
	leftCols := (availWidth * 55) / 100 / 18
	if leftCols < 1 {
		leftCols = 1
	}
	if leftCols > 4 {
		leftCols = 4
	}
	m.GridCols = leftCols

	availHeight := m.WindowHeight - 4
	maxRows := (availHeight - 4) / 5
	if maxRows < 1 {
		maxRows = 1
	}
	if maxRows > 4 {
		maxRows = 4
	}
	m.GridRows = maxRows

	m.PageSize = m.GridCols * m.GridRows
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.WindowWidth = msg.Width
		m.WindowHeight = msg.Height
		m.recalculateLayout()

		rightWidth := m.WindowWidth - sidebarWidth - (m.GridCols * 18) - 10
		if rightWidth < 30 {
			rightWidth = 30
		}
		rightHeight := (m.GridRows * 5)
		if rightHeight < 6 {
			rightHeight = 6
		}

		if !m.ViewportReady {
			m.Viewport = viewport.New(rightWidth, rightHeight)
			m.ViewportReady = true
		} else {
			m.Viewport.Width = rightWidth
			m.Viewport.Height = rightHeight
		}

		return m, nil

	case fox.TickMsg:
		if m.IsTesting {
			m.FoxFrame++
			return m, fox.TickCmd()
		}
		return m, nil

	case runner.TestResultMsg:
		m.IsTesting = false
		m.TestOutput = msg.Output
		m.StatusMsg = ""
		if msg.Passed {
			m.Progress.Passed[msg.ExerciseID] = true
		}
		m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
		manifest.SaveProgress(m.RootDir, m.Progress)
		return m, nil

	case ide.ResultMsg:
		m.StatusMsg = "IDE opened: " + msg.IDE
		if msg.Dir != "" {
			m.StatusMsg += " → " + msg.Dir
		}
		if msg.Err != nil {
			m.StatusMsg = "Failed to open IDE: " + msg.Err.Error()
		}
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)
	}

	var vpCmd tea.Cmd
	m.Viewport, vpCmd = m.Viewport.Update(msg)
	cmds = append(cmds, vpCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.ShowIdePicker {
		return m.handleIdePickerKey(msg)
	}

	switch msg.String() {

	case "q", "ctrl+c", "esc":
		return m, tea.Quit

	case "o", "O":
		return m.handleOpenInIDE()

	case "h", "H":
		m.ShowHint = !m.ShowHint

	case "tab":
		m.FocusedPanel = (m.FocusedPanel + 1) % 4

	case "shift+tab":
		m.FocusedPanel = (m.FocusedPanel + 3) % 4

	case "r", "enter":
		switch m.FocusedPanel {
		case PanelChapters:
			m.FocusedPanel = PanelLessons
		case PanelLessons:
			m.FocusedPanel = PanelTasks
		case PanelTasks:
			if !m.IsTesting {
				m.IsTesting = true
				m.ShowHint = false
				m.TestOutput = ""
				m.StatusMsg = ""
				curEx := m.Exercises[m.FocusedIdx]
				return m, tea.Batch(
					fox.TickCmd(),
					runner.RunTestCmd(m.RootDir, curEx),
				)
			}
		case PanelDetails:
			m.FocusedPanel = PanelTasks
		}

	case "up":
		m.moveUp()

	case "down":
		m.moveDown()

	case "left":
		m.moveLeft()

	case "right":
		m.moveRight()

	case "n", "N":
		if m.FocusedIdx < m.lessonEndIdx()-1 {
			m.focusTask(m.FocusedIdx + 1)
		}

	case "p", "P":
		if m.FocusedIdx > m.lessonStartIdx() {
			m.focusTask(m.FocusedIdx - 1)
		}
	}

	return m, nil
}

func (m *Model) moveUp() {
	switch m.FocusedPanel {
	case PanelChapters:
		if m.SelectedChapter > 0 {
			m.focusLesson(m.SelectedChapter-1, 0)
		}
	case PanelLessons:
		if m.SelectedLesson > 0 {
			m.focusLesson(m.SelectedChapter, m.SelectedLesson-1)
		}
	case PanelTasks:
		start := m.lessonStartIdx()
		if m.FocusedIdx-m.GridCols >= start {
			m.focusTask(m.FocusedIdx - m.GridCols)
		}
	}
}

func (m *Model) moveDown() {
	switch m.FocusedPanel {
	case PanelChapters:
		if m.SelectedChapter < len(m.Manifest.Categories)-1 {
			m.focusLesson(m.SelectedChapter+1, 0)
		}
	case PanelLessons:
		if m.SelectedLesson < len(m.currentCategory().Topics)-1 {
			m.focusLesson(m.SelectedChapter, m.SelectedLesson+1)
		}
	case PanelTasks:
		end := m.lessonEndIdx()
		if m.FocusedIdx+m.GridCols < end {
			m.focusTask(m.FocusedIdx + m.GridCols)
		}
	}
}

func (m *Model) moveLeft() {
	switch m.FocusedPanel {
	case PanelChapters:
		// stay
	case PanelLessons:
		m.FocusedPanel = PanelChapters
	case PanelTasks:
		start := m.lessonStartIdx()
		if m.FocusedIdx > start && m.FocusedIdx%m.GridCols > 0 {
			m.focusTask(m.FocusedIdx - 1)
		} else {
			m.FocusedPanel = PanelLessons
		}
	case PanelDetails:
		m.FocusedPanel = PanelTasks
	}
}

func (m *Model) moveRight() {
	switch m.FocusedPanel {
	case PanelChapters:
		m.FocusedPanel = PanelLessons
	case PanelLessons:
		m.FocusedPanel = PanelTasks
	case PanelTasks:
		end := m.lessonEndIdx()
		if m.FocusedIdx < end-1 && (m.FocusedIdx+1)%m.GridCols != 0 {
			m.focusTask(m.FocusedIdx + 1)
		} else {
			m.FocusedPanel = PanelDetails
		}
	case PanelDetails:
		m.FocusedPanel = PanelChapters
	}
}

func (m Model) handleOpenInIDE() (tea.Model, tea.Cmd) {
	if len(m.DetectedIDEs) == 0 {
		m.ShowIdePicker = true
		m.IdeCustom = true
		m.ideInput.Focus()
		return m, textinput.Blink
	}
	if m.Progress.PreferredIDE != "" {
		return m, m.openInIDECmd()
	}
	m.ShowIdePicker = true
	m.IdePickerIdx = 0
	return m, nil
}

func (m Model) handleIdePickerKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.IdeCustom {
		if msg.Type == tea.KeyEnter {
			value := m.ideInput.Value()
			m.IdeCustom = false
			m.ideInput.Blur()
			if value == "" {
				m.ShowIdePicker = false
				return m, nil
			}
			m.Progress.PreferredIDE = value
			manifest.SaveProgress(m.RootDir, m.Progress)
			m.ShowIdePicker = false
			return m, m.openInIDECmd()
		}
		if msg.String() == "esc" {
			m.IdeCustom = false
			m.ideInput.Blur()
			return m, nil
		}
		var tiCmd tea.Cmd
		m.ideInput, tiCmd = m.ideInput.Update(msg)
		return m, tiCmd
	}

	switch msg.String() {
	case "up":
		if m.IdePickerIdx > 0 {
			m.IdePickerIdx--
		}
	case "down":
		if m.IdePickerIdx < len(m.DetectedIDEs) {
			m.IdePickerIdx++
		}
	case "enter":
		if m.IdePickerIdx >= len(m.DetectedIDEs) {
			m.IdeCustom = true
			m.ideInput.Focus()
			return m, textinput.Blink
		}
		choice := m.DetectedIDEs[m.IdePickerIdx]
		m.Progress.PreferredIDE = choice.Key
		manifest.SaveProgress(m.RootDir, m.Progress)
		m.ShowIdePicker = false
		return m, m.openInIDECmd()
	case "esc":
		m.ShowIdePicker = false
	}
	return m, nil
}

func (m Model) openInIDECmd() tea.Cmd {
	return func() tea.Msg {
		if len(m.Exercises) == 0 {
			return ide.ResultMsg{Err: fmt.Errorf("no exercise selected")}
		}
		ex := m.Exercises[m.FocusedIdx]
		dir, err := ide.PrepareWorkspace(m.RootDir, ex, m.chapterNo(), m.currentTopic().Title)
		if err != nil {
			return ide.ResultMsg{Err: err}
		}

		key := m.Progress.PreferredIDE
		choice, ok := ide.FindByKey(m.DetectedIDEs, key)
		if !ok {
			choice = ide.IDE{Key: key, Command: key}
		}
		if err := ide.Launch(choice, dir); err != nil {
			return ide.ResultMsg{Dir: dir, IDE: key, Err: err}
		}
		return ide.ResultMsg{Dir: dir, IDE: choice.Name}
	}
}