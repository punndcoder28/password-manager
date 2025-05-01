package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/punndcoder28/password-manager/internal/passkey"
	"github.com/punndcoder28/password-manager/internal/session"
	"github.com/punndcoder28/password-manager/internal/storage"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password manager with a new passkey or validate existing one",
	Long: `Initialize or validate access to your password vault. If no passkey file exists,
creates a new one with the provided passkey. Otherwise, validates the provided passkey
against the existing one.

Example:
  password-manager init "my-secure-passkey"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		passkeyString := args[0]
		if passkeyString == "" {
			fmt.Println("passkey is required to access the password vault")
			os.Exit(1)
		}

		configDir := GetConfigDir()
		err := session.CreateSession(configDir)
		if err != nil {
			fmt.Printf("failed to create session: %v\n", err)
			os.Exit(1)
		}

		pm, err := passkey.NewPasskeyManager(configDir)
		if err != nil {
			fmt.Printf("failed to create passkey manager: %v\n", err)
			session.ClearSession(configDir)
			os.Exit(1)
		}

		if _, err := os.Stat(filepath.Join(configDir, "passkey.dat")); os.IsNotExist(err) {
			if err := pm.InitializePasskey(passkeyString); err != nil {
				fmt.Printf("failed to initialize passkey: %v\n", err)
				session.ClearSession(configDir)
				os.Exit(1)
			}
			fmt.Println("Password vault initialized successfully")
		} else {
			valid, err := pm.VerifyPasskey(passkeyString)
			if err != nil {
				fmt.Printf("failed to verify passkey: %v\n", err)
				session.ClearSession(configDir)
				os.Exit(1)
			}
			if !valid {
				fmt.Println("Invalid passkey")
				session.ClearSession(configDir)
				os.Exit(1)
			}
			fmt.Println("Access granted to password vault")
		}

		// Initialize the file handler
		fileHandler := storage.NewFileHandler(filepath.Join(configDir, "vault.json"))
		if err := fileHandler.Initialize(); err != nil {
			fmt.Printf("failed to initialize file handler: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
