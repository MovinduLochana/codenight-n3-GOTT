package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/fox"
	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/runner"
)

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
	ShowHint      bool
	WindowWidth   int
	WindowHeight  int
	Viewport      viewport.Model
	ViewportReady bool
}

func NewModel() (Model, error) {
	rootDir := manifest.FindRootDir()
	m, exercises, err := manifest.LoadManifest(rootDir)
	if err != nil {
		return Model{}, err
	}
	progress := manifest.LoadProgress(rootDir)

	focusedIdx := 0
	if progress.LastID != "" {
		for i, ex := range exercises {
			if ex.ID == progress.LastID {
				focusedIdx = i
				break
			}
		}
	}

	return Model{
		RootDir:    rootDir,
		Manifest:   m,
		Exercises:  exercises,
		Progress:   progress,
		FocusedIdx: focusedIdx,
		GridCols:   4,
		GridRows:   3,
		PageSize:   12,
		ShowHint:   false,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) recalculateLayout() {
	if m.WindowWidth == 0 || m.WindowHeight == 0 {
		return
	}

	// Calculate GridCols based on terminal width
	// Each card is 18 chars wide (16 content + 2 border)
	availWidth := m.WindowWidth - 6
	leftCols := (availWidth * 55) / 100 / 18
	if leftCols < 2 {
		leftCols = 2
	}
	if leftCols > 5 {
		leftCols = 5
	}
	m.GridCols = leftCols

	// Calculate GridRows based on terminal height
	// Available height overhead: 1 title header + 2 footer border = 3 lines
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

		rightWidth := m.WindowWidth - (m.GridCols * 18) - 10
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
		if msg.Passed {
			m.Progress.Passed[msg.ExerciseID] = true
		}
		m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
		manifest.SaveProgress(m.RootDir, m.Progress)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "r", "enter":
			if !m.IsTesting && len(m.Exercises) > 0 {
				m.IsTesting = true
				m.ShowHint = false
				m.TestOutput = ""
				curEx := m.Exercises[m.FocusedIdx]
				return m, tea.Batch(
					fox.TickCmd(),
					runner.RunTestCmd(m.RootDir, curEx),
				)
			}

		case "h", "H":
			m.ShowHint = !m.ShowHint

		case "n", "N":
			if m.FocusedIdx < len(m.Exercises)-1 {
				m.FocusedIdx++
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}

		case "p", "P":
			if m.FocusedIdx > 0 {
				m.FocusedIdx--
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}

		case "left":
			if m.FocusedIdx%m.GridCols > 0 {
				m.FocusedIdx--
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}

		case "right":
			if (m.FocusedIdx+1)%m.GridCols != 0 && m.FocusedIdx < len(m.Exercises)-1 {
				m.FocusedIdx++
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}

		case "up":
			if m.FocusedIdx >= m.GridCols {
				m.FocusedIdx -= m.GridCols
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}

		case "down":
			if m.FocusedIdx+m.GridCols < len(m.Exercises) {
				m.FocusedIdx += m.GridCols
				m.Progress.LastID = m.Exercises[m.FocusedIdx].ID
				manifest.SaveProgress(m.RootDir, m.Progress)
			}
		}
	}

	var vpCmd tea.Cmd
	m.Viewport, vpCmd = m.Viewport.Update(msg)
	cmds = append(cmds, vpCmd)

	return m, tea.Batch(cmds...)
}
