package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Browser represents a browser-like interface with tabs and content panels
type Browser struct {
	tabBar    *TabBar
	panel     *Panel
	width     int
	height    int
	focusedOn string // "tabs" or "panel"
}

// NewBrowser creates a new browser with initial tabs
func NewBrowser(initialTabs []Tab) *Browser {
	tabBar := NewTabBar(initialTabs)
	panel := NewPanel("")

	return &Browser{
		tabBar:    tabBar,
		panel:     panel,
		focusedOn: "tabs",
	}
}

// SetSize sets the dimensions of the browser
func (b *Browser) SetSize(width, height int) {
	b.width = width
	b.height = height

	// Tab bar is 1 line + 1 separator
	tabBarHeight := 1

	// Panel gets the rest of the space
	panelHeight := height - tabBarHeight - 2

	b.tabBar.SetSize(width, tabBarHeight)
	b.panel.SetSize(width, panelHeight)
}

// Init initializes the browser
func (b *Browser) Init() tea.Cmd {
	return tea.Batch(b.tabBar.Init(), b.panel.Init())
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
			b.tabBar.AddTab(Tab{
				Title:   fmt.Sprintf("Tab %d", len(b.tabBar.tabs)+1),
				Content: "New tab content",
			})
			return nil
		case "ctrl+w":
			// Close current tab
			if len(b.tabBar.tabs) > 0 {
				b.tabBar.RemoveTab(b.tabBar.ActiveIndex())
			}
			return nil
		}
	}

	if b.focusedOn == "tabs" {
		b.tabBar.Update(msg)
		// Update panel content based on active tab
		activeTab := b.tabBar.ActiveTab()
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

	// Render tab bar
	tabBarView := b.tabBar.View()

	// Separator line
	separator := strings.Repeat("─", b.width)

	// Render panel
	activeTab := b.tabBar.ActiveTab()
	b.panel.SetContent(activeTab.Content)
	panelView := b.panel.View()

	// Apply focus styling
	if b.focusedOn == "tabs" {
		tabBarView = lipgloss.NewStyle().
			Foreground(lipgloss.Color("228")).
			Render(tabBarView)
	}

	// Combine all parts
	content := lipgloss.JoinVertical(
		lipgloss.Top,
		tabBarView,
		separator,
		panelView,
	)

	return content
}

// AddTab adds a new tab to the browser
func (b *Browser) AddTab(tab Tab) {
	b.tabBar.AddTab(tab)
}

// RemoveTab removes a tab from the browser
func (b *Browser) RemoveTab(index int) {
	b.tabBar.RemoveTab(index)
}

// GetActiveTabIndex returns the index of the active tab
func (b *Browser) GetActiveTabIndex() int {
	return b.tabBar.ActiveIndex()
}

// UpdateTabContent updates the content of a specific tab
func (b *Browser) UpdateTabContent(index int, content string) {
	b.tabBar.UpdateTabContent(index, content)
}
