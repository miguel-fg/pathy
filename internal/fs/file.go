package fs

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type FileOpMsg struct {
	Success bool
	Err     error
}

func CreateFile(path string) tea.Cmd {
	return func() tea.Msg {
		f, err := os.Create(path)
		if err != nil {
			return FileOpMsg{Success: false, Err: err}
		}

		defer f.Close()

		return FileOpMsg{Success: true}
	}
}

func RenameFile(oldPath, newPath string) tea.Cmd {
	return func() tea.Msg {
		err := os.Rename(oldPath, newPath)
		if err != nil {
			return FileOpMsg{Success: false, Err: err}
		}
		return FileOpMsg{Success: true}
	}
}

func DeleteFile(path string) tea.Cmd {
	return func() tea.Msg {
		err := os.RemoveAll(path)
		if err != nil {
			return FileOpMsg{Success: false, Err: err}
		}
		return FileOpMsg{Success: true}
	}
}
