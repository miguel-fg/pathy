package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ConfirmationAction string

const ConfirmDelete ConfirmationAction = "delete"

type Confirmation struct {
	form   *huh.Form
	value  bool
	action ConfirmationAction
}

func NewConfirmation(width int, action ConfirmationAction, title string) *Confirmation {
	c := &Confirmation{
		action: action,
	}

	input := huh.NewConfirm().
		Title(title).
		Affirmative("Yes").
		Negative("No").
		Value(&c.value)

	group := huh.NewGroup(input)

	c.form = huh.NewForm(group).
		WithShowHelp(false).
		WithShowErrors(false).
		WithWidth(width - 8)

	return c
}

func (c *Confirmation) Init() tea.Cmd {
	return c.form.Init()
}

func (c *Confirmation) Update(msg tea.Msg) (bool, bool, tea.Cmd) {
	formModel, cmd := c.form.Update(msg)

	if f, ok := formModel.(*huh.Form); ok {
		c.form = f
	}

	switch c.form.State {
	case huh.StateCompleted:
		return true, c.value, cmd
	case huh.StateAborted:
		return true, false, cmd
	}

	return false, false, cmd
}

func (c *Confirmation) View() string {
	if fld := c.form.GetFocusedField(); fld != nil {
		return fld.View()
	}

	return c.form.View()
}

func (c *Confirmation) Action() ConfirmationAction {
	return c.action
}
