package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestUserRoleConstants(t *testing.T) {
	// Verify that the constants are declared and have the correct iota values.
	if int(Guest) != 0 {
		t.Errorf("Guest expected 0, got %d", Guest)
	}
	if int(Member) != 1 {
		t.Errorf("Member expected 1, got %d", Member)
	}
	if int(Admin) != 2 {
		t.Errorf("Admin expected 2, got %d", Admin)
	}
	if int(Owner) != 3 {
		t.Errorf("Owner expected 3, got %d", Owner)
	}
}

func TestGetPermission(t *testing.T) {
	tests := []struct {
		role UserRole
		want string
	}{
		{Guest, "Read Only"},
		{Member, "Read & Write"},
		{Admin, "Read, Write & Moderate"},
		{Owner, "Full Access"},
		{UserRole(99), "Unknown"},
		{UserRole(-1), "Unknown"},
	}

	for _, tt := range tests {
		got := GetPermission(tt.role)
		if got != tt.want {
			t.Errorf("GetPermission(%d) = %q; want %q", tt.role, got, tt.want)
		}
	}
}

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

	expected := []string{
		"Guest: Read Only",
		"Member: Read & Write",
		"Admin: Read, Write & Moderate",
		"Owner: Full Access",
	}

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines of output, got %d. Output was:\n%s", len(expected), len(lines), output)
		return
	}

	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("At line %d: expected %q, got %q", i+1, expected[i], line)
		}
	}
}
