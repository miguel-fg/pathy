package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	Label string
	Value string
	done  bool
}

type PromptSubmitMsg struct {
	Text string
}

func NewPrompt(label string) Prompt {
	return Prompt{Label: label}
}

func (p Prompt) Init() tea.Cmd {
	return nil
}

func (p Prompt) Update(msg tea.Msg) (Prompt, tea.Cmd) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.Type {
		case tea.KeyEsc:
			return p, func() tea.Msg { return PromptSubmitMsg{Text: ""} }
		case tea.KeyEnter:
			return p, func() tea.Msg { return PromptSubmitMsg{Text: p.Value} }
		case tea.KeyBackspace:
			if len(p.Value) > 0 {
				p.Value = p.Value[:len(p.Value)-1]
			}
		default:
			if len(m.String()) == 1 {
				p.Value += m.String()
			}
		}
	}

	return p, nil
}

func (p Prompt) View() string {
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	inputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	return labelStyle.Render(p.Label) + ": " + inputStyle.Render(p.Value) + "_"
}
