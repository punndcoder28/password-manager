package list

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles all the state changes based on messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.moveCursorUp()
			return m, nil

		case "down", "j":
			m.moveCursorDown()
			return m, nil

		case "enter", "space":
			m.toggleExpand()
			return m, nil

		case "r":
			m.togglePasswordRevealAtCursor()
			return m, nil

		case "R":
			m.toggleAllPasswordsReveal()
			return m, nil
		}
	}

	return m, nil
}

// moveCursorUp moves the cursor up in the tree
func (m *Model) moveCursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

// moveCursorDown moves the cursor down in the tree
func (m *Model) moveCursorDown() {
	maxCursor := m.getMaxCursorPosition()
	if m.cursor < maxCursor {
		m.cursor++
	}
}

// getMaxCursorPosition calculates the maximum cursor position based on expanded nodes
func (m *Model) getMaxCursorPosition() int {
	position := 0
	for _, node := range m.tree {
		position++ // Count the domain itself
		if node.Expanded {
			position += len(node.Entries) // Count the entries if expanded
		}
	}
	return position - 1 // Convert to 0-based index
}

// toggleExpand toggles the expansion state of the currently selected domain
func (m *Model) toggleExpand() {
	currentPos := 0
	for i := range m.tree {
		if currentPos == m.cursor {
			// Cursor is on a domain node
			m.tree[i].Expanded = !m.tree[i].Expanded
			return
		}
		currentPos++

		// If this domain is expanded, check if cursor is on one of its entries
		if m.tree[i].Expanded {
			for entryIdx := range m.tree[i].Entries {
				if currentPos == m.cursor {
					// Cursor is on an entry, toggle password reveal
					m.togglePasswordReveal(m.tree[i].Domain, entryIdx)
					return
				}
				currentPos++
			}
		}
	}
}

// togglePasswordRevealAtCursor toggles password reveal for the entry at cursor
func (m *Model) togglePasswordRevealAtCursor() {
	currentPos := 0
	for i := range m.tree {
		if currentPos == m.cursor {
			// Cursor is on a domain node, do nothing
			return
		}
		currentPos++

		// If this domain is expanded, check if cursor is on one of its entries
		if m.tree[i].Expanded {
			for entryIdx := range m.tree[i].Entries {
				if currentPos == m.cursor {
					// Cursor is on an entry, toggle password reveal
					m.togglePasswordReveal(m.tree[i].Domain, entryIdx)
					return
				}
				currentPos++
			}
		}
	}
}

// toggleAllPasswordsReveal toggles reveal state for all passwords
func (m *Model) toggleAllPasswordsReveal() {
	// Check if any password is currently revealed
	anyRevealed := false
	for _, domainMap := range m.revealPasswords {
		for _, revealed := range domainMap {
			if revealed {
				anyRevealed = true
				break
			}
		}
		if anyRevealed {
			break
		}
	}

	// If any are revealed, hide all. Otherwise, reveal all.
	if anyRevealed {
		m.revealPasswords = make(map[string]map[int]bool)
	} else {
		for _, node := range m.tree {
			if m.revealPasswords[node.Domain] == nil {
				m.revealPasswords[node.Domain] = make(map[int]bool)
			}
			for i := range node.Entries {
				m.revealPasswords[node.Domain][i] = true
			}
		}
	}
}

