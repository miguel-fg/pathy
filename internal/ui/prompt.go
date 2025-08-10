package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type PromptAction string

const (
	PromptCreate PromptAction = "create"
	PromptRename PromptAction = "rename"
)

type Prompt struct {
	form   *huh.Form
	value  string
	action PromptAction
}

func NewPrompt(width int, action PromptAction, title, placeholder, initialValue string) *Prompt {
	p := &Prompt{
		value:  initialValue,
		action: action,
	}

	input := huh.NewInput().
		Key("name").
		Prompt("\uf054 ").
		Value(&p.value).
		Title(title).
		Placeholder(placeholder)

	group := huh.NewGroup(input)

	p.form = huh.NewForm(group).
		WithShowHelp(false).
		WithShowErrors(false).
		WithWidth(width - 8)

	return p
}

func (p *Prompt) Init() tea.Cmd {
	return p.form.Init()
}

func (p *Prompt) Update(msg tea.Msg) (bool, string, tea.Cmd) {
	formModel, cmd := p.form.Update(msg)

	if f, ok := formModel.(*huh.Form); ok {
		p.form = f
	}

	switch p.form.State {
	case huh.StateCompleted:
		return true, p.value, cmd
	case huh.StateAborted:
		return true, "", cmd
	}

	return false, "", cmd
}

func (p *Prompt) View() string {
	if fld := p.form.GetFocusedField(); fld != nil {
		return fld.View()
	}

	return p.form.View()
}

func (p *Prompt) Action() PromptAction {
	return p.action
}
