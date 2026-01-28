package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Browser represents a browser-like interface with tabs and content panels
type Browser struct {
	tabs        *Tabs
	panel       *Panel
	terminal    *Terminal
	width       int
	height      int
	focusedOn   string // "tabs" or "content"
	contentMode string // "panel" or "terminal"
	theme       *Theme
}

// NewBrowser creates a new browser with initial tabs and default theme
func NewBrowser(initialTabs []Tab) *Browser {
	return NewBrowserWithTheme(initialTabs, DefaultTheme())
}

// NewBrowserWithTheme creates a new browser with initial tabs and custom theme
func NewBrowserWithTheme(initialTabs []Tab, theme *Theme) *Browser {
	tabs := NewTabsWithTheme(initialTabs, theme)
	panel := NewPanel("")
	terminal := NewTerminalWithTheme(theme)

	return &Browser{
		tabs:        tabs,
		panel:       panel,
		terminal:    terminal,
		focusedOn:   "tabs",
		contentMode: "panel",
		theme:       theme,
	}
}

// SetSize sets the dimensions of the browser
func (b *Browser) SetSize(width, height int) {
	b.width = width
	b.height = height

	// Tabs is 1 line + 2 separator lines
	tabsHeight := 6

	// Content area gets the rest of the space
	contentHeight := height - tabsHeight

	b.tabs.SetSize(width)
	b.panel.SetSize(width, contentHeight)
	b.terminal.SetSize(width, contentHeight)
}

// Init initializes the browser
func (b *Browser) Init() tea.Cmd {
	return tea.Batch(b.tabs.Init(), b.panel.Init(), b.terminal.Init())
}

// Update handles input for the browser
func (b *Browser) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if b.focusedOn == "tabs" {
				b.focusedOn = "panel"
			} else {
				b.focusedOn = "tabs"
			}
			return nil
		case "ctrl+t":
			// Create a new tab
			b.tabs.AddTab(Tab{
				Title:   fmt.Sprintf("Tab %d", len(b.tabs.GetTabs())+1),
				Content: "New tab content",
			})
			return nil
		case "ctrl+w":
			// Close current tab
			if len(b.tabs.GetTabs()) > 0 {
				b.tabs.RemoveTab(b.tabs.ActiveIndex())
			}
			return nil
		case "ctrl+l":
			// Toggle between panel and terminal modes
			if b.contentMode == "panel" {
				b.contentMode = "terminal"
			} else {
				b.contentMode = "panel"
			}
			return nil
		}
	}

	if b.focusedOn == "tabs" {
		b.tabs.Update(msg)
		// Update panel content based on active tab
		activeTab := b.tabs.ActiveTab()
		b.panel.SetContent(activeTab.Content)
	} else {
		// Content area is focused
		if b.contentMode == "panel" {
			b.panel.Update(msg)
		} else {
			b.terminal.Update(msg)
		}
	}

	return nil
}

// View renders the browser
func (b *Browser) View() string {
	if b.width == 0 || b.height == 0 {
		return ""
	}

	// Top separator
	topSeparator := strings.Repeat("─", b.width)
	if b.theme != nil {
		topSeparator = lipgloss.NewStyle().
			Foreground(lipgloss.Color(b.theme.GetSeparatorColor())).
			Render(topSeparator)
	}

	// Render tabs
	tabsView := b.tabs.View()

	// Apply focus styling to tabs
	if b.focusedOn == "tabs" && b.theme != nil {
		tabsView = b.theme.GetTabFocusedStyle().Render(tabsView)
	}

	// Bottom separator
	bottomSeparator := strings.Repeat("─", b.width)
	if b.theme != nil {
		bottomSeparator = lipgloss.NewStyle().
			Foreground(lipgloss.Color(b.theme.GetSeparatorColor())).
			Render(bottomSeparator)
	}

	// Render content (panel or terminal)
	var contentView string
	if b.contentMode == "panel" {
		activeTab := b.tabs.ActiveTab()
		b.panel.SetContent(activeTab.Content)
		contentView = b.panel.View()
	} else {
		contentView = b.terminal.View()
	}

	// Combine all parts
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		topSeparator,
		tabsView,
		bottomSeparator,
		contentView,
	)

	return content
}

// AddTab adds a new tab to the browser
func (b *Browser) AddTab(tab Tab) {
	b.tabs.AddTab(tab)
}

// RemoveTab removes a tab from the browser
func (b *Browser) RemoveTab(index int) {
	b.tabs.RemoveTab(index)
}

// GetActiveTabIndex returns the index of the active tab
func (b *Browser) GetActiveTabIndex() int {
	return b.tabs.ActiveIndex()
}

// UpdateTabContent updates the content of a specific tab
func (b *Browser) UpdateTabContent(index int, content string) {
	b.tabs.UpdateTabContent(index, content)
}

// SetTheme sets the theme for the browser and all its components
func (b *Browser) SetTheme(theme *Theme) {
	b.theme = theme
	b.tabs.SetTheme(theme)
	b.panel.SetTheme(theme)
	b.terminal.SetTheme(theme)
}

// GetTheme returns the current theme
func (b *Browser) GetTheme() *Theme {
	return b.theme
}

// SetContentMode sets the content display mode ("panel" or "terminal")
func (b *Browser) SetContentMode(mode string) {
	if mode == "panel" || mode == "terminal" {
		b.contentMode = mode
	}
}

// GetContentMode returns the current content mode
func (b *Browser) GetContentMode() string {
	return b.contentMode
}

// GetTerminal returns the terminal component
func (b *Browser) GetTerminal() *Terminal {
	return b.terminal
}

// GetPanel returns the panel component
func (b *Browser) GetPanel() *Panel {
	return b.panel
}
