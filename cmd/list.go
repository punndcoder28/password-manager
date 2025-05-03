package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/punndcoder28/password-manager/internal/session"
	"github.com/punndcoder28/password-manager/internal/storage"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the passwords in the vault",
	Long: `List all the passwords in the vault.

Example:
password-manager list
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := listPasswords(); err != nil {
			fmt.Printf("failed to list passwords: %v\n", err)
			os.Exit(1)
		}
	},
}

func listPasswords() error {
	configDir, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("error getting config directory: %w", err)
	}

	valid, err := session.ValidateSession(configDir)
	if err != nil {
		return fmt.Errorf("error validating session: %w", err)
	}

	if !valid {
		return fmt.Errorf("session expired. Please login again.")
	}

	fileHandler := storage.NewFileHandler(filepath.Join(configDir, "vault.json"))

	if _, err := os.Stat(filepath.Join(configDir, "vault.json")); os.IsNotExist(err) {
		return fmt.Errorf("vault not initialized. Please run 'init' command first")
	}

	entries, err := fileHandler.ListEntries()
	if err != nil {
		return fmt.Errorf("error listing entries: %w", err)
	}

	fmt.Println("Passwords:")
	for domain, entries := range entries {
		fmt.Printf("Domain: %s\n", domain)
		for _, entry := range entries {
			fmt.Printf("Username: %s\n", entry.Username)
			fmt.Printf("Password: %s\n", entry.Password)
			fmt.Println("--------------------------------")
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
