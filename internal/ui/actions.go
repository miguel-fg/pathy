package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type GoBackMsg struct {
	Dir string
}

func (m Model) GoBack() tea.Cmd {
	return func() tea.Msg {
		if dir, ok := m.history.Pop(); ok {
			return GoBackMsg{Dir: dir}
		}

		return nil
	}
}
