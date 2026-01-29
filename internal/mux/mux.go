package mux

import (
	"fmt"
	"sync"
	"terbox/internal/data"
	"time"
)

// Multiplexer manages multiple terminal sessions
type Multiplexer struct {
	sessions map[string]*data.TerminalSession
	order    []string
	active   string
	mu       sync.RWMutex
	config   *data.Config
}

// NewMultiplexer creates a new multiplexer
func NewMultiplexer(config *data.Config) *Multiplexer {
	return &Multiplexer{
		sessions: make(map[string]*data.TerminalSession),
		order:    []string{},
		config:   config,
	}
}

// CreateSession creates a new terminal session
func (m *Multiplexer) CreateSession(id string) (*data.TerminalSession, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.sessions[id]; exists {
		return nil, fmt.Errorf("session %s already exists", id)
	}

	session := data.NewTerminalSession(id, m.config.Shell)
	if err := session.Start(m.config.Shell); err != nil {
		return nil, err
	}

	m.sessions[id] = session
	m.order = append(m.order, id)

	if m.active == "" {
		m.active = id
	}

	return session, nil
}

// GetSession retrieves a session by ID
func (m *Multiplexer) GetSession(id string) (*data.TerminalSession, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	session, exists := m.sessions[id]
	if !exists {
		return nil, data.ErrSessionNotFound
	}
	return session, nil
}

// CloseSession closes a terminal session
func (m *Multiplexer) CloseSession(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return data.ErrSessionNotFound
	}

	if err := session.Close(); err != nil {
		return err
	}

	delete(m.sessions, id)

	// Remove from order
	for i, sid := range m.order {
		if sid == id {
			m.order = append(m.order[:i], m.order[i+1:]...)
			break
		}
	}

	// Update active session
	if m.active == id {
		if len(m.order) > 0 {
			m.active = m.order[0]
		} else {
			m.active = ""
		}
	}

	return nil
}

// ListSessions returns all session IDs in order
func (m *Multiplexer) ListSessions() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]string, len(m.order))
	copy(result, m.order)
	return result
}

// GetActive returns the currently active session
func (m *Multiplexer) GetActive() (*data.TerminalSession, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.active == "" {
		return nil, data.ErrSessionNotFound
	}

	session, exists := m.sessions[m.active]
	if !exists {
		return nil, data.ErrSessionNotFound
	}
	return session, nil
}

// SetActive sets the active session
func (m *Multiplexer) SetActive(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.sessions[id]; !exists {
		return data.ErrSessionNotFound
	}

	m.active = id
	return nil
}

// GetActiveID returns the active session ID
func (m *Multiplexer) GetActiveID() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.active
}

// NextSession switches to next session
func (m *Multiplexer) NextSession() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.order) == 0 {
		return data.ErrSessionNotFound
	}

	for i, id := range m.order {
		if id == m.active {
			m.active = m.order[(i+1)%len(m.order)]
			return nil
		}
	}

	m.active = m.order[0]
	return nil
}

// PrevSession switches to previous session
func (m *Multiplexer) PrevSession() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.order) == 0 {
		return data.ErrSessionNotFound
	}

	for i, id := range m.order {
		if id == m.active {
			idx := i - 1
			if idx < 0 {
				idx = len(m.order) - 1
			}
			m.active = m.order[idx]
			return nil
		}
	}

	m.active = m.order[0]
	return nil
}

// SessionCount returns the number of active sessions
func (m *Multiplexer) SessionCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.sessions)
}

// CleanupDeadSessions removes dead sessions
func (m *Multiplexer) CleanupDeadSessions() {
	m.mu.Lock()
	defer m.mu.Unlock()

	deadSessions := []string{}
	for id, session := range m.sessions {
		if !session.IsAlive() {
			deadSessions = append(deadSessions, id)
		}
	}

	for _, id := range deadSessions {
		delete(m.sessions, id)

		// Remove from order
		for i, sid := range m.order {
			if sid == id {
				m.order = append(m.order[:i], m.order[i+1:]...)
				break
			}
		}
	}

	// Update active if needed
	if m.active != "" {
		if _, exists := m.sessions[m.active]; !exists {
			if len(m.order) > 0 {
				m.active = m.order[0]
			} else {
				m.active = ""
			}
		}
	}
}

// GetSessionInfo returns information about a session
func (m *Multiplexer) GetSessionInfo(id string) *SessionInfo {
	session, err := m.GetSession(id)
	if err != nil {
		return nil
	}

	return &SessionInfo{
		ID:          id,
		Name:        session.GetName(),
		LastCommand: session.GetLastCommand(),
		CreatedAt:   session.CreatedAt,
		IsAlive:     session.IsAlive(),
	}
}

// SessionInfo contains metadata about a session
type SessionInfo struct {
	ID          string
	Name        string
	LastCommand string
	CreatedAt   time.Time
	IsAlive     bool
}
