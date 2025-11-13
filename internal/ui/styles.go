package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	primaryColor   = lipgloss.Color("#00D9FF")
	secondaryColor = lipgloss.Color("#7D56F4")
	successColor   = lipgloss.Color("#04B575")
	warningColor   = lipgloss.Color("#FF9D00")
	errorColor     = lipgloss.Color("#FF4757")
	mutedColor     = lipgloss.Color("#6C7A89")
	highlightColor = lipgloss.Color("#FEE715")
	textColor      = lipgloss.Color("#FFFFFF")
	dimTextColor   = lipgloss.Color("#95A3B3")

	// Title style
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginBottom(1)

	// Header style
	headerStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			MarginBottom(1)

	// Domain styles
	domainCollapsedStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Bold(true)

	domainExpandedStyle = lipgloss.NewStyle().
				Foreground(successColor).
				Bold(true)

	domainSelectedStyle = lipgloss.NewStyle().
				Foreground(highlightColor).
				Bold(true).
				Background(lipgloss.Color("#2C3E50"))

	// Entry field styles
	usernameStyle = lipgloss.NewStyle().
			Foreground(primaryColor)

	passwordStyle = lipgloss.NewStyle().
			Foreground(warningColor)

	metadataStyle = lipgloss.NewStyle().
			Foreground(dimTextColor).
			Italic(true)

	// Tree drawing styles
	treeLineStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	// Help text style
	helpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			MarginTop(1).
			Padding(1, 2)

	// Error style
	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	// Count badge style
	countStyle = lipgloss.NewStyle().
			Foreground(dimTextColor).
			Italic(true)
)

