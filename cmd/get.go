package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password",
	Long:  "Get a password from the password manager",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]
		if domain == "" {
			fmt.Println("Domain is needed to get password")
			os.Exit(1)
		}

		username := args[1]
		if username == "" {
			fmt.Println("Username is needed to get password")
			os.Exit(1)
		}

		err := getPassword(domain, username)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func getPassword(domain string, username string) error {
	fileHandler, err := ValidateAndGetFileHandler()
	if err != nil {
		return err
	}

	password, err := fileHandler.GetPassword(domain, username)
	if err != nil {
		return err
	}

	clipboard.Write(clipboard.FmtText, []byte(password))
	for range 10 {
		time.Sleep(50 * time.Millisecond)
		if string(clipboard.Read(clipboard.FmtText)) == password {
			break
		}
	}
	fmt.Println("Password copied to clipboard!")
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
