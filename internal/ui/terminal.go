package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Terminal represents an interactive terminal panel
type Terminal struct {
	width        int
	height       int
	content      []string // Lines of content/history
	inputBuffer  string   // Current input line being typed
	scrollOffset int      // For scrolling through history
	theme        *Theme
	style        lipgloss.Style
	maxLines     int // Maximum lines to keep in history (default 1000)
}

// NewTerminal creates a new terminal with default theme
func NewTerminal() *Terminal {
	return NewTerminalWithTheme(DefaultTheme())
}

// NewTerminalWithTheme creates a new terminal with a custom theme
func NewTerminalWithTheme(theme *Theme) *Terminal {
	return &Terminal{
		content:  []string{""},
		theme:    theme,
		style:    theme.GetPanelStyle(),
		maxLines: 1000,
	}
}

// SetSize sets the dimensions of the terminal
func (t *Terminal) SetSize(width, height int) {
	t.width = width
	t.height = height
}

// Init returns no command
func (t *Terminal) Init() tea.Cmd {
	return nil
}

// Update handles input for the terminal
func (t *Terminal) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Execute command (add to history)
			t.ExecuteCommand(t.inputBuffer)
			t.inputBuffer = ""
		case tea.KeyBackspace:
			if len(t.inputBuffer) > 0 {
				t.inputBuffer = t.inputBuffer[:len(t.inputBuffer)-1]
			}
		case tea.KeyUp:
			t.scrollUp()
		case tea.KeyDown:
			t.scrollDown()
		default:
			// Add regular character input
			if len(msg.String()) == 1 {
				t.inputBuffer += msg.String()
			}
		}
	}
	return nil
}

// View renders the terminal
func (t *Terminal) View() string {
	contentHeight := t.height - 1 // Reserve 1 line for input
	if contentHeight < 1 {
		contentHeight = 1
	}

	// Get visible lines from content
	visibleLines := []string{}
	startIdx := len(t.content) - contentHeight - t.scrollOffset
	if startIdx < 0 {
		startIdx = 0
	}
	endIdx := startIdx + contentHeight

	if endIdx > len(t.content) {
		endIdx = len(t.content)
	}

	for i := startIdx; i < endIdx; i++ {
		if i >= 0 && i < len(t.content) {
			visibleLines = append(visibleLines, t.truncateLine(t.content[i], t.width))
		}
	}

	// Pad to fill height
	for len(visibleLines) < contentHeight {
		visibleLines = append(visibleLines, "")
	}

	// Build output
	output := strings.Join(visibleLines, "\n")

	// Add input line
	inputLine := "$ " + t.inputBuffer + "_"
	if len(inputLine) > t.width {
		inputLine = inputLine[:t.width]
	}

	result := output + "\n" + inputLine
	return result
}

// ExecuteCommand adds a command and its output to the terminal
func (t *Terminal) ExecuteCommand(command string) {
	// Add command to history
	t.content = append(t.content, "$ "+command)

	// Trim content if exceeds maxLines
	if len(t.content) > t.maxLines {
		t.content = t.content[len(t.content)-t.maxLines:]
	}

	// Reset scroll offset when new command is executed
	t.scrollOffset = 0
}

// WriteOutput writes output to the terminal
func (t *Terminal) WriteOutput(output string) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		t.content = append(t.content, line)
	}

	// Trim content if exceeds maxLines
	if len(t.content) > t.maxLines {
		t.content = t.content[len(t.content)-t.maxLines:]
	}

	t.scrollOffset = 0
}

// ClearContent clears all terminal content
func (t *Terminal) ClearContent() {
	t.content = []string{""}
	t.inputBuffer = ""
	t.scrollOffset = 0
}

// GetInputBuffer returns the current input being typed
func (t *Terminal) GetInputBuffer() string {
	return t.inputBuffer
}

// SetInputBuffer sets the input buffer manually
func (t *Terminal) SetInputBuffer(input string) {
	t.inputBuffer = input
}

// GetContent returns all terminal content
func (t *Terminal) GetContent() []string {
	return t.content
}

// GetHistory returns formatted history as a single string
func (t *Terminal) GetHistory() string {
	return strings.Join(t.content, "\n")
}

// scrollUp scrolls up through history
func (t *Terminal) scrollUp() {
	maxScroll := len(t.content) - (t.height - 1)
	if maxScroll < 0 {
		maxScroll = 0
	}
	if t.scrollOffset < maxScroll {
		t.scrollOffset++
	}
}

// scrollDown scrolls down through history
func (t *Terminal) scrollDown() {
	if t.scrollOffset > 0 {
		t.scrollOffset--
	}
}

// truncateLine truncates a line to fit within the terminal width
func (t *Terminal) truncateLine(line string, width int) string {
	if len(line) > width {
		return line[:width]
	}
	return line
}

// SetTheme sets the theme for the terminal
func (t *Terminal) SetTheme(theme *Theme) {
	t.theme = theme
	t.style = theme.GetPanelStyle()
}

// GetTheme returns the current theme
func (t *Terminal) GetTheme() *Theme {
	return t.theme
}

// SetMaxLines sets the maximum number of lines to keep in history
func (t *Terminal) SetMaxLines(max int) {
	t.maxLines = max
	if len(t.content) > max {
		t.content = t.content[len(t.content)-max:]
	}
}

// GetMaxLines returns the maximum number of lines
func (t *Terminal) GetMaxLines() int {
	return t.maxLines
}
