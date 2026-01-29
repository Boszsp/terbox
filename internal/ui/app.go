package ui

import (
	"fmt"
	"terbox/internal/data"
	"terbox/internal/mux"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// App represents the main application
type App struct {
	multiplexer  *mux.Multiplexer
	tabBar       *TabBar
	terminal     *Terminal
	config       *data.Config
	theme        *Theme
	width        int
	height       int
	sessionID    int
	helpMode     bool
	settingsMode bool
}

// NewApp creates a new application
func NewApp(config *data.Config) *App {
	m := mux.NewMultiplexer(config)
	return &App{
		multiplexer:  m,
		tabBar:       NewTabBarWithMux(m),
		terminal:     NewTerminal(),
		config:       config,
		theme:        DefaultTheme(),
		sessionID:    1,
		helpMode:     false,
		settingsMode: false,
	}
}

// Init initializes the app
func (a *App) Init() tea.Cmd {
	var cmds []tea.Cmd
	if a.tabBar != nil {
		cmds = append(cmds, a.tabBar.Init())
	}
	if a.terminal != nil {
		cmds = append(cmds, a.terminal.Init())
	}
	// create first session
	cmds = append(cmds, a.createNewSession())
	return tea.Batch(cmds...)
}

// Update handles all messages
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height

		// Update component sizes
		if a.tabBar != nil {
			a.tabBar.SetSize(a.width, 1)
		}
		if a.terminal != nil {
			a.terminal.SetSize(a.width, a.height-2)
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return a, tea.Quit
		case "ctrl+t":
			cmds = append(cmds, a.createNewSession())
		case "ctrl+w":
			cmds = append(cmds, a.closeCurrentSession())
		case "ctrl+s":
			a.settingsMode = !a.settingsMode
		case "ctrl+h":
			a.helpMode = !a.helpMode
		case "ctrl+right":
			if a.tabBar != nil {
				a.tabBar.NextTab()
			}
		case "ctrl+left":
			if a.tabBar != nil {
				a.tabBar.PrevTab()
			}
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			idx := int(msg.String()[0]-'0') - 1
			if a.tabBar != nil {
				a.tabBar.SelectTab(idx)
			}
		}

	case SessionUpdatedMsg:
		// Update tab when session changes
	}

	// Delegate updates to components (guard nil)
	if a.tabBar != nil {
		a.tabBar.Update(msg)
	}
	if a.terminal != nil {
		a.terminal.Update(msg)
	}

	return a, tea.Batch(cmds...)
}

// View renders the entire app
func (a *App) View() string {
	if a.helpMode {
		return a.renderHelp()
	}

	if a.settingsMode {
		return a.renderSettings()
	}

	tabView := a.tabBar.View()
	terminalView := a.terminal.View()

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		tabView,
		lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			Render(terminalView),
	)

	return content
}

// createNewSession creates a new terminal session
func (a *App) createNewSession() tea.Cmd {
	sessionID := fmt.Sprintf("session-%d", a.sessionID)
	a.sessionID++

	session, err := a.multiplexer.CreateSession(sessionID)
	if err != nil {
		return nil
	}

	session.SetName(fmt.Sprintf("shell-%d", a.sessionID-1))
	a.tabBar.UpdateSessions()

	return nil
}

// closeCurrentSession closes the current session
func (a *App) closeCurrentSession() tea.Cmd {
	activeID := a.tabBar.GetActiveSessionID()
	if activeID != "" {
		a.multiplexer.CloseSession(activeID)
		a.tabBar.UpdateSessions()
	}
	return nil
}

// renderHelp renders the help screen
func (a *App) renderHelp() string {
	helpText := `
╔════════════════════════════════════════════════════════════════╗
║                    TERBOX - HELP                               ║
╚════════════════════════════════════════════════════════════════╝

KEYBOARD SHORTCUTS:
  Ctrl+T            Create new terminal tab
  Ctrl+W            Close current tab
  Ctrl+S            Open settings
  Ctrl+H            Show this help
  Ctrl+Right        Switch to next tab
  Ctrl+Left         Switch to previous tab
  1-9               Jump to specific tab (1=first, 9=ninth)
  Ctrl+Q            Quit application

FEATURES:
  • Web browser-like tab interface
  • One tab per terminal session
  • Auto-renaming tabs based on latest command
  • Configurable shell (default: /bin/sh)
  • Multiple themes available
  • Cross-platform (Linux and macOS)

TAB MANAGEMENT:
  • Each tab represents an independent terminal session
  • Tabs automatically rename to show the last command
  • Click on tabs to switch (mouse support)
  • Close tabs without affecting others

SETTINGS:
  Press Ctrl+S to open settings and:
  • Change default shell
  • Select theme
  • Configure keybindings
  • View advanced options

Press any key to return to terminal...
`

	box := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(2, 4).
		Width(a.width - 4).
		Render(helpText)

	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		box,
	)
}

// renderSettings renders the settings screen
func (a *App) renderSettings() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	settingsText := fmt.Sprintf(`
╔════════════════════════════════════════════════════════════════╗
║                    TERBOX - SETTINGS                           ║
╚════════════════════════════════════════════════════════════════╝

CURRENT CONFIGURATION:
  Default Shell:    %s
  Active Theme:     %s
  Open Sessions:    %d
  Last Update:      %s

AVAILABLE THEMES:
  • default       - Default light theme
  • dark          - Dark theme
  • dracula       - Dracula theme
  • nord          - Nord theme

KEYBINDINGS:
  New Tab       (Ctrl+T)       Switch Previous  (Ctrl+Left)
  Close Tab     (Ctrl+W)       Jump to Tab      (1-9)
  Settings      (Ctrl+S)       Quit             (Ctrl+Q)
  Help          (Ctrl+H)       Switch Next      (Ctrl+Right)

To change settings, edit ~/.config/terbox/config.json

Press Ctrl+S to close settings...
`, a.config.Shell, a.config.Theme, a.multiplexer.SessionCount(), now)

	box := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PrimaryColor).
		Padding(2, 4).
		Width(a.width - 4).
		Render(settingsText)

	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		box,
	)
}
