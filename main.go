package main

import (
	"fmt"
	"log"
	"os"

	"terbox/internal/data"
	"terbox/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Load configuration
	config, err := data.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		config = data.DefaultConfig()
	}

	// Create the application
	app := ui.NewApp(config)

	// Create Bubble Tea program
	p := tea.NewProgram(app, tea.WithAltScreen(), tea.WithMouseCellMotion())

	// Run the application
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}
