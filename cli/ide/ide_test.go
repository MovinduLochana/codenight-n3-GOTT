package ide

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/manifest"
)

func TestPrepareWorkspace(t *testing.T) {
	root := t.TempDir()
	exDir := filepath.Join(root, "exercises", "fundamentals", "vars")
	if err := os.MkdirAll(exDir, 0o755); err != nil {
		t.Fatal(err)
	}
	docDir := filepath.Join(root, "content", "fundamentals", "vars")
	if err := os.MkdirAll(docDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"1.go", "2.go", "3.go", "1_test.go", "2_test.go"} {
		if err := os.WriteFile(filepath.Join(exDir, f), []byte("// "+f), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(docDir, "task_beginner.md"), []byte("# task"), 0o644); err != nil {
		t.Fatal(err)
	}

	ex := manifest.Exercise{
		ID:       "x1",
		FilePath: "exercises/fundamentals/vars/1.go",
		TestPath: "exercises/fundamentals/vars/1_test.go",
		DocPath:  "content/fundamentals/vars/task_beginner.md",
	}

	dir, err := PrepareWorkspace(root, ex, 1, "Variables")
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(root, ".gostlings-workspace", "1 - Variables")
	if dir != want {
		t.Fatalf("dir = %q, want %q", dir, want)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]bool{}
	for _, e := range entries {
		got[e.Name()] = true
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 files, got %v", got)
	}
	for _, f := range []string{"1.go", "1_test.go", "instructions.md"} {
		if !got[f] {
			t.Fatalf("missing %s in workspace: %v", f, got)
		}
	}

	// Write-through: editing workspace 1.go edits the repo file (hard/sym link).
	if err := os.WriteFile(filepath.Join(dir, "1.go"), []byte("// edited"), 0o644); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(exDir, "1.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "// edited" {
		t.Fatalf("repo file not write-through: %q", string(data))
	}

	// Switching to task 2 must drop task 1 files.
	ex.FilePath = "exercises/fundamentals/vars/2.go"
	ex.TestPath = "exercises/fundamentals/vars/2_test.go"
	ex.DocPath = ""
	if _, err := PrepareWorkspace(root, ex, 1, "Variables"); err != nil {
		t.Fatal(err)
	}
	entries, _ = os.ReadDir(dir)
	got = map[string]bool{}
	for _, e := range entries {
		got[e.Name()] = true
	}
	if got["1.go"] || !got["2.go"] || !got["2_test.go"] {
		t.Fatalf("stale files not cleaned: %v", got)
	}
}

func TestLaunch(t *testing.T) {
	if _, err := os.Stat("/bin/true"); err != nil {
		t.Skip("/bin/true not available")
	}
	if err := Launch(IDE{Name: "true", Command: "/bin/true"}, t.TempDir()); err != nil {
		t.Fatal(err)
	}
}