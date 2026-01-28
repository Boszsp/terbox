package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tab represents a single tab
type Tab struct {
	Title   string
	Content string
}

// TabBar manages multiple tabs with navigation
type TabBar struct {
	tabs          []Tab
	active        int
	width         int
	height        int
	style         lipgloss.Style
	activeStyle   lipgloss.Style
	inactiveStyle lipgloss.Style
	tabPositions  []int // Store the x position of each tab for mouse clicks
}

// NewTabBar creates a new tab bar with given tabs
func NewTabBar(tabs []Tab) *TabBar {
	return &TabBar{
		tabs:   tabs,
		active: 0,
		style:  lipgloss.NewStyle(),
		activeStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("63")).
			Padding(0, 1),
		inactiveStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// SetSize sets the dimensions of the tab bar
func (tb *TabBar) SetSize(width, height int) {
	tb.width = width
	tb.height = height
}

// Init returns no command
func (tb *TabBar) Init() tea.Cmd {
	return nil
}

// Update handles key inputs and mouse clicks for tab navigation
func (tb *TabBar) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right", "l":
			if tb.active < len(tb.tabs)-1 {
				tb.active++
			}
		case "left", "h":
			if tb.active > 0 {
				tb.active--
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// Number shortcuts: 1 = first tab, 2 = second tab, etc.
			tabNum := int(msg.String()[0]-'0') - 1
			if tabNum >= 0 && tabNum < len(tb.tabs) {
				tb.active = tabNum
			}
		}
	case tea.MouseMsg:
		// Handle mouse clicks on tabs
		if msg.Type == tea.MouseLeft {
			for i, pos := range tb.tabPositions {
				// Check if click is within the tab's x range
				if msg.X >= pos && msg.X < pos+len(tb.tabs[i].Title)+2 {
					tb.active = i
					break
				}
			}
		}
	}
	return nil
}

// View renders the tab bar
func (tb *TabBar) View() string {
	if len(tb.tabs) == 0 {
		return ""
	}

	var tabStrings []string
	tb.tabPositions = []int{} // Reset positions
	currentPos := 0

	for i, tab := range tb.tabs {
		tb.tabPositions = append(tb.tabPositions, currentPos)
		var renderedTab string
		if i == tb.active {
			renderedTab = tb.activeStyle.Render(tab.Title)
		} else {
			renderedTab = tb.inactiveStyle.Render(tab.Title)
		}
		tabStrings = append(tabStrings, renderedTab)
		currentPos += lipgloss.Width(renderedTab)
	}

	line := strings.Join(tabStrings, "")
	line = line + strings.Repeat(" ", max(0, tb.width-lipgloss.Width(line)))
	return line
}

// ActiveTab returns the currently active tab
func (tb *TabBar) ActiveTab() Tab {
	if tb.active >= 0 && tb.active < len(tb.tabs) {
		return tb.tabs[tb.active]
	}
	return Tab{}
}

// ActiveIndex returns the index of the active tab
func (tb *TabBar) ActiveIndex() int {
	return tb.active
}

// AddTab adds a new tab
func (tb *TabBar) AddTab(tab Tab) {
	tb.tabs = append(tb.tabs, tab)
}

// RemoveTab removes a tab at the given index
func (tb *TabBar) RemoveTab(index int) {
	if index >= 0 && index < len(tb.tabs) {
		tb.tabs = append(tb.tabs[:index], tb.tabs[index+1:]...)
		if tb.active >= len(tb.tabs) && tb.active > 0 {
			tb.active--
		}
	}
}

// UpdateTabContent updates the content of a tab
func (tb *TabBar) UpdateTabContent(index int, content string) {
	if index >= 0 && index < len(tb.tabs) {
		tb.tabs[index].Content = content
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
