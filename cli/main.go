package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Manifest data models
type Manifest struct {
	Title      string     `json:"title"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Topics []Topic `json:"topics"`
}

type Topic struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	ContentPath string     `json:"content_path"`
	Exercises   []Exercise `json:"exercises"`
}

type Exercise struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Level         string `json:"level"`
	FilePath      string `json:"file_path"`
	TestPath      string `json:"test_path"`
	DocPath       string `json:"doc_path"`
	CategoryTitle string `json:"-"`
	TopicTitle    string `json:"-"`
}

// Progress persistence model
type Progress struct {
	Passed map[string]bool `json:"passed"`
	LastID string          `json:"last_id"`
}

// ANSI colors
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

func findRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	for {
		manifestPath := filepath.Join(dir, "exercises", "manifest.json")
		if _, err := os.Stat(manifestPath); err == nil {
			return dir
		}
		manifestPathRoot := filepath.Join(dir, "manifest.json")
		if _, err := os.Stat(manifestPathRoot); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "."
}

func loadManifest(rootDir string) (*Manifest, []Exercise, error) {
	manifestPaths := []string{
		filepath.Join(rootDir, "exercises", "manifest.json"),
		filepath.Join(rootDir, "manifest.json"),
	}

	var data []byte
	var err error
	for _, p := range manifestPaths {
		data, err = os.ReadFile(p)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, nil, fmt.Errorf("could not find manifest.json: %w", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, nil, fmt.Errorf("failed to parse manifest.json: %w", err)
	}

	var flat []Exercise
	for _, cat := range manifest.Categories {
		for _, top := range cat.Topics {
			for _, ex := range top.Exercises {
				ex.CategoryTitle = cat.Title
				ex.TopicTitle = top.Title
				flat = append(flat, ex)
			}
		}
	}

	return &manifest, flat, nil
}

func loadProgress(rootDir string) Progress {
	path := filepath.Join(rootDir, ".gostlings_progress.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return Progress{Passed: make(map[string]bool)}
	}
	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		return Progress{Passed: make(map[string]bool)}
	}
	if p.Passed == nil {
		p.Passed = make(map[string]bool)
	}
	return p
}

func saveProgress(rootDir string, p Progress) {
	path := filepath.Join(rootDir, ".gostlings_progress.json")
	data, _ := json.MarshalIndent(p, "", "  ")
	_ = os.WriteFile(path, data, 0644)
}

func renderHeader(currentIdx, total int, passedCount int) {
	fmt.Print("\033[H\033[2J") // Clear terminal
	fmt.Printf("%s%s=== GOSTLINGS — Interactive Go Workshop ===%s\n", Bold, Cyan, Reset)

	// Render progress bar
	barLength := 30
	filled := 0
	if total > 0 {
		filled = (passedCount * barLength) / total
	}
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barLength-filled)

	fmt.Printf("Progress: [%s%s%s] %d/%d passed\n", Green, bar, Reset, passedCount, total)
	fmt.Println(strings.Repeat("─", 50))
}

func runTest(rootDir string, ex Exercise) (bool, string) {
	fullFilePath := filepath.Join(rootDir, ex.FilePath)
	fullTestPath := filepath.Join(rootDir, ex.TestPath)

	// Check if exercise file exists
	if _, err := os.Stat(fullFilePath); err != nil {
		return false, fmt.Sprintf("Exercise file not found: %s", ex.FilePath)
	}

	// Check if test file exists
	if _, err := os.Stat(fullTestPath); err != nil {
		// Fallback: run 'go run <file>' to check compilation
		cmd := exec.Command("go", "run", fullFilePath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return false, fmt.Sprintf("Compilation/Execution failed:\n%s", string(out))
		}
		return true, fmt.Sprintf("Code compiled and executed cleanly! Output:\n%s", string(out))
	}

	// Run 'go test' with both exercise file and test file
	dir := filepath.Dir(fullFilePath)
	baseFile := filepath.Base(fullFilePath)
	baseTest := filepath.Base(fullTestPath)

	cmd := exec.Command("go", "test", "-v", baseFile, baseTest)
	cmd.Dir = dir

	out, err := cmd.CombinedOutput()
	outputStr := string(out)

	if err != nil {
		return false, outputStr
	}
	return true, outputStr
}

func renderDoc(rootDir string, docPath string) {
	fullPath := filepath.Join(rootDir, docPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("%sNo explanation found at %s%s\n", Red, docPath, Reset)
		return
	}
	fmt.Printf("\n%s=== Exercise Explanation & Hints ===%s\n", Yellow, Reset)
	fmt.Println(string(data))
	fmt.Println(strings.Repeat("─", 50))
}

func renderList(exercises []Exercise, progress Progress) {
	fmt.Printf("\n%s=== Exercise List ===%s\n", Bold, Reset)
	for i, ex := range exercises {
		status := fmt.Sprintf("%s[ ]%s", Yellow, Reset)
		if progress.Passed[ex.ID] {
			status = fmt.Sprintf("%s[✓]%s", Green, Reset)
		}
		fmt.Printf("%2d. %s %s > %s > %s (%s) [%s]\n",
			i+1, status, ex.CategoryTitle, ex.TopicTitle, ex.Title, ex.Level, ex.FilePath)
	}
	fmt.Println(strings.Repeat("─", 50))
}

func main() {
	rootDir := findRootDir()
	_, exercises, err := loadManifest(rootDir)
	if err != nil {
		fmt.Printf("%sError loading manifest: %v%s\n", Red, err, Reset)
		os.Exit(1)
	}

	if len(exercises) == 0 {
		fmt.Printf("%sNo exercises found in manifest.%s\n", Red, Reset)
		os.Exit(1)
	}

	progress := loadProgress(rootDir)

	// Resume from last exercise or first pending
	currentIdx := 0
	if progress.LastID != "" {
		for i, ex := range exercises {
			if ex.ID == progress.LastID {
				currentIdx = i
				break
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		passedCount := 0
		for _, ex := range exercises {
			if progress.Passed[ex.ID] {
				passedCount++
			}
		}

		renderHeader(currentIdx, len(exercises), passedCount)
		curEx := exercises[currentIdx]

		fmt.Printf("%sCurrent Exercise [%d/%d]:%s %s%s%s (%s)\n",
			Bold, currentIdx+1, len(exercises), Reset, Cyan, curEx.Title, Reset, curEx.Level)
		fmt.Printf("Category: %s | Topic: %s\n", curEx.CategoryTitle, curEx.TopicTitle)
		fmt.Printf("File:     %s%s%s\n", Bold, curEx.FilePath, Reset)

		statusStr := fmt.Sprintf("%s[ PENDING ]%s", Yellow, Reset)
		if progress.Passed[curEx.ID] {
			statusStr = fmt.Sprintf("%s[ ✓ PASSED ]%s", Green, Reset)
		}
		fmt.Printf("Status:   %s\n\n", statusStr)

		fmt.Printf("%sOptions:%s [r] Run test  [n] Next  [p] Prev  [l] List  [h] Hint  [q] Quit\n", Bold, Reset)
		fmt.Print("Enter command: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		cmd := strings.ToLower(strings.TrimSpace(input))

		switch {
		case cmd == "q" || cmd == "quit":
			fmt.Println("Goodbye! Keep coding in Go!")
			return

		case cmd == "n" || cmd == "next":
			if currentIdx < len(exercises)-1 {
				currentIdx++
				progress.LastID = exercises[currentIdx].ID
				saveProgress(rootDir, progress)
			} else {
				fmt.Printf("%sAlready at the last exercise!%s\n", Yellow, Reset)
			}

		case cmd == "p" || cmd == "prev" || cmd == "previous":
			if currentIdx > 0 {
				currentIdx--
				progress.LastID = exercises[currentIdx].ID
				saveProgress(rootDir, progress)
			} else {
				fmt.Printf("%sAlready at the first exercise!%s\n", Yellow, Reset)
			}

		case cmd == "l" || cmd == "list":
			renderList(exercises, progress)
			fmt.Print("\nEnter exercise number to jump to (or Press Enter to return): ")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			if numStr != "" {
				var target int
				if _, err := fmt.Sscanf(numStr, "%d", &target); err == nil && target >= 1 && target <= len(exercises) {
					currentIdx = target - 1
					progress.LastID = exercises[currentIdx].ID
					saveProgress(rootDir, progress)
				}
			}

		case cmd == "h" || cmd == "hint":
			renderDoc(rootDir, curEx.DocPath)
			fmt.Print("\nPress Enter to continue...")
			_, _ = reader.ReadString('\n')

		case cmd == "r" || cmd == "run" || cmd == "":
			fmt.Printf("\n%sRunning test for %s...%s\n\n", Cyan, curEx.FilePath, Reset)
			passed, output := runTest(rootDir, curEx)
			if passed {
				fmt.Printf("%s✓ SUCCESS! Exercise completed!%s\n\n", Green, Reset)
				fmt.Println(output)
				progress.Passed[curEx.ID] = true
				saveProgress(rootDir, progress)

				if currentIdx < len(exercises)-1 {
					fmt.Print("Press Enter to advance to next exercise...")
					_, _ = reader.ReadString('\n')
					currentIdx++
					progress.LastID = exercises[currentIdx].ID
					saveProgress(rootDir, progress)
				}
			} else {
				fmt.Printf("%s✗ FAILED! Keep trying! Output:%s\n\n", Red, Reset)
				fmt.Println(output)
				fmt.Print("\nPress Enter to return to menu...")
				_, _ = reader.ReadString('\n')
			}
		default:
			// Check if user entered an exercise number directly
			var target int
			if _, err := fmt.Sscanf(cmd, "%d", &target); err == nil && target >= 1 && target <= len(exercises) {
				currentIdx = target - 1
				progress.LastID = exercises[currentIdx].ID
				saveProgress(rootDir, progress)
			}
		}
	}
}
