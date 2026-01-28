package main

// These imports will be used later on the tutorial. If you save the file
// now, Go might complain they are unused, but that's fine.
// You may also need to run `go mod tidy` to download bubbletea and its
// dependencies.
import (
	"fmt"
	"os"

	"terbox/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	browser *ui.Browser
	width   int
	height  int
}

func initialModel() model {
	tabs := []ui.Tab{
		{
			Title:   "Home",
			Content: "Welcome to terbox!\n\nThis is a browser-like TUI.\n\nUse arrow keys to navigate tabs.\nPress Tab to switch focus.\nPress Ctrl+T to create a new tab.\nPress Ctrl+W to close current tab.",
		},
		{
			Title:   "About",
			Content: "Terbox - Terminal Browser\n\nA Bubble Tea application demonstrating\nreusable UI components.\n\nFeatures:\n- Tabbed interface\n- Reusable components\n- Keyboard navigation",
		},
	}

	return model{
		browser: ui.NewBrowser(tabs),
	}
}

func (m model) Init() tea.Cmd {
	return m.browser.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.browser.SetSize(m.width, m.height)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		m.browser.Update(msg)

	case tea.MouseMsg:
		m.browser.Update(msg)
	}

	return m, nil
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Loading..."
	}
	return m.browser.View() + "\n\nPress q or Ctrl+C to quit | Click tabs or h/l to navigate | 1-9: Jump to tab | Ctrl+T: new tab | Ctrl+W: close tab"
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
