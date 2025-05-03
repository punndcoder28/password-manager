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
		configDir, err := GetConfigDir()
		if err != nil {
			fmt.Println("Error getting config directory:", err)
		}

		valid, err := session.ValidateSession(configDir)
		if err != nil {
			fmt.Println("Error validating session:", err)
		}

		if !valid {
			fmt.Println("Session expired. Please login again.")
			return
		}

		fileHandler := storage.NewFileHandler(filepath.Join(configDir, "vault.json"))

		if _, err := os.Stat(filepath.Join(configDir, "vault.json")); os.IsNotExist(err) {
			fmt.Println("Vault not initialized. Please run 'init' command first")
			return
		}

		entries, err := fileHandler.ListEntries()
		if err != nil {
			fmt.Println("Error listing entries:", err)
			return
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
