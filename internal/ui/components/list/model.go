package list

import (
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/punndcoder28/password-manager/vault"
)

// TreeNode represents a node in the tree structure
type TreeNode struct {
	Domain   string
	Entries  []vault.Entry
	Expanded bool
}

// Model represents the Bubble Tea model for the password list component
type Model struct {
	entries         map[string][]vault.Entry
	tree            []TreeNode
	cursor          int
	selectedDomain  string
	selectedEntry   int
	revealPasswords map[string]map[int]bool // domain -> entry index -> revealed
	err             error
	width           int
	height          int
}

// New creates a new list model with the provided password entries
func New(entries map[string][]vault.Entry) Model {
	tree := buildTree(entries)

	return Model{
		entries:         entries,
		tree:            tree,
		cursor:          0,
		selectedDomain:  "",
		selectedEntry:   -1,
		revealPasswords: make(map[string]map[int]bool),
		width:           80,
		height:          24,
	}
}

// buildTree creates a sorted tree structure from the password entries
func buildTree(entries map[string][]vault.Entry) []TreeNode {
	tree := make([]TreeNode, 0, len(entries))

	// Get sorted domain names
	domains := make([]string, 0, len(entries))
	for domain := range entries {
		domains = append(domains, domain)
	}
	sort.Strings(domains)

	// Build tree nodes
	for _, domain := range domains {
		tree = append(tree, TreeNode{
			Domain:   domain,
			Entries:  entries[domain],
			Expanded: false,
		})
	}

	return tree
}

// Init initializes the model (required by Bubble Tea)
func (m Model) Init() tea.Cmd {
	return nil
}

// getTotalEntryCount returns the total number of password entries
func (m Model) getTotalEntryCount() int {
	count := 0
	for _, entries := range m.entries {
		count += len(entries)
	}
	return count
}

// isPasswordRevealed checks if a password for a specific domain and entry index is revealed
func (m Model) isPasswordRevealed(domain string, entryIndex int) bool {
	if domainMap, exists := m.revealPasswords[domain]; exists {
		return domainMap[entryIndex]
	}
	return false
}

// togglePasswordReveal toggles the reveal state of a password
func (m *Model) togglePasswordReveal(domain string, entryIndex int) {
	if m.revealPasswords[domain] == nil {
		m.revealPasswords[domain] = make(map[int]bool)
	}
	m.revealPasswords[domain][entryIndex] = !m.revealPasswords[domain][entryIndex]
}

