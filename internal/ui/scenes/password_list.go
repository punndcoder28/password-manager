package scenes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/punndcoder28/password-manager/internal/ui/components/list"
	"github.com/punndcoder28/password-manager/vault"
)

// PasswordListScene represents the full-screen scene for listing passwords
// This is a thin wrapper around the list component that can be extended
// to include additional UI elements (e.g., status bar, search box, etc.)
type PasswordListScene struct {
	listModel list.Model
}

// NewPasswordListScene creates a new password list scene
func NewPasswordListScene(entries map[string][]vault.Entry) PasswordListScene {
	return PasswordListScene{
		listModel: list.New(entries),
	}
}

// Init initializes the scene (required by Bubble Tea)
func (s PasswordListScene) Init() tea.Cmd {
	return s.listModel.Init()
}

// Update handles state changes
func (s PasswordListScene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := s.listModel.Update(msg)
	
	// Type assertion to convert back to list.Model
	if m, ok := updatedModel.(list.Model); ok {
		s.listModel = m
	}
	
	return s, cmd
}

// View renders the scene
func (s PasswordListScene) View() string {
	return s.listModel.View()
}

