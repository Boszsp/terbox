package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Browser represents a browser-like interface with tabs and content panels
type Browser struct {
	tabs      *Tabs
	panel     *Panel
	width     int
	height    int
	focusedOn string // "tabs" or "panel"
	theme     *Theme
}

// NewBrowser creates a new browser with initial tabs and default theme
func NewBrowser(initialTabs []Tab) *Browser {
	return NewBrowserWithTheme(initialTabs, DefaultTheme())
}

// NewBrowserWithTheme creates a new browser with initial tabs and custom theme
func NewBrowserWithTheme(initialTabs []Tab, theme *Theme) *Browser {
	tabs := NewTabsWithTheme(initialTabs, theme)
	panel := NewPanel("")

	return &Browser{
		tabs:      tabs,
		panel:     panel,
		focusedOn: "tabs",
		theme:     theme,
	}
}

// SetSize sets the dimensions of the browser
func (b *Browser) SetSize(width, height int) {
	b.width = width
	b.height = height

	// Tabs is 1 line + 2 separator lines
	tabsHeight := 6

	// Panel gets the rest of the space
	panelHeight := height - tabsHeight

	b.tabs.SetSize(width)
	b.panel.SetSize(width, panelHeight)
}

// Init initializes the browser
func (b *Browser) Init() tea.Cmd {
	return tea.Batch(b.tabs.Init(), b.panel.Init())
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
		}
	}

	if b.focusedOn == "tabs" {
		b.tabs.Update(msg)
		// Update panel content based on active tab
		activeTab := b.tabs.ActiveTab()
		b.panel.SetContent(activeTab.Content)
	} else {
		b.panel.Update(msg)
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

	// Render panel
	activeTab := b.tabs.ActiveTab()
	b.panel.SetContent(activeTab.Content)
	panelView := b.panel.View()

	// Combine all parts
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		topSeparator,
		tabsView,
		bottomSeparator,
		panelView,
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
}

// GetTheme returns the current theme
func (b *Browser) GetTheme() *Theme {
	return b.theme
}
