# Terbox Architecture & Implementation Guide

## Project Structure

```
terbox/
├── main.go                      # Application entry point
├── go.mod                       # Go module definition
├── README.md                    # User documentation
├── CONCEPT.md                   # Project concept and vision
├── ARCHITECTURE.md              # This file
│
└── internal/
    ├── data/                    # Data models and state management
    │   ├── config.go           # Configuration management
    │   ├── session.go          # Terminal session representation
    │   ├── errors.go           # Custom error types
    │   └── [SESSION STATE]     # Terminal output buffering
    │
    ├── mux/                    # Terminal multiplexer (session management)
    │   └── mux.go             # Main multiplexer logic
    │
    ├── ui/                     # User interface components
    │   ├── component.go        # Base component interface
    │   ├── app.go             # Main app component
    │   ├── tabbar.go          # Tab bar implementation
    │   ├── terminal.go        # Terminal display
    │   ├── panel.go           # Content panels
    │   ├── tabs.go            # Advanced tab management
    │   ├── browser.go         # Browser window container
    │   ├── list.go            # List component
    │   ├── theme.go           # Theme definitions
    │   ├── messages.go        # Event messages
    │   └── styles.go          # UI styling
    │
    └── utils/                 # Utility functions
        ├── platform.go        # Platform detection
        └── strings.go         # String manipulation
```

## Architecture Overview

### Layered Architecture

```
┌─────────────────────────────────────┐
│         UI Layer                    │
│  (Browser, TabBar, Terminal, etc)   │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│      Multiplexer Layer              │
│   (Session Management, mux.go)      │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│      Data Layer                     │
│  (Config, Sessions, Errors)         │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│  Utils & Platform Layer             │
│  (Strings, Platform Detection)      │
└─────────────────────────────────────┘
```

## Core Components

### 1. **Data Layer** (`internal/data/`)

#### `config.go` - Configuration Management
```go
type Config struct {
    Shell       string                // Default shell (e.g., /bin/sh)
    Theme       string                // Active theme name
    KeyBindings map[string]string     // Customizable keybindings
}
```

**Functions:**
- `DefaultConfig()` - Returns sensible defaults
- `LoadConfig()` - Loads from `~/.config/terbox/config.json`
- `SaveConfig()` - Persists config changes
- `GetConfigPath()` - Resolves config file location

#### `session.go` - Terminal Session
```go
type TerminalSession struct {
    ID          string           // Unique session identifier
    Name        string           // Display name (tab title)
    Cmd         *exec.Cmd        // Running shell process
    Input       io.WriteCloser   // stdin pipe
    Output      io.Reader        // stdout pipe
    LastCommand string           // Latest command for tab renaming
    CreatedAt   time.Time        // Session creation time
}
```

**Methods:**
- `NewTerminalSession()` - Creates new session
- `Start()` - Launches shell process
- `WriteCommand()` - Sends input to shell
- `Close()` - Terminates session
- `IsAlive()` - Checks if process is running
- `GetName()/SetName()` - Tab name management

#### `errors.go` - Error Types
- `ErrSessionNotFound`
- `ErrSessionNotStarted`
- `ErrInvalidShell`

### 2. **Multiplexer Layer** (`internal/mux/`)

#### `mux.go` - Terminal Multiplexer
Manages multiple terminal sessions with tab-like switching.

```go
type Multiplexer struct {
    sessions map[string]*TerminalSession  // Map of all sessions
    order    []string                     // Order for tab display
    active   string                       // Currently active session ID
    config   *Config                      // Configuration reference
}
```

**Core Methods:**
- `CreateSession(id)` - Create new terminal session
- `GetSession(id)` - Retrieve session by ID
- `CloseSession(id)` - Terminate session
- `ListSessions()` - Get all session IDs in order
- `SetActive(id)` - Switch active session
- `NextSession()` / `PrevSession()` - Session switching
- `SessionCount()` - Number of active sessions
- `CleanupDeadSessions()` - Remove terminated sessions
- `GetSessionInfo(id)` - Get session metadata

**Session Management:**
```
┌─────────────────────────────────────┐
│  Multiplexer                        │
│  ┌─────────┬─────────┬─────────┐   │
│  │ Session │ Session │ Session │   │
│  │    1    │    2    │    3    │   │
│  └─────────┴─────────┴─────────┘   │
│  Active: Session 2                  │
└─────────────────────────────────────┘
```

### 3. **UI Layer** (`internal/ui/`)

#### `component.go` - Base Component Interface
```go
type Component interface {
    Init() bubbletea.Cmd
    Update(msg bubbletea.Msg) bubbletea.Cmd
    View() string
    SetSize(width, height int)
}
```

#### `app.go` - Main Application
The top-level Bubble Tea model that orchestrates all components.

**Responsibilities:**
- Creates and manages multiplexer
- Handles global keybindings (Ctrl+T, Ctrl+S, etc.)
- Routes messages to sub-components
- Renders help and settings screens
- Manages application lifecycle

**Key Methods:**
- `Init()` - Initialize multiplexer and first session
- `Update()` - Handle keyboard/mouse input
- `View()` - Render complete UI

**Keybindings:**
| Binding | Action |
|---------|--------|
| Ctrl+T | New terminal tab |
| Ctrl+W | Close current tab |
| Ctrl+S | Open settings |
| Ctrl+H | Show help |
| Ctrl+Right | Next tab |
| Ctrl+Left | Previous tab |
| 1-9 | Jump to tab |
| Ctrl+Q | Quit |

#### `tabbar.go` - Tab Navigation
Displays tabs and handles switching.

```go
type TabBar struct {
    mux       *Multiplexer          // Reference to multiplexer
    sessions  []string              // Current session IDs
    activeIdx int                   // Active tab index
    // ... rendering fields
}
```

**Methods:**
- `UpdateSessions()` - Sync with multiplexer
- `NextTab()` / `PrevTab()` / `SelectTab()` - Navigation
- `GetActiveSessionID()` - Current session

#### `terminal.go` - Terminal Display
Shows the active terminal session content.

#### `panel.go` - Content Panel
Container for displaying information.

#### `browser.go` - Browser Window
Main window container combining all components.

#### `messages.go` - Event System
Custom Bubble Tea message types:
- `SessionUpdatedMsg` - When session state changes
- `ThemeChangedMsg` - Theme selection changed
- `SettingsUpdatedMsg` - Settings modified
- `CommandExecutedMsg` - New command ran
- `TabClosedMsg` - Tab closed
- `NewTabMsg` - Create new tab request
- `QuitMsg` - Application quit request

#### `theme.go` - Color Schemes
Theme definitions with color palettes.

### 4. **Utilities** (`internal/utils/`)

#### `platform.go` - Platform Detection
- `GetShell()` - Get platform default shell
- `IsValidShell()` - Verify shell path
- `IsLinux()` / `IsMacOS()` / `IsWindows()` - Platform checks

#### `strings.go` - String Utilities
- `TruncateString()` - Limit string length
- `ParseCommand()` - Extract command name
- `FormatSessionName()` - Create tab names
- `PadString()` / `CenterString()` - String formatting

## Data Flow

### Session Creation Flow
```
User presses Ctrl+T
         ↓
    App.Update()
         ↓
    CreateSession()
         ↓
    Multiplexer.CreateSession()
         ↓
    TerminalSession.Start()
         ↓
    exec.Command() starts shell
         ↓
    TabBar.UpdateSessions()
         ↓
    View() renders new tab
```

### Command Execution Flow
```
User types command in terminal
         ↓
    Terminal captures input
         ↓
    Send to active session
         ↓
    TerminalSession.WriteCommand()
         ↓
    Write to shell stdin
         ↓
    Shell executes command
         ↓
    Output captured from stdout
         ↓
    SessionUpdatedMsg sent
         ↓
    Tab name auto-updates
         ↓
    View() re-renders
```

### Tab Switching Flow
```
User presses Ctrl+Right
         ↓
    App.Update() detects key
         ↓
    TabBar.NextTab()
         ↓
    Multiplexer.SetActive()
         ↓
    Active session ID changes
         ↓
    Terminal switches to new session
         ↓
    View() displays new content
```

## Key Design Decisions

### 1. **Separation of Concerns**
- **Data Layer**: Manages state (sessions, config)
- **Mux Layer**: Orchestrates sessions
- **UI Layer**: Handles all rendering and user interaction
- **Utils**: Platform-specific and helper functions

### 2. **Message-Driven Architecture**
Uses Bubble Tea's event system for loose coupling:
- Components don't directly call each other
- Messages flow through the main app
- Easy to add new features without modifying existing code

### 3. **Thread-Safe Session Management**
- Terminal sessions use `sync.RWMutex` for thread safety
- Multiplexer synchronizes access to session map
- Safe concurrent access from UI and background goroutines

### 4. **Stateless Components**
- UI components don't hold state
- All state in data layer
- Components are "dumb" renderers

### 5. **Configuration Persistence**
- Config stored as JSON in `~/.config/terbox/config.json`
- Loads on startup, saves on change
- Defaults for new installations

## Adding New Features

### Add a New Keybinding

1. Define the keybinding in `DefaultConfig()` in `config.go`
2. Add handling in `App.Update()` in `app.go`
3. Create corresponding handler method
4. Update `HELP` text

### Add a New Theme

1. Create theme variant in `theme.go`
2. Add to theme selector in settings
3. Save selection to config
4. Apply colors to components

### Add a New Component

1. Create new file in `internal/ui/`
2. Implement `Component` interface
3. Add to main app layout
4. Handle its messages in `App.Update()`

### Add Command to Session

1. Call `session.WriteCommand(cmd)` from UI
2. Command captured in `session.LastCommand`
3. Tab auto-renames on next render

## Configuration File

Located at `~/.config/terbox/config.json`:

```json
{
  "shell": "/bin/bash",
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

## Performance Considerations

### Session Limits
- Default max history: 1000 lines per session
- Configurable in `TerminalSession.maxLines`
- Prevents excessive memory usage

### Cleanup
- Dead sessions detected and removed via `CleanupDeadSessions()`
- Runs periodically (can be on interval)
- Prevents zombie processes

### Rendering
- Only active tab's content fully rendered
- Inactive tabs show metadata only
- Efficient Bubble Tea render cycles

## Testing Strategy

### Unit Tests
- `data/` - Config loading, session management
- `mux/` - Session creation, switching, cleanup
- `utils/` - String formatting, platform detection

### Integration Tests
- App creation and initialization
- Session lifecycle (create, switch, close)
- Configuration persistence

### Manual Testing
- Keyboard shortcuts
- Mouse interaction
- Tab switching with multiple sessions
- Theme switching
- Settings panel

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- Standard library for shell execution and file I/O

## Future Enhancements

1. **Split Panes** - Divide tabs into panes
2. **Session Persistence** - Save/restore sessions across restarts
3. **SSH Manager** - Integrated SSH connection manager
4. **Command History** - Search and replay previous commands
5. **Logging** - Save terminal output to files
6. **Plugins** - Extensible plugin system
7. **Workspaces** - Group tabs into projects
8. **Copy/Paste** - System clipboard integration

---

This architecture provides a solid foundation for a modern, user-friendly terminal multiplexer while remaining maintainable and extensible.
