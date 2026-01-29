package utils

import (
	"fmt"
	"strings"
)

// TruncateString truncates a string to a maximum length
func TruncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

// ParseCommand extracts the command name from a command string
func ParseCommand(cmdStr string) string {
	parts := strings.Fields(cmdStr)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// FormatSessionName creates a formatted session name
func FormatSessionName(command string, index int) string {
	if command == "" {
		return fmt.Sprintf("shell-%d", index)
	}
	return TruncateString(command, 20)
}

// PadString pads a string to a given length
func PadString(s string, length int) string {
	if len(s) >= length {
		return s
	}
	return s + strings.Repeat(" ", length-len(s))
}

// CenterString centers a string within a given width
func CenterString(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := (width - len(s)) / 2
	return strings.Repeat(" ", padding) + s
}
