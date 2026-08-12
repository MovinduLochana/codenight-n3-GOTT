package main

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := strings.TrimSpace(buf.String())
	lines := strings.Split(output, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	if len(lines) < 3 {
		t.Errorf("Expected at least 3 lines of output, got %d. Output:\n%s", len(lines), output)
		return
	}

	// Verify "Colors: [Red Green Blue]"
	expectedColors := "Colors: [Red Green Blue]"
	if lines[0] != expectedColors {
		t.Errorf("Expected line 1 to be %q, got %q", expectedColors, lines[0])
	}

	// Verify "Length: 3"
	expectedLength := "Length: 3"
	if lines[1] != expectedLength {
		t.Errorf("Expected line 2 to be %q, got %q", expectedLength, lines[1])
	}

	// Verify "Capacity: <cap>" (cap >= 3)
	if !strings.HasPrefix(lines[2], "Capacity: ") {
		t.Errorf("Expected line 3 to start with 'Capacity: ', got %q", lines[2])
	} else {
		capStr := strings.TrimPrefix(lines[2], "Capacity: ")
		capVal, err := strconv.Atoi(capStr)
		if err != nil {
			t.Errorf("Failed to parse capacity value %q: %v", capStr, err)
		} else if capVal < 3 {
			t.Errorf("Expected capacity to be at least 3, got %d", capVal)
		}
	}
}
