package ui

import "github.com/charmbracelet/lipgloss"

var styles = struct {
	Title       lipgloss.Style
	Cursor      lipgloss.Style
	Dir         lipgloss.Style
	Border      lipgloss.Style
	Subtle      lipgloss.Style
	SideEntry   lipgloss.Style
	SideDir     lipgloss.Style
	BorderSides lipgloss.Style
}{
	Title:     lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")),
	Cursor:    lipgloss.NewStyle().Foreground(lipgloss.Color("212")),
	Dir:       lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Bold(true),
	Border:    lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("206")),
	Subtle:    lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
	SideEntry: lipgloss.NewStyle().Foreground(lipgloss.Color("250")),
	SideDir:   lipgloss.NewStyle().Foreground(lipgloss.Color("27")).Bold(true), BorderSides: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("237")),
}
