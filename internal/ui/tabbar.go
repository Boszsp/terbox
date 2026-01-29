package ui

import (
	"fmt"
	"strings"
	"terbox/internal/mux"

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
	mux           *mux.Multiplexer
	sessions      []string
	activeIdx     int
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

// NewTabBarWithMux creates a new tab bar with multiplexer
func NewTabBarWithMux(m *mux.Multiplexer) *TabBar {
	return &TabBar{
		mux:       m,
		tabs:      []Tab{},
		active:    0,
		activeIdx: 0,
		style:     lipgloss.NewStyle(),
		activeStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("63")).
			Padding(0, 1),
		inactiveStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1),
	}
}

// Init returns no command
func (tb *TabBar) Init() tea.Cmd {
	tb.UpdateSessions()
	return nil
}

// Update handles key inputs and mouse clicks for tab navigation
func (tb *TabBar) Update(msg tea.Msg) tea.Cmd {
	if tb == nil {
		return nil
	}
	// Ensure sessions are up-to-date when using multiplexer
	if tb.mux != nil {
		tb.UpdateSessions()
	}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		tb.SetSize(msg.Width, msg.Height)
	case SessionUpdatedMsg:
		tb.UpdateSessions()
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
				// Determine the tab width depending on mode
				var tabWidth int
				if tb.mux != nil {
					// Use session label width: " [index] name "
					name := ""
					if i < len(tb.sessions) {
						info := tb.mux.GetSessionInfo(tb.sessions[i])
						if info != nil {
							if info.LastCommand != "" {
								name = truncateStr(info.LastCommand, 20)
							} else {
								name = info.Name
							}
						}
					}
					tabLabel := fmt.Sprintf(" [%d] %s ", i+1, name)
					tabWidth = lipgloss.Width(tabLabel)
				} else {
					if i < len(tb.tabs) {
						tabWidth = len(tb.tabs[i].Title) + 2
					} else {
						tabWidth = 0
					}
				}

				// Check if click is within the tab's x range
				if tabWidth > 0 && msg.X >= pos && msg.X < pos+tabWidth {
					if tb.mux != nil {
						tb.activeIdx = i
						if i < len(tb.sessions) {
							tb.mux.SetActive(tb.sessions[i])
						}
					} else {
						tb.active = i
					}
					break
				}
			}
		}
	}
	return nil
}

// View renders the tab bar
func (tb *TabBar) View() string {
	if tb.mux != nil {
		// Use multiplexer mode
		tb.sessions = tb.mux.ListSessions()
		var tabStrings []string
		tb.tabPositions = []int{}
		currentPos := 0

		for i, sessionID := range tb.sessions {
			tb.tabPositions = append(tb.tabPositions, currentPos)
			info := tb.mux.GetSessionInfo(sessionID)
			if info == nil {
				continue
			}

			tabName := info.Name
			if info.LastCommand != "" {
				tabName = truncateStr(info.LastCommand, 20)
			}

			tabLabel := fmt.Sprintf(" [%d] %s ", i+1, tabName)
			var renderedTab string
			if i == tb.activeIdx {
				renderedTab = tb.activeStyle.Render(tabLabel)
			} else {
				renderedTab = tb.inactiveStyle.Render(tabLabel)
			}
			tabStrings = append(tabStrings, renderedTab)
			currentPos += lipgloss.Width(renderedTab)
		}

		line := strings.Join(tabStrings, "")
		line = line + strings.Repeat(" ", max(0, tb.width-lipgloss.Width(line)))
		return line
	}

	// Use standard tab mode
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

// SetSize sets the dimensions of the tab bar
func (tb *TabBar) SetSize(width, height int) {
	tb.width = width
	tb.height = height
}

// UpdateSessions updates the sessions list from multiplexer
func (tb *TabBar) UpdateSessions() {
	if tb.mux == nil {
		return
	}
	tb.sessions = tb.mux.ListSessions()
	if tb.activeIdx >= len(tb.sessions) && len(tb.sessions) > 0 {
		tb.activeIdx = len(tb.sessions) - 1
	}
}

// truncateStr truncates a string to a maximum length
func truncateStr(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
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

// NextTab switches to the next tab
func (tb *TabBar) NextTab() {
	if len(tb.sessions) == 0 && len(tb.tabs) == 0 {
		return
	}

	if tb.mux != nil {
		if len(tb.sessions) > 0 {
			tb.activeIdx = (tb.activeIdx + 1) % len(tb.sessions)
			if tb.activeIdx < len(tb.sessions) {
				tb.mux.SetActive(tb.sessions[tb.activeIdx])
			}
		}
	} else {
		if len(tb.tabs) > 0 {
			tb.active = (tb.active + 1) % len(tb.tabs)
		}
	}
}

// PrevTab switches to the previous tab
func (tb *TabBar) PrevTab() {
	if len(tb.sessions) == 0 && len(tb.tabs) == 0 {
		return
	}

	if tb.mux != nil {
		if len(tb.sessions) > 0 {
			tb.activeIdx--
			if tb.activeIdx < 0 {
				tb.activeIdx = len(tb.sessions) - 1
			}
			tb.mux.SetActive(tb.sessions[tb.activeIdx])
		}
	} else {
		if len(tb.tabs) > 0 {
			tb.active--
			if tb.active < 0 {
				tb.active = len(tb.tabs) - 1
			}
		}
	}
}

// SelectTab switches to a specific tab by index
func (tb *TabBar) SelectTab(idx int) {
	if tb.mux != nil {
		if idx >= 0 && idx < len(tb.sessions) {
			tb.activeIdx = idx
			tb.mux.SetActive(tb.sessions[tb.activeIdx])
		}
	} else {
		if idx >= 0 && idx < len(tb.tabs) {
			tb.active = idx
		}
	}
}

// GetActiveSessionID returns the active session ID
func (tb *TabBar) GetActiveSessionID() string {
	if tb.activeIdx >= 0 && tb.activeIdx < len(tb.sessions) {
		return tb.sessions[tb.activeIdx]
	}
	return ""
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
