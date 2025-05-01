package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/punndcoder28/password-manager/vault"
)

type FileHandler struct {
	filePath string
	mu       sync.Mutex
}

func NewFileHandler(filePath string) *FileHandler {
	return &FileHandler{
		filePath: filePath,
	}
}

func (fh *FileHandler) Initialize() error {
	fh.mu.Lock()
	dir := filepath.Dir(fh.filePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		fh.mu.Unlock()
		return fmt.Errorf("failed to create directory: %w", err)
	}

	_, err := os.Stat(fh.filePath)
	fh.mu.Unlock()

	if err != nil && os.IsNotExist(err) {
		newVault := &vault.Vault{
			Entries: make(map[string]vault.Entry),
		}
		return fh.WriteVault(newVault)
	}
	return nil
}

func (fh *FileHandler) WriteVault(vault *vault.Vault) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	data, err := json.MarshalIndent(vault, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal vault: %w", err)
	}

	tempFile := fh.filePath + ".tmp"
	if err := os.WriteFile(tempFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write temporary file: %w", err)
	}

	if err := os.Rename(tempFile, fh.filePath); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to rename temporary file: %w", err)
	}

	return nil
}

func (fh *FileHandler) ReadVault() (*vault.Vault, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	data, err := os.ReadFile(fh.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var vault *vault.Vault
	if err := json.Unmarshal(data, &vault); err != nil {
		return nil, fmt.Errorf("failed to unmarshal vault: %w", err)
	}

	return vault, nil
}

// SHOULD NEVER BE USED UNLESS YOU WANT TO DELETE
// THE VAULT AND LOOSE ALL YOUR PASSWORDS
func (fh *FileHandler) DeleteVault() error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	if err := os.Remove(fh.filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (fh *FileHandler) AddEntry(domain string, entry *vault.Entry) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.ReadVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	if _, exists := vault.Entries[domain]; exists {
		return fmt.Errorf("entry for domain %s already exists. Try updating instead", domain)
	}

	now := time.Now()
	entry.CreatedAt = now
	entry.UpdatedAt = now

	vault.Entries[domain] = *entry

	return fh.WriteVault(vault)
}

func (fh *FileHandler) GetEntry(domain string) (*vault.Entry, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.ReadVault()
	if err != nil {
		return nil, fmt.Errorf("failed to read vault: %w", err)
	}

	entry, exists := vault.Entries[domain]
	if !exists {
		return nil, fmt.Errorf("entry for domain %s not found", domain)
	}

	return &entry, nil
}

func (fh *FileHandler) UpdateEntry(domain string, entry *vault.Entry) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.ReadVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	if _, exists := vault.Entries[domain]; !exists {
		return fmt.Errorf("entry for domain %s not found. Try adding instead", domain)
	}

	now := time.Now()
	entry.UpdatedAt = now

	vault.Entries[domain] = *entry

	return fh.WriteVault(vault)
}

func (fh *FileHandler) DeactivateEntry(domain string) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.ReadVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	entry, exists := vault.Entries[domain]
	if !exists {
		return fmt.Errorf("entry for domain %s not found. Try adding instead", domain)
	}

	if !entry.IsActive {
		return fmt.Errorf("entry for domain %s is already deactivated", domain)
	}

	now := time.Now()
	entry.DeactivatedAt = now
	entry.IsActive = false

	vault.Entries[domain] = entry

	return fh.WriteVault(vault)
}
