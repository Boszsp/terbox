package data

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds application configuration
type Config struct {
	Shell       string            `json:"shell"`
	Theme       string            `json:"theme"`
	KeyBindings map[string]string `json:"keybindings"`
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Shell: "/bin/sh",
		Theme: "default",
		KeyBindings: map[string]string{
			"new_tab":   "ctrl+t",
			"close_tab": "ctrl+w",
			"settings":  "ctrl+s",
			"help":      "ctrl+h",
			"next_tab":  "ctrl+right",
			"prev_tab":  "ctrl+left",
			"quit":      "ctrl+q",
		},
	}
}

// GetConfigPath returns the path to the config file
func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(home, ".config", "terbox")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// LoadConfig loads configuration from file, returns default if not found
func LoadConfig() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return DefaultConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return DefaultConfig(), nil
		}
		return nil, err
	}

	config := DefaultConfig()
	if err := json.Unmarshal(data, config); err != nil {
		return DefaultConfig(), nil
	}
	return config, nil
}

// SaveConfig saves configuration to file
func SaveConfig(config *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
