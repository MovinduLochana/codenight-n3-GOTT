package ide

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
)

// IDE represents an editor or IDE that can open a directory.
type IDE struct {
	Name    string // human readable label shown in the picker
	Key     string // stable identifier persisted in progress
	Command string // executable (or command line) to run, may be empty for macOS apps
	AppName string // macOS .app bundle name, empty otherwise
}

// ResultMsg is sent back to the TUI after an open-in-IDE attempt.
type ResultMsg struct {
	Dir      string
	IDE      string
	Finished bool // true when the IDE ran in the terminal and has now exited
	Err      error
}

var knownCLIs = []IDE{
	{Name: "VS Code", Key: "code", Command: "code"},
	{Name: "VS Code (Insiders)", Key: "code-insiders", Command: "code-insiders"},
	{Name: "Cursor", Key: "cursor", Command: "cursor"},
	{Name: "Windsurf", Key: "windsurf", Command: "windsurf"},
	{Name: "VSCodium", Key: "codium", Command: "codium"},
	{Name: "Zed", Key: "zed", Command: "zed"},
	{Name: "GoLand", Key: "goland", Command: "goland"},
	{Name: "IntelliJ IDEA", Key: "idea", Command: "idea"},
	{Name: "Sublime Text", Key: "subl", Command: "subl"},
	{Name: "Vim", Key: "vim", Command: "vim"},
	{Name: "Neovim", Key: "nvim", Command: "nvim"},
}

var macApps = []IDE{
	{Name: "VS Code", Key: "app:vscode", AppName: "Visual Studio Code"},
	{Name: "Cursor", Key: "app:cursor", AppName: "Cursor"},
	{Name: "Windsurf", Key: "app:windsurf", AppName: "Windsurf"},
	{Name: "GoLand", Key: "app:goland", AppName: "GoLand"},
	{Name: "IntelliJ IDEA", Key: "app:idea", AppName: "IntelliJ IDEA"},
	{Name: "Zed", Key: "app:zed", AppName: "Zed"},
}

// Detect returns the list of editors/IDEs available on this machine.
func Detect() []IDE {
	seen := map[string]bool{}
	var out []IDE

	add := func(e IDE) {
		key := e.Key
		if key == "" {
			key = e.Command
		}
		if key == "" || seen[key] {
			return
		}
		seen[key] = true
		out = append(out, e)
	}

	if v := os.Getenv("GOSTLINGS_IDE"); v != "" {
		add(IDE{Name: v, Key: "env:gostlings_ide", Command: v})
	}

	for _, e := range knownCLIs {
		if p, err := exec.LookPath(e.Command); err == nil {
			e.Command = p
			add(e)
		}
	}

	for _, env := range []string{"VISUAL", "EDITOR"} {
		if v := os.Getenv(env); v != "" {
			parts := strings.Fields(v)
			if len(parts) > 0 {
				if p, err := exec.LookPath(parts[0]); err == nil {
					add(IDE{Name: fmt.Sprintf("%s (%s)", env, v), Key: "env:" + parts[0], Command: p})
				}
			}
		}
	}

	if runtime.GOOS == "darwin" {
		for _, app := range macApps {
			if appInstalled(app.AppName) {
				add(app)
			}
		}
	}

	sort.SliceStable(out, func(i, j int) bool { return out[i].Name < out[j].Name })

	return out
}

// FindByKey returns the IDE in list whose Key matches key.
func FindByKey(list []IDE, key string) (IDE, bool) {
	for _, e := range list {
		if e.Key == key {
			return e, true
		}
	}
	return IDE{}, false
}

// IsTerminal reports whether the IDE runs inside the terminal (e.g. vim,
// nvim) as opposed to opening its own GUI window.
func (e IDE) IsTerminal() bool {
	return isTerminalCommand(e.Command) || (e.AppName == "" && isTerminalCommand(e.Name))
}

// CommandForDir returns the exec.Cmd used to open dir in the given IDE.
// Callers may either Start() it (GUI IDEs) or run it in the foreground while
// suspending the TUI (terminal editors).
func (e IDE) CommandForDir(dir string) (*exec.Cmd, error) {
	if runtime.GOOS == "darwin" && e.AppName != "" {
		return exec.Command("open", "-a", e.AppName, dir), nil
	}

	cmdLine := e.Command
	if cmdLine == "" {
		return nil, fmt.Errorf("no IDE command configured for %q", e.Name)
	}

	args := strings.Fields(cmdLine)
	if len(args) == 0 {
		return nil, fmt.Errorf("no IDE command configured for %q", e.Name)
	}

	if _, err := os.Stat(cmdLine); err == nil || len(args) == 1 {
		return exec.Command(args[0], dir), nil
	}
	return exec.Command(args[0], append(args[1:], dir)...), nil
}

// Launch opens dir in the given IDE without blocking the caller.
func Launch(e IDE, dir string) error {
	cmd, err := e.CommandForDir(dir)
	if err != nil {
		return err
	}
	detach(cmd)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to launch %q: %w", e.Command, err)
	}
	return nil
}

// PrepareWorkspace creates .gostlings-workspace/<chapter no> - <lesson name>/
// containing only the files relevant to the given task.
func PrepareWorkspace(rootDir string, ex manifest.Exercise, chapterNo int, lessonTitle string) (string, error) {
	wsRoot := filepath.Join(rootDir, ".gostlings-workspace")
	if err := os.MkdirAll(wsRoot, 0o755); err != nil {
		return "", err
	}

	dir := filepath.Join(wsRoot, fmt.Sprintf("%d - %s", chapterNo, sanitizeName(lessonTitle)))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}

	entries, err := os.ReadDir(dir)
	if err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				_ = os.Remove(filepath.Join(dir, entry.Name()))
			}
		}
	}

	type srcDest struct{ src, dest string }
	sources := []srcDest{
		{filepath.Join(rootDir, ex.FilePath), filepath.Join(dir, filepath.Base(ex.FilePath))},
	}
	if ex.TestPath != "" {
		sources = append(sources, srcDest{filepath.Join(rootDir, ex.TestPath), filepath.Join(dir, filepath.Base(ex.TestPath))})
	}
	if ex.DocPath != "" {
		sources = append(sources, srcDest{filepath.Join(rootDir, ex.DocPath), filepath.Join(dir, "instructions.md")})
	}

	for _, s := range sources {
		if _, err := os.Stat(s.src); err != nil {
			continue
		}
		if err := linkOrCopy(s.src, s.dest); err != nil {
			return dir, err
		}
	}

	return dir, nil
}

// linkOrCopy links src to dest (hard link first, then symlink) falling back
// to a plain copy. Hard links give write-through on all OSes without needing
// admin rights on Windows.
func linkOrCopy(src, dest string) error {
	if err := os.Link(src, dest); err == nil {
		return nil
	}
	if err := os.Symlink(src, dest); err == nil {
		return nil
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, data, 0o644)
}

func sanitizeName(s string) string {
	replacer := strings.NewReplacer(
		"/", "-", "\\", "-", ":", "-", "*", "-",
		"?", "-", `"`, "-", "<", "-", ">", "-", "|", "-",
	)
	return strings.TrimSpace(replacer.Replace(s))
}

func isTerminalCommand(cmdLine string) bool {
	fields := strings.Fields(cmdLine)
	if len(fields) == 0 {
		return false
	}
	base := strings.ToLower(filepath.Base(fields[0]))
	base = strings.TrimSuffix(base, ".exe")
	switch base {
	case "vim", "vi", "nvim", "nano", "emacs", "micro",
		"hx", "helix", "kak", "kakoune", "ne", "mg", "jed":
		return true
	}
	return false
}

func appInstalled(appName string) bool {
	paths := []string{
		filepath.Join("/Applications", appName+".app"),
		filepath.Join(os.Getenv("HOME"), "Applications", appName+".app"),
	}
	for _, p := range paths {
		if fi, err := os.Stat(p); err == nil && fi.IsDir() {
			return true
		}
	}
	return false
}
