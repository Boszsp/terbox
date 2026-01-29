package data

import (
	"io"
	"os"
	"os/exec"
	"sync"
	"time"
)

// TerminalSession represents a single terminal session
type TerminalSession struct {
	ID          string
	Name        string
	Cmd         *exec.Cmd
	Input       io.WriteCloser
	Output      io.Reader
	LastCommand string
	CreatedAt   time.Time
	mu          sync.RWMutex
}

// NewTerminalSession creates a new terminal session
func NewTerminalSession(id string, shell string) *TerminalSession {
	return &TerminalSession{
		ID:        id,
		Name:      "shell",
		CreatedAt: time.Now(),
	}
}

// Start starts the terminal session
func (ts *TerminalSession) Start(shell string) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	cmd := exec.Command(shell)
	cmd.Env = os.Environ()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	ts.Cmd = cmd
	ts.Input = stdin
	ts.Output = stdout
	return nil
}

// WriteCommand writes a command to the session
func (ts *TerminalSession) WriteCommand(command string) error {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if ts.Input == nil {
		return ErrSessionNotStarted
	}

	_, err := ts.Input.Write([]byte(command + "\n"))
	if err != nil {
		return err
	}

	ts.mu.Lock()
	ts.LastCommand = command
	ts.mu.Unlock()

	return nil
}

// GetName returns the session name (tab name)
func (ts *TerminalSession) GetName() string {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.Name
}

// SetName sets the session name
func (ts *TerminalSession) SetName(name string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.Name = name
}

// GetLastCommand returns the last command run
func (ts *TerminalSession) GetLastCommand() string {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	return ts.LastCommand
}

// Close closes the terminal session
func (ts *TerminalSession) Close() error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	if ts.Input != nil {
		ts.Input.Close()
	}

	if ts.Cmd != nil && ts.Cmd.Process != nil {
		return ts.Cmd.Process.Kill()
	}

	return nil
}

// IsAlive checks if the session is still running
func (ts *TerminalSession) IsAlive() bool {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	if ts.Cmd == nil || ts.Cmd.Process == nil {
		return false
	}

	// Try to find the process
	proc, err := os.FindProcess(ts.Cmd.Process.Pid)
	if err != nil {
		return false
	}

	// Send signal 0 to check if process exists
	err = proc.Signal(os.Signal(nil))
	return err == nil
}
