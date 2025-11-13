package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/punndcoder28/password-manager/internal/ui"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the passwords in the vault",
	Long: `List all the passwords in the vault in an interactive tree view.

Example:
password-manager list

Navigation:
- Use arrow keys (↑/↓) or vim keys (j/k) to navigate
- Press Enter or Space to expand/collapse domains or toggle password reveal
- Press 'r' to reveal/hide the selected password
- Press 'R' to reveal/hide all passwords
- Press 'q' to quit
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

	entries, err := fileHandler.ListEntriesWithMetadata()
	if err != nil {
		return fmt.Errorf("error listing entries: %w", err)
	}

	// Check if there are any entries
	if len(entries) == 0 {
		fmt.Println("No passwords found in the vault.")
		return nil
	}

	// Create and run the Bubble Tea program
	model := ui.InitialModel(entries)
	program := tea.NewProgram(model, tea.WithAltScreen())
	
	if _, err := program.Run(); err != nil {
		return fmt.Errorf("error running UI: %w", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
}
