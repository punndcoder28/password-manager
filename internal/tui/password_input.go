package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputField int

const (
	FieldWebsite InputField = iota
	FieldUsername
	FieldPassword
)

type PasswordInputModel struct {
	website  string
	username string
	password string

	currentField InputField
	input        string
	err          error

	onComplete func() tea.Msg
}

type PasswordData struct {
	Website  string
	Username string
	Password string
}

func NewPasswordInput(onComplete func(PasswordData) tea.Cmd) PasswordInputModel {
	return PasswordInputModel{
		currentField: FieldWebsite,
		onComplete: func() tea.Msg {
			return onComplete(PasswordData{
				Website:  "",
				Username: "",
				Password: "",
			})
		},
	}
}

func (m PasswordInputModel) Init() tea.Cmd {
	return nil
}

func (m PasswordInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.currentField {
			case FieldWebsite:
				m.website = m.input
				m.input = ""
				m.currentField = FieldUsername
			case FieldUsername:
				m.username = m.input
				m.input = ""
				m.currentField = FieldPassword
			case FieldPassword:
				m.password = m.input
				m.input = ""
				return m, func() tea.Msg {
					return m.onComplete()
				}
			}
		case tea.KeyEsc:
			return m, tea.Quit
		default:
			if msg.Type == tea.KeyRunes {
				m.input += string(msg.Runes)
			}
		}
	}
	return m, nil
}

func (m PasswordInputModel) View() string {
	var s string

	titleStyle := lipgloss.NewStyle().Bold(true)
	currentStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	s += titleStyle.Render("New Password Entry") + "\n\n"

	if m.currentField == FieldWebsite {
		s += currentStyle.Render("Website: ") + m.input + "\n"
	} else {
		s += "Website: " + m.website + "\n"
	}

	if m.currentField == FieldUsername {
		s += currentStyle.Render("Username: ") + m.input + "\n"
	} else {
		s += "Username: " + m.username + "\n"
	}

	if m.currentField == FieldPassword {
		s += currentStyle.Render("Password: ") + strings.Repeat("*", len(m.input)) + "\n"
	} else if m.password != "" {
		s += "Password: " + strings.Repeat("*", len(m.password)) + "\n"
	} else {
		s += "Password: \n"
	}

	if m.currentField == FieldPassword {
		s += currentStyle.Render("Notes: ") + m.input + "\n"
	} else {
		s += "Notes: " + m.password + "\n"
	}

	s += "\nPress Enter to confirm, Esc to cancel\n"

	return s
}
