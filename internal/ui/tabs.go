package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tabs represents a scrollable horizontal tab component
type Tabs struct {
	tabs          []Tab
	active        int
	width         int
	scrollOffset  int // How many tabs to scroll left
	style         lipgloss.Style
	activeStyle   lipgloss.Style
	inactiveStyle lipgloss.Style
	tabPositions  []int // Store the x position of each tab for mouse clicks
}

// NewTabs creates a new scrollable tabs component
func NewTabs(tabs []Tab) *Tabs {
	return &Tabs{
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

// SetSize sets the width of the tabs component
func (t *Tabs) SetSize(width int) {
	t.width = width
	t.ensureActiveTabVisible()
}

// Init returns no command
func (t *Tabs) Init() tea.Cmd {
	return nil
}

// Update handles key inputs and mouse clicks for tab navigation
func (t *Tabs) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right", "l":
			if t.active < len(t.tabs)-1 {
				t.active++
				t.ensureActiveTabVisible()
			}
		case "left", "h":
			if t.active > 0 {
				t.active--
				t.ensureActiveTabVisible()
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			// Number shortcuts: 1 = first tab, 2 = second tab, etc.
			tabNum := int(msg.String()[0]-'0') - 1
			if tabNum >= 0 && tabNum < len(t.tabs) {
				t.active = tabNum
				t.ensureActiveTabVisible()
			}
		}
	case tea.MouseMsg:
		// Handle mouse clicks on tabs
		if msg.Type == tea.MouseLeft {
			for i, pos := range t.tabPositions {
				// Check if click is within the tab's x range
				if msg.X >= pos && msg.X < pos+len(t.tabs[i].Title)+2 {
					t.active = i
					t.ensureActiveTabVisible()
					break
				}
			}
		}
	}
	return nil
}

// View renders the scrollable tabs
func (t *Tabs) View() string {
	if len(t.tabs) == 0 {
		return ""
	}

	var tabStrings []string
	t.tabPositions = []int{} // Reset positions
	currentPos := 0

	// Calculate visible range based on scroll offset
	visibleTabs := t.getVisibleTabs()

	for i := visibleTabs.start; i < visibleTabs.end; i++ {
		tab := t.tabs[i]
		t.tabPositions = append(t.tabPositions, currentPos)

		var renderedTab string
		if i == t.active {
			renderedTab = t.activeStyle.Render(tab.Title)
		} else {
			renderedTab = t.inactiveStyle.Render(tab.Title)
		}
		tabStrings = append(tabStrings, renderedTab)
		currentPos += lipgloss.Width(renderedTab)
	}

	line := strings.Join(tabStrings, "")

	// Add scroll indicators if needed
	if t.scrollOffset > 0 {
		line = "◀ " + line
	}
	if visibleTabs.end < len(t.tabs) {
		line = line + " ▶"
	}

	// Pad to full width
	line = line + strings.Repeat(" ", max(0, t.width-lipgloss.Width(line)))
	if lipgloss.Width(line) > t.width {
		line = line[:t.width]
	}

	return line
}

// visibleRange represents the start and end indices of visible tabs
type visibleRange struct {
	start int
	end   int
}

// getVisibleTabs calculates which tabs should be visible based on width and scroll offset
func (t *Tabs) getVisibleTabs() visibleRange {
	if len(t.tabs) == 0 {
		return visibleRange{0, 0}
	}

	availableWidth := t.width - 4 // Reserve space for scroll indicators and padding
	if availableWidth < 10 {
		availableWidth = 10
	}

	// Start from scroll offset
	start := t.scrollOffset
	currentWidth := 0
	end := start

	// Add tabs until we exceed available width
	for end < len(t.tabs) {
		tab := t.tabs[end]
		tabWidth := len(tab.Title) + 2 // +2 for padding
		if currentWidth+tabWidth > availableWidth && end > start {
			break
		}
		currentWidth += tabWidth
		end++
	}

	return visibleRange{start, end}
}

// ensureActiveTabVisible ensures the active tab is within the visible range
func (t *Tabs) ensureActiveTabVisible() {
	visibleRange := t.getVisibleTabs()

	// If active tab is before visible range, scroll left
	if t.active < visibleRange.start {
		t.scrollOffset = t.active
		return
	}

	// If active tab is after visible range, scroll right
	if t.active >= visibleRange.end {
		// Scroll so the active tab is at the end of the visible range
		availableWidth := t.width - 4
		if availableWidth < 10 {
			availableWidth = 10
		}

		currentWidth := 0
		scrollStart := t.active

		// Work backwards from active tab
		for scrollStart > 0 {
			tab := t.tabs[scrollStart]
			tabWidth := len(tab.Title) + 2
			if currentWidth+tabWidth > availableWidth {
				break
			}
			currentWidth += tabWidth
			scrollStart--
		}

		t.scrollOffset = scrollStart
	}
}

// ActiveTab returns the currently active tab
func (t *Tabs) ActiveTab() Tab {
	if t.active >= 0 && t.active < len(t.tabs) {
		return t.tabs[t.active]
	}
	return Tab{}
}

// ActiveIndex returns the index of the active tab
func (t *Tabs) ActiveIndex() int {
	return t.active
}

// AddTab adds a new tab
func (t *Tabs) AddTab(tab Tab) {
	t.tabs = append(t.tabs, tab)
}

// RemoveTab removes a tab at the given index
func (t *Tabs) RemoveTab(index int) {
	if index >= 0 && index < len(t.tabs) {
		t.tabs = append(t.tabs[:index], t.tabs[index+1:]...)
		if t.active >= len(t.tabs) && t.active > 0 {
			t.active--
		}
		t.ensureActiveTabVisible()
	}
}

// UpdateTabContent updates the content of a tab
func (t *Tabs) UpdateTabContent(index int, content string) {
	if index >= 0 && index < len(t.tabs) {
		t.tabs[index].Content = content
	}
}

// GetTabs returns all tabs
func (t *Tabs) GetTabs() []Tab {
	return t.tabs
}

// SetActive sets the active tab index
func (t *Tabs) SetActive(index int) {
	if index >= 0 && index < len(t.tabs) {
		t.active = index
		t.ensureActiveTabVisible()
	}
}

// GetScrollOffset returns the current scroll offset
func (t *Tabs) GetScrollOffset() int {
	return t.scrollOffset
}

// SetScrollOffset sets the scroll offset
func (t *Tabs) SetScrollOffset(offset int) {
	if offset < 0 {
		offset = 0
	}
	if offset >= len(t.tabs) {
		offset = len(t.tabs) - 1
	}
	t.scrollOffset = offset
}
