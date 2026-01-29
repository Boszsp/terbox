package data

import "fmt"

var (
	ErrSessionNotFound   = fmt.Errorf("session not found")
	ErrSessionNotStarted = fmt.Errorf("session not started")
	ErrInvalidShell      = fmt.Errorf("invalid shell")
)
