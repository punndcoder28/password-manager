/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.password-manager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
