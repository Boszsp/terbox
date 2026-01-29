# Terbox - App Concept Document

## Overview
**Terbox** is a user-friendly terminal multiplexer with a web browser-like interface. It provides a modern, intuitive way to manage multiple terminal sessions without requiring deep knowledge of complex terminal multiplexers like tmux.

---

## Core Purpose
Replace tmux/screen with a simpler, more accessible multi-tab terminal interface that leverages familiar web browser navigation patterns.

---

## Key Features

### 1. Multi-Tab Interface
- **Web Browser Style Tabs**: Users switch between terminal sessions using familiar tab UI
- **One Tab = One Terminal Session**: Each tab represents a completely independent terminal session
- **Tab Management**: 
  - Create new tabs (new terminal sessions)
  - Close tabs
  - Reorder/separate tabs
  - Switch tabs using mouse or keyboard shortcuts

### 2. Auto-Renaming Tabs
- Tabs automatically rename based on the **latest command executed** in that terminal
- Helps users quickly identify what each terminal session is doing
- Example: Running `npm start` in a tab renames it to "npm start"

### 3. Starter Shell Configuration
- **Default Shell**: `/bin/sh` (configurable)
- Users can set their preferred shell in settings
- Each new terminal session launches with the configured shell
- Configuration persists across sessions

### 4. Special Tabs
- **Settings Tab** (Ctrl+S): Configure shell, theme, appearance, keybindings
- **Help Tab** (Ctrl+H): Display help documentation and keyboard shortcuts

### 5. Terminal Session Manager
- **Session Persistence**: Terminal sessions maintain state while tabs are open
- **Independent Sessions**: Each tab's terminal is completely isolated from others
- **Memory Management**: Sessions can be closed to free resources
- **Session Switching**: Fast switching between tabs without losing context

### 6. Theme Support
- **Multiple Themes**: Users can choose different color schemes and visual styles
- **Per-Session Persistence**: Theme preference saved across launches
- **Live Theme Switching**: Change theme without restarting the app

### 7. Tab Separation
- **Detach Tabs**: Ability to separate a tab into a new window or instance
- **Advanced Workflow**: Useful for organizing terminals across multiple screens or workspaces

### 8. Cross-Platform Support
- **Linux**: Full support
- **macOS**: Full support
- **Shell Compatibility**: Works with bash, zsh, fish, ksh, etc.

---

## Architecture Overview

### Components
1. **UI Layer** (`internal/ui/`)
   - `tabbar.go`: Tab bar rendering and management
   - `tabs.go`: Individual tab component logic
   - `terminal.go`: Terminal emulation and rendering
   - `panel.go`: Main panel layout
   - `component.go`: Base component interface
   - `browser.go`: Browser/app-level management
   - `theme.go`: Theme rendering and switching
   - `list.go`: List component for menus

2. **Data Layer** (`internal/data/`)
   - Terminal session state
   - Configuration/settings storage
   - Theme definitions

3. **Utils** (`internal/utils/`)
   - Helper functions
   - Cross-platform utilities

---

## User Workflow

### Starting the App
```
$ ./terbox
```
- Opens with one default terminal tab (shell)
- User sees familiar tab bar at top/bottom
- Full terminal interface in main area

### Daily Usage
1. **Create new tab**: `Ctrl+T` → launches new terminal session
2. **Switch tabs**: Click tab or `Ctrl+[Number]` or `Ctrl+[Arrow]`
3. **Run command**: Type commands normally
4. **Tab auto-renames**: Tab name updates to show command
5. **Access settings**: `Ctrl+S` → opens settings tab
6. **Get help**: `Ctrl+H` → opens help tab
7. **Close tab**: `Ctrl+W` → closes terminal session

### Advanced Usage
- **Separate tab**: Right-click tab → "Detach in new window"
- **Change theme**: Settings → Theme selection
- **Configure shell**: Settings → Shell preference

---

## Comparison with tmux

| Feature | Terbox | tmux |
|---------|--------|------|
| **Learning Curve** | Beginner-friendly | Steep |
| **Keybindings** | Familiar (web-like) | Complex prefix-based |
| **Visual UI** | Modern tab interface | Text-based commands |
| **Accessibility** | High (intuitive) | Medium (powerful but complex) |
| **Scriptability** | User-friendly config | Advanced scripting |
| **Tab Management** | Visual and click-based | Command-based |
| **Terminal Multiplexing** | ✓ | ✓ |
| **Session Persistence** | ✓ (while app runs) | ✓ (persistent across sessions) |

---

## Technical Specifications

### Input/Output
- **Input**: Keyboard, mouse events
- **Output**: Terminal emulation display
- **Interaction**: TUI (Terminal User Interface) using rendering library

### Key Behaviors
1. **Multi-tab rendering**: All visible tabs show content simultaneously
2. **Command capture**: System captures latest command from each terminal
3. **State management**: Each tab maintains independent terminal state
4. **Configuration persistence**: Settings saved to config file

### Dependencies
- Terminal emulation library (cross-platform)
- TUI rendering framework
- Platform-specific shell execution (Linux/macOS)

---

## Use Cases

### Development Teams
- Run multiple services in parallel (frontend, backend, database)
- Switch between projects easily
- See at a glance which service is running

### DevOps Engineers
- Monitor multiple servers/containers
- Run parallel deployment tasks
- Quick context switching without terminal complexity

### System Administrators
- Manage multiple SSH sessions
- Build/deployment monitoring
- File operations across systems

### General Users
- Beginner-friendly multiplexing
- No steep learning curve
- Intuitive interface

---

## Advantages

✓ **Web Browser Familiarity**: Tab interface everyone knows  
✓ **No Learning Curve**: Intuitive for new users  
✓ **Visual Feedback**: Auto-rename shows what's running  
✓ **Cross-Platform**: Works on Linux and macOS  
✓ **Customizable**: Theme and shell options  
✓ **Session Management**: Easy terminal organization  
✓ **Mouse-Friendly**: No keyboard-only navigation required  

---

## Future Enhancements
- Session persistence across app restarts
- Tab grouping/workspaces
- Split panes within tabs
- Command history/replay
- Terminal logging
- SSH session manager
- Shell integration plugins

---

## Summary
Terbox is a modern, accessible alternative to tmux that brings web browser-style tab management to terminal multiplexing. It maintains powerful multi-session capabilities while remaining friendly and intuitive for users of all technical levels.
