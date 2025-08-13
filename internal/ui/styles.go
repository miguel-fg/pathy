package ui

import "github.com/charmbracelet/lipgloss"

var styles = struct {
	Title  lipgloss.Style
	Cursor lipgloss.Style
	Dir    lipgloss.Style
	Border lipgloss.Style
	Subtle lipgloss.Style
}{
	Title:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")),
	Cursor: lipgloss.NewStyle().Foreground(lipgloss.Color("212")),
	Dir:    lipgloss.NewStyle().Foreground(lipgloss.Color("33")).Bold(true),
	Border: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderBottomForeground(lipgloss.Color("212")).Margin(1).Padding(1, 2),
	Subtle: lipgloss.NewStyle().Foreground(lipgloss.Color("237")),
}
