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
	theme   *Theme
}

// NewPanel creates a new panel with the given title and default theme
func NewPanel(title string) *Panel {
	return NewPanelWithTheme(title, DefaultTheme())
}

// NewPanelWithTheme creates a new panel with the given title and custom theme
func NewPanelWithTheme(title string, theme *Theme) *Panel {
	return &Panel{
		title: title,
		style: theme.GetPanelStyle(),
		theme: theme,
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

// SetTheme sets the theme for the panel
func (p *Panel) SetTheme(theme *Theme) {
	p.theme = theme
	p.style = theme.GetPanelStyle()
}

// GetTheme returns the current theme
func (p *Panel) GetTheme() *Theme {
	return p.theme
}
