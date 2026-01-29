package ui

// SessionUpdatedMsg is sent when a session is updated
type SessionUpdatedMsg struct {
	SessionID string
}

// ThemeChangedMsg is sent when the theme changes
type ThemeChangedMsg struct {
	Theme string
}

// SettingsUpdatedMsg is sent when settings are updated
type SettingsUpdatedMsg struct{}

// CommandExecutedMsg is sent when a command is executed
type CommandExecutedMsg struct {
	SessionID string
	Command   string
}

// TabClosedMsg is sent when a tab is closed
type TabClosedMsg struct {
	SessionID string
}

// NewTabMsg is sent when a new tab is requested
type NewTabMsg struct{}

// QuitMsg is sent when quit is requested
type QuitMsg struct{}
