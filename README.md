# terbox

A terminal-based browser UI written in Go using Bubble Tea with scrollable tabs and customizable themes.

## Features

- **Scrollable Horizontal Tabs** - Tabs automatically scroll when they exceed available width
- **Themeable UI** - Easy-to-customize color schemes for the entire application
- **Keyboard Navigation** - Navigate tabs and content with keyboard shortcuts
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

## Keyboard Shortcuts

- `h` / `←` - Previous tab
- `l` / `→` - Next tab
- `1-9` - Jump to tab number
- `Tab` - Switch focus between tabs and content
- `Ctrl+T` - Create a new tab
- `Ctrl+W` - Close current tab
- `q` / `Ctrl+C` - Quit