package common

import (
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// FormatTimeAgo formats a time as a human-readable "time ago" string
// Example: "2 days ago", "1 week ago"
func FormatTimeAgo(t time.Time) string {
	if t.IsZero() {
		return "Unknown"
	}
	return humanize.Time(t)
}

// Pluralize returns singular or plural form based on count
// Example: Pluralize(1, "entry", "entries") -> "entry"
//          Pluralize(5, "entry", "entries") -> "entries"
func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

// MaskPassword returns a string of asterisks matching the password length
// Example: MaskPassword("mypass") -> "******"
func MaskPassword(password string) string {
	return strings.Repeat("*", len(password))
}

// TruncateString truncates a string to a maximum length and adds ellipsis
// Example: TruncateString("very long string", 10) -> "very lo..."
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

// Icons contains commonly used unicode icons
var Icons = struct {
	Lock       string
	Key        string
	User       string
	Calendar   string
	Clock      string
	Search     string
	Check      string
	Cross      string
	Warning    string
	Info       string
	ArrowRight string
	ArrowDown  string
	Cursor     string
}{
	Lock:       "🔐",
	Key:        "🔑",
	User:       "👤",
	Calendar:   "📅",
	Clock:      "🕒",
	Search:     "🔍",
	Check:      "✅",
	Cross:      "❌",
	Warning:    "⚠️",
	Info:       "ℹ️",
	ArrowRight: "▶",
	ArrowDown:  "▼",
	Cursor:     "❯",
}

