package ui

import (
	"os"

	"pathy/internal/fs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	cwd     string
	files   []os.DirEntry
	history *fs.History
	cursor  int
	width   int
	height  int
	err     error
}

func NewModel(startDir string) Model {
	return Model{cwd: startDir, history: fs.NewHistory()}
}

func (m Model) Init() tea.Cmd {
	return fs.LoadFiles(m.cwd)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.files)-1 {
				m.cursor++
			}
		case "enter", "l":
			if len(m.files) > 0 && m.files[m.cursor].IsDir() {
				m.history.Push(m.cwd)
				newPath := fs.Join(m.cwd, m.files[m.cursor].Name())
				m.cwd = newPath
				return m, fs.LoadFiles(m.cwd)
			}
		case "backspace", "h":
			if m.history.HasPrevious() {
				return m, m.GoBack()
			}
		}

	case GoBackMsg:
		m.cwd = msg.Dir
		return m, fs.LoadFiles(m.cwd)

	case fs.FilesLoadedMsg:
		if msg.Dir != m.cwd {
			return m, nil
		}

		m.files = msg.Entries
		m.cursor = 0

	case fs.ErrMsg:
		m.err = msg.Err
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	s := styles.Title.Render("Pathy v0.0.1 — "+m.cwd) + "\n\n"

	for i, f := range m.files {
		name := f.Name()

		if f.IsDir() {
			name = styles.Dir.Render(name + "/")
		}

		cursor := " "
		if i == m.cursor {
			cursor = styles.Cursor.Render("→ ")
		}
		s += cursor + name + "\n"
	}

	if len(m.files) == 0 {
		s += "(empty directory)\n"
	}

	appFrame := styles.Border.Width(m.width-4).Height(m.height-4).Margin(1).Padding(1, 2)

	content := appFrame.Render(s)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
