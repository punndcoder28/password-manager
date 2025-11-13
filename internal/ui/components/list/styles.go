package list

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/punndcoder28/password-manager/internal/ui/common"
)

// List component-specific styles
// These styles are specific to the tree-based password list view

// Domain node styles
var (
	domainCollapsedStyle = lipgloss.NewStyle().
				Foreground(common.SecondaryColor).
				Bold(true)

	domainExpandedStyle = lipgloss.NewStyle().
				Foreground(common.SuccessColor).
				Bold(true)

	domainSelectedStyle = lipgloss.NewStyle().
				Foreground(common.HighlightColor).
				Bold(true).
				Background(common.SelectedBgColor)
)

// Entry field styles
var (
	usernameStyle = lipgloss.NewStyle().
			Foreground(common.PrimaryColor)

	passwordStyle = lipgloss.NewStyle().
			Foreground(common.WarningColor)

	treeLineStyle = lipgloss.NewStyle().
			Foreground(common.MutedColor)
)

// Count badge style
var countStyle = lipgloss.NewStyle().
	Foreground(common.DimTextColor).
	Italic(true)

