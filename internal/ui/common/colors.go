package common

import "github.com/charmbracelet/lipgloss"

// Color palette for the entire application
// These colors maintain consistency across all UI components
var (
	// Primary colors
	PrimaryColor   = lipgloss.Color("#00D9FF") // Cyan - main interactive elements
	SecondaryColor = lipgloss.Color("#7D56F4") // Purple - secondary elements
	SuccessColor   = lipgloss.Color("#04B575") // Green - success states
	WarningColor   = lipgloss.Color("#FF9D00") // Orange - warnings/passwords
	ErrorColor     = lipgloss.Color("#FF4757") // Red - errors
	HighlightColor = lipgloss.Color("#FEE715") // Yellow - highlights/selections

	// Text colors
	TextColor    = lipgloss.Color("#FFFFFF") // White - primary text
	DimTextColor = lipgloss.Color("#95A3B3") // Gray - secondary text
	MutedColor   = lipgloss.Color("#6C7A89") // Dark gray - muted elements

	// Background colors
	SelectedBgColor = lipgloss.Color("#2C3E50") // Dark blue-gray - selected item background
)

