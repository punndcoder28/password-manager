package cmd

import (
	"fmt"
	"os"
	"time"

	vaultPackage "github.com/punndcoder28/password-manager/internal/vault"
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
	fileHandler, err := ValidateAndGetFileHandler()
	if err != nil {
		return err
	}

	passwordEntry := &vaultPackage.Entry{
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
