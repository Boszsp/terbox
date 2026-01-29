# Terbox

A user-friendly terminal multiplexer with a web browser-like interface. Run multiple terminal sessions in parallel with intuitive tab management - no complex keybindings required!

## Overview

**Terbox** brings the simplicity of web browser tabs to terminal multiplexing. Unlike `tmux` or `screen`, Terbox uses a modern, familiar UI that anyone can use instantly.

## Features

- **Web Browser-Style Tabs** - Switch between terminal sessions just like browser tabs
- **Auto-Renaming Tabs** - Tabs automatically show the last command you ran
- **Terminal Session Manager** - Manage multiple independent terminal sessions
- **Configurable Shell** - Set your preferred shell (bash, zsh, fish, etc.)
- **Theme Support** - Choose from multiple color schemes
- **Cross-Platform** - Works on Linux and macOS
- **Mouse-Friendly** - Click tabs to switch (no complex key combinations)
- **Detachable Tabs** - Separate tabs into new windows for advanced workflows
- **Mouse Support** - Click on tabs to navigate
- **Tab Management** - Create and close tabs dynamically

## Theme System

The application supports easy theme customization. Choose from built-in themes or create your own:

### Using Built-in Themes

```go
package main

import (
	"terbox/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create browser with default theme
	browser := ui.NewBrowser(tabs)
	
	// Or use a specific theme
	browser := ui.NewBrowserWithTheme(tabs, ui.DarkTheme())
	browser := ui.NewBrowserWithTheme(tabs, ui.LightTheme())
}
```

### Creating a Custom Theme

```go
myTheme := &ui.Theme{
	TabActiveFg:      "255",    // White
	TabActiveBg:      "63",     // Blue
	TabInactiveFg:    "240",    // Dark gray
	TabFocusedFg:     "228",    // Yellow
	SeparatorColor:   "239",    // Dark gray
	PanelFg:          "255",    // White
}

browser := ui.NewBrowserWithTheme(tabs, myTheme)
```

Available built-in themes:
- `DefaultTheme()` - Classic blue and white theme
- `DarkTheme()` - Dark background with bright text
- `LightTheme()` - Light background with dark text

## Terminal Component

Use the interactive terminal component in the content area:

```go
package main

import (
	"terbox/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create an interactive terminal
	terminal := ui.NewTerminal()
	
	// With a custom theme
	terminal := ui.NewTerminalWithTheme(ui.DarkTheme())
	
	// Write output to terminal
	terminal.WriteOutput("Hello, World!")
	
	// Get user input
	input := terminal.GetInputBuffer()
	
	// Clear terminal
	terminal.ClearContent()
	
	// Get all history
	history := terminal.GetHistory()
}
```

### Terminal Features

- **Input Buffer** - Type commands with visual feedback
- **History Scrolling** - Scroll up/down through previous output (↑ / ↓)
- **Output Display** - Write text output to the terminal
- **Line Management** - Configurable maximum lines to keep in memory
- **Clear Command** - Clear all content

### Terminal Keyboard Shortcuts

- **Regular Keys** - Type input
- **Backspace** - Delete last character
- **Enter** - Execute command
- **↑ / ↓** - Scroll through history

## Keyboard Shortcuts

- `h` / `←` - Previous tab
- `l` / `→` - Next tab
- `1-9` - Jump to tab number
- `Tab` - Switch focus between tabs and content
- `Ctrl+T` - Create a new tab
- `Ctrl+W` - Close current tab
- `Ctrl+L` - Toggle between panel and terminal modes (in content area)
- `q` / `Ctrl+C` - Quit