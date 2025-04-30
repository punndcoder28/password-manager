package cmd

import (
	"os"
	"path/filepath"

	"github.com/punndcoder28/password-manager/internal/storage"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "password-manager",
		Short: "A secure command-line password manager for storing and managing your credentials",
		Long: `A secure command-line password manager that helps you store and manage your passwords.

Features:
  • Secure storage of passwords using strong encryption
  • Passkey protection for your password vault
  • Easy access to stored credentials
  • Command-line interface for quick operations

Usage:
  password-manager [command] [flags]

Examples:
  # Initialize a new password vault
  password-manager init

  # Add a new password
  password-manager add github

  # Retrieve a stored password
  password-manager get github

  # List all stored passwords
  password-manager list

For more information about a specific command, use:
  password-manager [command] --help`,
	}

	// Global file handler instance
	fileHandler *storage.FileHandler
)

// GetConfigDir returns the configuration directory path
func GetConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".config", "password-manager")
}

// GetFileHandler returns the global file handler instance
func GetFileHandler() *storage.FileHandler {
	return fileHandler
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
