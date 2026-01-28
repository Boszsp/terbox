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
		case "ctrl+h":
			 m.browser.AddTab(
				ui.Tab{
					Title: "Help",
					Content: "This is the help tab.\n\n- Use arrow keys to navigate between tabs.\n- Press Tab to switch focus between tabs and content panel.\n- Press Ctrl+T to open a new tab.\n- Press Ctrl+W to close the current tab.\n- Press q or Ctrl+C to quit the application.",
				},
			);
			return  m, nil
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
	return m.browser.View() 
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
