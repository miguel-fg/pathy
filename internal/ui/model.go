package ui

import (
	"os"

	"pathy/internal/fs"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
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

	// prompt state
	promptActive bool
	promptForm   *huh.Form
	promptValue  string
	promptAction string
}

func NewModel(startDir string) Model {
	return Model{cwd: startDir, history: fs.NewHistory()}
}

func (m Model) Init() tea.Cmd {
	return fs.LoadFiles(m.cwd)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.promptActive {
		formModel, cmd := m.promptForm.Update(msg)

		if f, ok := formModel.(*huh.Form); ok {
			m.promptForm = f
		}

		if m.promptForm.State == huh.StateCompleted {
			name := m.promptForm.GetString("name")

			m.promptActive = false

			switch m.promptAction {
			case "create":
				return m, fs.CreateFile(fs.Join(m.cwd, name))
			case "rename":
				if len(m.files) > 0 {
					old := fs.Join(m.cwd, m.files[m.cursor].Name())
					newp := fs.Join(m.cwd, name)
					return m, fs.RenameFile(old, newp)
				}
			}

			return m, nil
		}

		if m.promptForm.State == huh.StateAborted {
			m.promptActive = false
		}

		return m, cmd
	}

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
		case "a":
			m.promptValue = ""

			input := huh.NewInput().Key("name").Prompt("\uf054 ").Value(&m.promptValue).Title("Create new file").Placeholder("myfile.ext")
			group := huh.NewGroup(input)
			m.promptForm = huh.NewForm(group).WithShowHelp(false).WithShowErrors(false).WithWidth(m.width - 8)
			m.promptActive = true
			m.promptAction = "create"

			return m, m.promptForm.Init()
		case "d":
			if len(m.files) > 0 {
				target := fs.Join(m.cwd, m.files[m.cursor].Name())
				return m, fs.DeleteFile(target)
			}
		case "r":
			if len(m.files) > 0 {
				oldname := m.files[m.cursor].Name()
				m.promptValue = oldname
				input := huh.NewInput().Key("name").Prompt("\uf054 ").Value(&m.promptValue).Title("Rename file").Placeholder("myfile.ext")
				group := huh.NewGroup(input)
				m.promptForm = huh.NewForm(group).WithShowHelp(false).WithShowErrors(false).WithWidth(m.width - 8)
				m.promptActive = true
				m.promptAction = "rename"
				return m, m.promptForm.Init()
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

	case fs.FileOpMsg:
		if msg.Err != nil {
			m.err = msg.Err
		} else {
			return m, fs.LoadFiles(m.cwd)
		}
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

	if m.promptActive && m.promptForm != nil {
		if fld := m.promptForm.GetFocusedField(); fld != nil {
			s += "\n" + fld.View() + "\n"
		} else {
			s += "\n" + m.promptForm.View() + "\n"
		}
	}

	appFrame := styles.Border.Width(m.width-4).Height(m.height-4).Margin(1).Padding(1, 2)

	content := appFrame.Render(s)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
