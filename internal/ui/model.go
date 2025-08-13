package ui

import (
	"fmt"
	"os"
	"strings"

	"pathy/internal/fs"

	"github.com/charmbracelet/bubbles/viewport"
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

	activePrompt       *Prompt
	activeConfirmation *Confirmation

	viewport viewport.Model
}

func NewModel(startDir string) Model {
	vp := viewport.New(0, 0)

	return Model{cwd: startDir, history: fs.NewHistory(startDir), viewport: vp}
}

func (m Model) Init() tea.Cmd {
	return fs.LoadFiles(m.cwd)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.activePrompt != nil {
		done, val, cmd := m.activePrompt.Update(msg)

		if done {

			var finishCmd = cmd
			action := m.activePrompt.Action()
			m.activePrompt = nil

			switch action {
			case PromptCreate:
				if val != "" {
					return m, tea.Batch(finishCmd, fs.CreateFile(fs.Join(m.cwd, val)))
				}
			case PromptRename:
				if val != "" && len(m.files) > 0 {
					old := fs.Join(m.cwd, m.files[m.cursor].Name())
					newp := fs.Join(m.cwd, val)
					return m, tea.Batch(finishCmd, fs.RenameFile(old, newp))
				}
			}
			return m, finishCmd
		}
		return m, cmd
	}

	if m.activeConfirmation != nil {
		done, val, cmd := m.activeConfirmation.Update(msg)

		if done {
			var finishCmd = cmd
			action := m.activeConfirmation.Action()
			m.activeConfirmation = nil

			switch action {
			case ConfirmDelete:
				if val && len(m.files) > 0 {
					target := fs.Join(m.cwd, m.files[m.cursor].Name())
					return m, tea.Batch(finishCmd, fs.DeleteFile(target))
				}
			}
			return m, finishCmd
		}
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.viewport.Width = msg.Width - 8
		m.viewport.Height = msg.Height - 12
		return m, nil

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
			m.activePrompt = NewPrompt(m.width, PromptCreate, "Create new file", "my_file.ext", "")
			return m, m.activePrompt.Init()
		case "d":
			if len(m.files) > 0 {
				m.activeConfirmation = NewConfirmation(m.width, ConfirmDelete, "Confirm file delete")
			}
		case "r":
			if len(m.files) > 0 {
				m.activePrompt = NewPrompt(m.width, PromptRename, "Rename file", "my_file.ext", m.files[m.cursor].Name())
				return m, m.activePrompt.Init()
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

	var fileContent strings.Builder
	for i, f := range m.files {
		name := f.Name()

		if f.IsDir() {
			name = styles.Dir.Render(name + "/")
		}

		cursor := " "
		if i == m.cursor {
			cursor = styles.Cursor.Render("→ ")
		}

		fileContent.WriteString(cursor + name + "\n")
	}

	if len(m.files) == 0 {
		fileContent.WriteString("(empty directory)\n")
	}

	m.viewport.SetContent(fileContent.String())

	if m.cursor >= m.viewport.YOffset+m.viewport.Height-1 {
		m.viewport.YOffset = m.cursor - m.viewport.Height + 1
	} else if m.cursor < m.viewport.YOffset {
		m.viewport.YOffset = m.cursor
	}

	title := styles.Title.Render("Pathy v0.0.1 — "+m.cwd) + "\n\n"
	viewportContent := m.viewport.View()

	var s strings.Builder
	s.WriteString(title)
	s.WriteString(viewportContent)

	if len(m.files) > m.viewport.Height {
		scrollInfo := fmt.Sprintf(" (%d-%d of %d)", m.viewport.YOffset+1, min(m.viewport.YOffset+m.viewport.Height), len(m.files))
		s.WriteString("\n" + styles.Subtle.Render(scrollInfo))
	}

	if m.activePrompt != nil {
		s.WriteString("\n" + m.activePrompt.View() + "\n")
	}

	if m.activeConfirmation != nil {
		s.WriteString("\n" + m.activeConfirmation.View() + "\n")
	}

	appFrame := styles.Border.Width(m.width-4).Height(m.height-4).Margin(1).Padding(1, 2)

	content := appFrame.Render(s.String())

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
