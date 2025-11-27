package list

import (
	"fmt"
	"strings"

	"github.com/punndcoder28/password-manager/internal/ui/common"
	vaultPackage "github.com/punndcoder28/password-manager/internal/vault"
)

// View renders the list UI
func (m Model) View() string {
	if m.err != nil {
		return common.ErrorStyle.Render(fmt.Sprintf("Error: %v", m.err))
	}

	var s strings.Builder

	// Title
	totalEntries := m.getTotalEntryCount()
	title := fmt.Sprintf("%s Password Vault (%d %s)",
		common.Icons.Lock,
		totalEntries,
		common.Pluralize(totalEntries, "entry", "entries"))
	s.WriteString(common.TitleStyle.Render(title))
	s.WriteString("\n\n")

	// Render tree
	currentPos := 0
	for i, node := range m.tree {
		// Render domain
		s.WriteString(m.renderDomain(node, i, currentPos))
		currentPos++

		// Render entries if expanded
		if node.Expanded {
			for entryIdx, entry := range node.Entries {
				isLast := entryIdx == len(node.Entries)-1
				s.WriteString(m.renderEntry(node.Domain, entry, entryIdx, currentPos, isLast))
				currentPos++
			}
		}
	}

	// Help text
	s.WriteString("\n")
	s.WriteString(common.HelpStyle.Render(getHelpText()))

	return s.String()
}

// renderDomain renders a domain node in the tree
func (m Model) renderDomain(node TreeNode, nodeIndex int, position int) string {
	var s strings.Builder

	// Determine if this node is selected
	isSelected := m.cursor == position

	// Expand/collapse icon
	icon := common.Icons.ArrowRight
	if node.Expanded {
		icon = common.Icons.ArrowDown
	}

	// Entry count
	entryCount := len(node.Entries)
	countText := fmt.Sprintf("(%d %s)", entryCount, common.Pluralize(entryCount, "account", "accounts"))

	// Build the domain line
	domainText := fmt.Sprintf("%s %s %s", icon, node.Domain, countText)

	// Apply style based on selection and expansion state
	var styledText string
	if isSelected {
		styledText = domainSelectedStyle.Render(domainText)
		s.WriteString(common.Icons.Cursor + " ")
	} else {
		s.WriteString("  ")
		if node.Expanded {
			styledText = domainExpandedStyle.Render(domainText)
		} else {
			styledText = domainCollapsedStyle.Render(domainText)
		}
	}

	s.WriteString(styledText)
	s.WriteString("\n")

	return s.String()
}

// renderEntry renders a password entry in the tree
func (m Model) renderEntry(domain string, entry vaultPackage.Entry, entryIdx int, position int, isLast bool) string {
	var s strings.Builder

	// Determine if this entry is selected
	isSelected := m.cursor == position

	// Tree drawing characters
	var prefix string
	if isLast {
		prefix = "  └─"
	} else {
		prefix = "  ├─"
	}

	// Selection indicator
	if isSelected {
		s.WriteString(common.Icons.Cursor + " ")
	} else {
		s.WriteString("  ")
	}

	// Entry header (username)
	s.WriteString(treeLineStyle.Render(prefix))
	s.WriteString(" ")
	s.WriteString(usernameStyle.Render(fmt.Sprintf("%s %s", common.Icons.User, entry.Username)))
	s.WriteString("\n")

	// Determine line prefix for nested items
	var nestedPrefix string
	if isLast {
		nestedPrefix = "     "
	} else {
		nestedPrefix = "  │  "
	}

	// Selection padding
	selectionPadding := "  "

	// Password field
	password := m.formatPassword(domain, entry.Password, entryIdx)
	s.WriteString(selectionPadding)
	s.WriteString(treeLineStyle.Render(nestedPrefix))
	s.WriteString(" ")
	s.WriteString(passwordStyle.Render(fmt.Sprintf("%s %s", common.Icons.Key, password)))
	s.WriteString("\n")

	// Created date
	createdText := fmt.Sprintf("%s Created: %s", common.Icons.Calendar, common.FormatTimeAgo(entry.CreatedAt))
	s.WriteString(selectionPadding)
	s.WriteString(treeLineStyle.Render(nestedPrefix))
	s.WriteString(" ")
	s.WriteString(common.MetadataStyle.Render(createdText))
	s.WriteString("\n")

	// Updated date
	updatedText := fmt.Sprintf("%s Updated: %s", common.Icons.Clock, common.FormatTimeAgo(entry.UpdatedAt))
	s.WriteString(selectionPadding)
	s.WriteString(treeLineStyle.Render(nestedPrefix))
	s.WriteString(" ")
	s.WriteString(common.MetadataStyle.Render(updatedText))
	s.WriteString("\n")

	return s.String()
}

// formatPassword formats the password, either masked or revealed
func (m Model) formatPassword(domain string, password string, entryIdx int) string {
	if m.isPasswordRevealed(domain, entryIdx) {
		return password
	}
	return common.MaskPassword(password)
}

// getHelpText returns the help text for keyboard shortcuts
func getHelpText() string {
	return "[↑/↓ or j/k: navigate] [Enter/Space: expand/toggle] [r: reveal password] [R: reveal all] [q: quit]"
}
