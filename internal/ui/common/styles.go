package common

import "github.com/charmbracelet/lipgloss"

// Common styles that can be reused across all UI components
// These provide a consistent look and feel throughout the application

// TitleStyle is used for main titles/headers
var TitleStyle = lipgloss.NewStyle().
	Foreground(PrimaryColor).
	Bold(true).
	MarginBottom(1)

// HeaderStyle is used for section headers
var HeaderStyle = lipgloss.NewStyle().
	Foreground(MutedColor).
	Italic(true).
	MarginBottom(1)

// ErrorStyle is used for error messages
var ErrorStyle = lipgloss.NewStyle().
	Foreground(ErrorColor).
	Bold(true)

// SuccessStyle is used for success messages
var SuccessStyle = lipgloss.NewStyle().
	Foreground(SuccessColor).
	Bold(true)

// HelpStyle is used for help text and keyboard shortcuts
var HelpStyle = lipgloss.NewStyle().
	Foreground(MutedColor).
	MarginTop(1).
	Padding(1, 2)

// HighlightedStyle is used for highlighted/selected items
var HighlightedStyle = lipgloss.NewStyle().
	Foreground(HighlightColor).
	Bold(true).
	Background(SelectedBgColor)

// MetadataStyle is used for secondary information
var MetadataStyle = lipgloss.NewStyle().
	Foreground(DimTextColor).
	Italic(true)

// SubtleStyle is used for subtle, less important text
var SubtleStyle = lipgloss.NewStyle().
	Foreground(MutedColor)

