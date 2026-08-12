package main

import (
	"bytes"
	"io"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestWorker(t *testing.T) {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	// Send jobs
	jobs <- 10
	jobs <- 20
	jobs <- 30
	close(jobs)

	// Run worker
	worker(1, jobs, results)

	// Read results
	res := []int{<-results, <-results, <-results}
	sort.Ints(res)

	want := []int{20, 40, 60}
	for i, v := range want {
		if res[i] != v {
			t.Errorf("worker results mismatch: got %v, want %v", res, want)
			break
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

	expected := "Results: [2 4 6 8 10]"
	if output != expected {
		t.Errorf("Expected output %q, got %q", expected, output)
	}
}
