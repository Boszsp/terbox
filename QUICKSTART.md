# Terbox - Quick Start Guide

## Installation

### Build from Source
```bash
git clone <repository>
cd terbox
go mod tidy
go build
```

### Run the Application
```bash
./terbox
```

---

## First Time Usage

When you start Terbox for the first time:

1. **Default Shell**: Opens with `/bin/sh` (configurable)
2. **First Tab**: Shows as "shell-1" until you run a command
3. **Configuration**: Automatically created at `~/.config/terbox/config.json`

---

## Essential Keyboard Shortcuts

| Action | Shortcut |
|--------|----------|
| New Terminal Tab | `Ctrl+T` |
| Close Current Tab | `Ctrl+W` |
| Next Tab | `Ctrl+Right` |
| Previous Tab | `Ctrl+Left` |
| Jump to Tab 1-9 | `1`-`9` |
| Settings | `Ctrl+S` |
| Help | `Ctrl+H` |
| Quit | `Ctrl+Q` |

---

## Configuration

### Location
```
~/.config/terbox/config.json
```

### Default Configuration
```json
{
  "shell": "/bin/sh",
  "theme": "default",
  "keybindings": {
    "new_tab": "ctrl+t",
    "close_tab": "ctrl+w",
    "settings": "ctrl+s",
    "help": "ctrl+h",
    "next_tab": "ctrl+right",
    "prev_tab": "ctrl+left",
    "quit": "ctrl+q"
  }
}
```

### Changing the Default Shell
Edit `~/.config/terbox/config.json` and change the `shell` field:

```json
{
  "shell": "/bin/bash",
  ...
}
```

Supported shells:
- `/bin/sh` - Default
- `/bin/bash` - Bash shell
- `/usr/bin/zsh` - Zsh shell
- `/usr/bin/fish` - Fish shell
- Any shell available on your system

---

## Features

### Multi-Tab Management
- Create new tabs with `Ctrl+T`
- Each tab is an independent terminal session
- Switch between tabs using arrow keys or number shortcuts
- Close tabs with `Ctrl+W`

### Auto-Renaming Tabs
Tabs automatically show the last command you ran:
- Run `npm start` → tab becomes "npm start"
- Run `docker ps` → tab becomes "docker ps"
- Helps you quickly identify what's running in each tab

### Settings
Press `Ctrl+S` to view:
- Current shell configuration
- Active theme
- Number of open sessions
- Keybinding reference

### Help
Press `Ctrl+H` to view:
- All keyboard shortcuts
- Feature descriptions
- Usage examples
- Configuration tips

---

## Common Workflows

### Development with Multiple Services

```bash
# Terminal 1: Frontend development
npm start

# Terminal 2: Backend API (Ctrl+T to create)
cd backend && npm start

# Terminal 3: Database (Ctrl+T to create)
docker-compose up

# Switch between them with Ctrl+Right/Ctrl+Left
```

### System Administration

```bash
# Terminal 1: Monitor logs
tail -f /var/log/syslog

# Terminal 2: System monitoring (Ctrl+T)
htop

# Terminal 3: File operations (Ctrl+T)
cd /var/www
ls -la
```

### Git Workflow

```bash
# Terminal 1: Code editing
vim src/main.go

# Terminal 2: Git operations (Ctrl+T)
git status
git add .
git commit -m "message"
```

---

## Troubleshooting

### Issue: Can't create new tabs
**Solution**: Check that the shell path in config is valid
```bash
which bash
# Use the output path in ~/.config/terbox/config.json
```

### Issue: Tab name not updating
**Solution**: The auto-rename happens after the command completes. Press Enter after typing your command.

### Issue: Lost terminal output
**Solution**: Use scrolling (arrow keys) or check history with `Ctrl+L` to clear, or close and reopen the tab.

### Issue: Terminal is slow
**Solution**: 
- Reduce terminal history size in code if needed
- Close unused tabs
- Check system resources with `top`

---

## Advanced Usage

### Creating SSH Connections

```bash
# Terminal 1: Open new tab
Ctrl+T

# Terminal 2: Connect via SSH
ssh user@remote-server

# You now have a persistent SSH session in a tab
```

### Parallel Command Execution

```bash
# Terminal 1: Long-running build
make all

# Terminal 2: Run tests (Ctrl+T)
make test

# Terminal 3: Monitor output (Ctrl+T)
tail -f build.log

# All run in parallel!
```

---

## Tips & Tricks

1. **Quick Navigation**: Use number keys (1-9) instead of arrow keys for faster tab switching
2. **Session Organization**: Open tabs in logical order - you'll see them left to right
3. **Shell Selection**: Use your preferred shell for different projects by temporarily changing config
4. **Persistent Config**: Configuration is saved automatically and persists across restarts
5. **Mouse Support**: You can click on tabs to switch (if your terminal supports it)

---

## Keyboard Shortcut Cheat Sheet

```
┌─────────────────────────────────────────────────────────┐
│         TERBOX KEYBOARD SHORTCUT REFERENCE               │
├─────────────────────────────────────────────────────────┤
│ Ctrl+T     │ Create new terminal tab                    │
│ Ctrl+W     │ Close current tab                          │
│ Ctrl+Right │ Switch to next tab                         │
│ Ctrl+Left  │ Switch to previous tab                     │
│ 1-9        │ Jump directly to tab 1-9                   │
│ Ctrl+S     │ Open settings (shows configuration)        │
│ Ctrl+H     │ Open help (shows this cheat sheet)         │
│ Ctrl+Q     │ Quit Terbox                                │
├─────────────────────────────────────────────────────────┤
│ All normal shell commands work inside each tab!         │
│ Each tab is a fully independent terminal session.       │
└─────────────────────────────────────────────────────────┘
```

---

## Project Files

- `CONCEPT.md` - Detailed project vision and features
- `ARCHITECTURE.md` - Technical architecture and design
- `IMPLEMENTATION.md` - Complete implementation details
- `main.go` - Application entry point
- `internal/data/` - Configuration and session management
- `internal/mux/` - Terminal multiplexer
- `internal/ui/` - User interface components
- `internal/utils/` - Utility functions

---

## Getting Help

1. Press `Ctrl+H` in the application for help
2. Edit `~/.config/terbox/config.json` to customize
3. Check the log output in the terminal for error messages
4. Review `ARCHITECTURE.md` for technical details
5. Refer to `CONCEPT.md` for feature descriptions

---

## System Requirements

- **OS**: Linux or macOS
- **Go**: 1.21 or later (for building)
- **Terminal**: Any terminal emulator with ANSI support
- **Shell**: bash, zsh, fish, or any POSIX-compatible shell

---

**Enjoy using Terbox! 🎉**

It's designed to be simple, intuitive, and powerful - like your favorite browser, but for terminals.
