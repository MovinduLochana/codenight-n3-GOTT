package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO: Implement CountLines(filename string) (int, error).
// It should open the file, scan it line by line using bufio.NewScanner, and count the lines.
// Make sure to close the file properly.

func main() {
	count, err := CountLines("sample.txt")
	if err == nil {
		fmt.Printf("Line count: %d\n", count)
	}
}
