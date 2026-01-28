package ui

import tea "github.com/charmbracelet/bubbletea"

// Component is the base interface for all UI components
type Component interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	View() string
}

// Dimensions holds width and height for a component
type Dimensions struct {
	Width  int
	Height int
}

// Position holds x and y coordinates
type Position struct {
	X int
	Y int
}
