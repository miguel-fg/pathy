package fs

import (
	"os"
	"path/filepath"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
)

type FilesLoadedMsg struct {
	Dir     string
	Entries []os.DirEntry
}

type ParentLoadedMsg struct {
	Dir     string
	Entries []os.DirEntry
}

type ErrMsg struct {
	Dir string
	Err error
}

func LoadFiles(dir string) tea.Cmd {
	return func() tea.Msg {
		files, err := os.ReadDir(dir)
		if err != nil {
			return ErrMsg{Dir: dir, Err: err}
		}
		return FilesLoadedMsg{Dir: dir, Entries: files}
	}
}

func LoadParent(dir string) tea.Cmd {
	return func() tea.Msg {
		files, err := os.ReadDir(dir)
		if err != nil {
			return ErrMsg{Dir: dir, Err: err}
		}
		return ParentLoadedMsg{Dir: dir, Entries: files}
	}
}

func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func HomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}
