package main

import "testing"

func TestSign(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{42, "positive"},
		{-3, "negative"},
		{0, "zero"},
	}
	for _, tt := range tests {
		if got := Sign(tt.n); got != tt.want {
			t.Errorf("Sign(%d) = %q; want %q", tt.n, got, tt.want)
		}
	}
}
