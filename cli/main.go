package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Mozilla-Campus-Club-of-SLIIT/codenight-n3-GO/ui"
)

func main() {
	model, err := ui.NewModel()
	if err != nil {
		fmt.Printf("Error initializing gostlings CLI: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running gostlings TUI: %v\n", err)
		os.Exit(1)
	}
}
