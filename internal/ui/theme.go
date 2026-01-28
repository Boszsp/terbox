package ui

import "github.com/charmbracelet/lipgloss"

// Theme defines the color scheme for the UI
type Theme struct {
	// Tab colors
	TabActiveFg    string
	TabActiveBg    string
	TabInactiveFg  string
	TabInactiveBg  string
	TabFocusedFg   string
	TabBorderColor string

	// Panel colors
	PanelFg          string
	PanelBg          string
	PanelBorderColor string

	// General colors
	SeparatorColor  string
	BackgroundColor string
}

// DefaultTheme returns the default color scheme
func DefaultTheme() *Theme {
	return &Theme{
		// Tab styling
		TabActiveFg:    "255", // White
		TabActiveBg:    "63",  // Blue
		TabInactiveFg:  "240", // Dark gray
		TabInactiveBg:  "",    // Transparent
		TabFocusedFg:   "228", // Yellow
		TabBorderColor: "239", // Dark gray

		// Panel styling
		PanelFg:          "255", // White
		PanelBg:          "",    // Transparent
		PanelBorderColor: "239", // Dark gray

		// General
		SeparatorColor:  "239", // Dark gray
		BackgroundColor: "",    // Transparent
	}
}

// DarkTheme returns a dark color scheme
func DarkTheme() *Theme {
	return &Theme{
		// Tab styling
		TabActiveFg:    "15",  // Bright white
		TabActiveBg:    "17",  // Dark blue
		TabInactiveFg:  "245", // Light gray
		TabInactiveBg:  "",    // Transparent
		TabFocusedFg:   "226", // Bright yellow
		TabBorderColor: "238", // Very dark gray

		// Panel styling
		PanelFg:          "15",  // Bright white
		PanelBg:          "",    // Transparent
		PanelBorderColor: "238", // Very dark gray

		// General
		SeparatorColor:  "238", // Very dark gray
		BackgroundColor: "",    // Transparent
	}
}

// LightTheme returns a light color scheme
func LightTheme() *Theme {
	return &Theme{
		// Tab styling
		TabActiveFg:    "0",   // Black
		TabActiveBg:    "231", // Nearly white
		TabInactiveFg:  "8",   // Dark gray
		TabInactiveBg:  "",    // Transparent
		TabFocusedFg:   "226", // Yellow
		TabBorderColor: "250", // Light gray

		// Panel styling
		PanelFg:          "0",   // Black
		PanelBg:          "",    // Transparent
		PanelBorderColor: "250", // Light gray

		// General
		SeparatorColor:  "250", // Light gray
		BackgroundColor: "",    // Transparent
	}
}

// GetTabActiveStyle returns the style for active tabs
func (t *Theme) GetTabActiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.TabActiveFg)).
		Background(lipgloss.Color(t.TabActiveBg)).
		Padding(0, 1)
}

// GetTabInactiveStyle returns the style for inactive tabs
func (t *Theme) GetTabInactiveStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.TabInactiveFg)).
		Padding(0, 1)
}

// GetTabFocusedStyle returns the style for focused tabs
func (t *Theme) GetTabFocusedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.TabFocusedFg))
}

// GetPanelStyle returns the style for panels
func (t *Theme) GetPanelStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.PanelFg))
}

// GetSeparatorColor returns the separator color
func (t *Theme) GetSeparatorColor() string {
	return t.SeparatorColor
}
