package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/punndcoder28/password-manager/internal/session"
	"github.com/punndcoder28/password-manager/internal/storage"
)

func ValidateAndGetFileHandler() (*storage.FileHandler, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return nil, fmt.Errorf("error getting config directory: %w", err)
	}

	valid, err := session.ValidateSession(configDir)
	if err != nil {
		return nil, fmt.Errorf("error validating session: %w", err)
	}

	if !valid {
		return nil, fmt.Errorf("session expired. Please login again")
	}

	fileHandler := storage.NewFileHandler(filepath.Join(configDir, "vault.json"))

	if _, err := os.Stat(filepath.Join(configDir, "vault.json")); os.IsNotExist(err) {
		return nil, fmt.Errorf("vault not initialized. Please run 'init' command first")
	}

	return fileHandler, nil
}
