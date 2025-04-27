package cmd

import (
	"fmt"

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
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
