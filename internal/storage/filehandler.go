package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	vaultPackage "github.com/punndcoder28/password-manager/vault"
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
	defer fh.mu.Unlock()

	dir := filepath.Dir(fh.filePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	_, err := os.Stat(fh.filePath)
	if err != nil && os.IsNotExist(err) {
		newVault := &vaultPackage.Vault{
			Entries: make(map[string][]vaultPackage.Entry),
		}
		return fh.writeVault(newVault)
	}
	return nil
}

func (fh *FileHandler) writeVault(vault *vaultPackage.Vault) error {
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

func (fh *FileHandler) readVault() (*vaultPackage.Vault, error) {
	data, err := os.ReadFile(fh.filePath)
	fmt.Println(fh.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var vault *vaultPackage.Vault
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

func (fh *FileHandler) AddEntry(domain string, entry *vaultPackage.Entry) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.readVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	if vault.Entries == nil {
		vault.Entries = make(map[string][]vaultPackage.Entry)
	}
	if vault.Entries[domain] == nil {
		vault.Entries[domain] = make([]vaultPackage.Entry, 0)
	}

	for _, e := range vault.Entries[domain] {
		if e.Username == entry.Username {
			return fmt.Errorf("entry for username %s in domain %s already exists. Try updating instead", entry.Username, domain)
		}
	}

	now := time.Now()
	entry.CreatedAt = now
	entry.UpdatedAt = now

	vault.Entries[domain] = append(vault.Entries[domain], *entry)

	return fh.writeVault(vault)
}

func (fh *FileHandler) GetEntry(domain string, username string) (*vaultPackage.Entry, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.readVault()
	if err != nil {
		return nil, fmt.Errorf("failed to read vault: %w", err)
	}

	entries, exists := vault.Entries[domain]
	if !exists {
		return nil, fmt.Errorf("no entries found for domain %s", domain)
	}

	for _, entry := range entries {
		if entry.Username == username {
			return &entry, nil
		}
	}

	return nil, fmt.Errorf("entry for username %s in domain %s not found", username, domain)
}

func (fh *FileHandler) UpdateEntry(domain string, username string, entry *vaultPackage.Entry) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.readVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	entries, exists := vault.Entries[domain]
	if !exists {
		return fmt.Errorf("no entries found for domain %s", domain)
	}

	found := false
	for i, e := range entries {
		if e.Username == username {
			now := time.Now()
			entry.UpdatedAt = now
			entries[i] = *entry
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("entry for username %s in domain %s not found. Try adding instead", username, domain)
	}

	vault.Entries[domain] = entries
	return fh.writeVault(vault)
}

func (fh *FileHandler) DeactivateEntry(domain string, username string) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.readVault()
	if err != nil {
		return fmt.Errorf("failed to read vault: %w", err)
	}

	entries, exists := vault.Entries[domain]
	if !exists {
		return fmt.Errorf("no entries found for domain %s", domain)
	}

	found := false
	for i, entry := range entries {
		if entry.Username == username {
			if !entry.IsActive {
				return fmt.Errorf("entry for username %s in domain %s is already deactivated", username, domain)
			}
			now := time.Now()
			entry.DeactivatedAt = now
			entry.IsActive = false
			entries[i] = entry
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("entry for username %s in domain %s not found", username, domain)
	}

	vault.Entries[domain] = entries
	return fh.writeVault(vault)
}

func (fh *FileHandler) ListEntries() (map[string][]vaultPackage.Entry, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	vault, err := fh.readVault()
	if err != nil {
		return nil, fmt.Errorf("error while reading vault: %w", err)
	}

	entries := make(map[string][]vaultPackage.Entry)
	for domain, domainEntries := range vault.Entries {
		entries[domain] = make([]vaultPackage.Entry, 0)
		for _, entry := range domainEntries {
			if entry.IsActive {
				entries[domain] = append(entries[domain], entry)
			}
		}
	}

	return entries, nil
}
