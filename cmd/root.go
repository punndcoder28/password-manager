package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
)

func GetConfigDir() (string, error) {
	configDir := filepath.Join(os.Getenv("HOME"), ".config", "password-manager")

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create config directory: %w", err)
	}

	return configDir, nil
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
