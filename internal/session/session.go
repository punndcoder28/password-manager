package session

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Session struct {
	CreatedAt time.Time
	ExpiresAt time.Time
}

func CreateSession(configDir string) error {
	session := &Session{
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
	}

	data, err := json.Marshal(session)
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	filepath := filepath.Join(configDir, "session.json")
	return os.WriteFile(filepath, data, 0600)
}

func ClearSession(configDir string) error {
	sessionPath := filepath.Join(configDir, "session.json")
	return os.Remove(sessionPath)
}

func ValidateSession(configDir string) (bool, error) {
	sessionPath := filepath.Join(configDir, "session.json")
	data, err := os.ReadFile(sessionPath)
	if err != nil {
		return false, err
	}

	var session Session
	if err := json.Unmarshal(data, &session); err != nil {
		return false, err
	}

	if time.Now().After(session.ExpiresAt) {
		return false, nil
	}

	return true, nil
}
