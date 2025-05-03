package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	fileHandler, err := ValidateAndGetFileHandler()
	if err != nil {
		return err
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
