package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/punndcoder28/password-manager/internal/session"
	"github.com/punndcoder28/password-manager/internal/storage"
	"github.com/punndcoder28/password-manager/vault"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password to the password vault",
	Long: `Add a new password to the password vault. The password is encrypted at rest using the passkey.
	
	Example:
	password-manager add <website> <username> <password>
	`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		website := args[0]
		if website == "" {
			fmt.Println("website is required")
			os.Exit(1)
		}

		username := args[1]
		if username == "" {
			fmt.Println("username is required")
			os.Exit(1)
		}

		password := args[2]
		if password == "" {
			fmt.Println("password is required")
			os.Exit(1)
		}

		if err := addPassword(website, username, password); err != nil {
			fmt.Printf("failed to add password: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Password added successfully")
	},
}

func addPassword(website string, username string, password string) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	valid, err := session.ValidateSession(configDir)
	if err != nil {
		return fmt.Errorf("failed to validate session: %w", err)
	}

	if !valid {
		return fmt.Errorf("session is invalid")
	}

	// Create a new file handler instance
	fileHandler := storage.NewFileHandler(filepath.Join(configDir, "vault.json"))

	// Check if the vault file exists
	if _, err := os.Stat(filepath.Join(configDir, "vault.json")); os.IsNotExist(err) {
		return fmt.Errorf("vault not initialized. Please run 'init' command first")
	}

	passwordEntry := &vault.Entry{
		Username:  username,
		Password:  password,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return fileHandler.AddEntry(website, passwordEntry)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
