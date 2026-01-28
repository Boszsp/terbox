package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Panel displays content with borders and styling
type Panel struct {
	title   string
	content string
	width   int
	height  int
	style   lipgloss.Style
}

// NewPanel creates a new panel with the given title
func NewPanel(title string) *Panel {
	return &Panel{
		title: title,
		style: lipgloss.NewStyle(),
	}
}

// SetSize sets the dimensions of the panel
func (p *Panel) SetSize(width, height int) {
	p.width = width
	p.height = height
}

// SetContent sets the content of the panel
func (p *Panel) SetContent(content string) {
	p.content = content
}

// SetStyle sets the style of the panel
func (p *Panel) SetStyle(style lipgloss.Style) {
	p.style = style
}

// Init returns no command
func (p *Panel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the panel
func (p *Panel) Update(msg tea.Msg) tea.Cmd {
	return nil
}

// View renders the panel
func (p *Panel) View() string {
	contentHeight := p.height
	if contentHeight < 1 {
		contentHeight = 1
	}

	// Truncate or pad content lines
	lines := strings.Split(p.content, "\n")
	if len(lines) > contentHeight {
		lines = lines[:contentHeight]
	}
	for len(lines) < contentHeight {
		lines = append(lines, "")
	}

	content := strings.Join(lines, "\n")
	return content
}
