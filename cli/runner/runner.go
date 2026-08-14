package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
)

type TestResultMsg struct {
	ExerciseID string
	Passed     bool
	Output     string
}

func RunTestCmd(rootDir string, ex manifest.Exercise) tea.Cmd {
	return func() tea.Msg {
		fullFilePath := filepath.Join(rootDir, ex.FilePath)
		fullTestPath := filepath.Join(rootDir, ex.TestPath)

		if _, err := os.Stat(fullFilePath); err != nil {
			return TestResultMsg{
				ExerciseID: ex.ID,
				Passed:     false,
				Output:     fmt.Sprintf("Exercise file not found: %s", ex.FilePath),
			}
		}

		if _, err := os.Stat(fullTestPath); err != nil {
			cmd := exec.Command("go", "run", fullFilePath)
			out, err := cmd.CombinedOutput()
			if err != nil {
				return TestResultMsg{
					ExerciseID: ex.ID,
					Passed:     false,
					Output:     fmt.Sprintf("Compilation/Execution failed:\n%s", string(out)),
				}
			}
			return TestResultMsg{
				ExerciseID: ex.ID,
				Passed:     true,
				Output:     fmt.Sprintf("Code compiled and executed cleanly!\nOutput:\n%s", string(out)),
			}
		}

		dir := filepath.Dir(fullFilePath)
		baseFile := filepath.Base(fullFilePath)
		baseTest := filepath.Base(fullTestPath)

		cmd := exec.Command("go", "test", "-v", baseFile, baseTest)
		cmd.Dir = dir

		out, err := cmd.CombinedOutput()
		outputStr := string(out)

		if err != nil {
			return TestResultMsg{
				ExerciseID: ex.ID,
				Passed:     false,
				Output:     outputStr,
			}
		}

		return TestResultMsg{
			ExerciseID: ex.ID,
			Passed:     true,
			Output:     outputStr,
		}
	}
}
