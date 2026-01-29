package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Component is the base interface for all UI components
type Component interface {
	Init() tea.Cmd
	Update(msg tea.Msg) tea.Cmd
	View() string
	SetSize(width, height int)
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

// BaseComponent provides common functionality for components
type BaseComponent struct {
	Width  int
	Height int
}

// SetSize sets the size of the component
func (bc *BaseComponent) SetSize(width, height int) {
	bc.Width = width
	bc.Height = height
}

// Style constants
var (
	// Colors
	PrimaryColor   = lipgloss.Color("39")
	SecondaryColor = lipgloss.Color("245")
	ErrorColor     = lipgloss.Color("196")
	SuccessColor   = lipgloss.Color("46")

	// Styles
	BorderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(PrimaryColor)

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(PrimaryColor)

	TabActiveStyle = lipgloss.NewStyle().
			Background(PrimaryColor).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1)

	TabInactiveStyle = lipgloss.NewStyle().
				Foreground(SecondaryColor).
				Padding(0, 1)
)
