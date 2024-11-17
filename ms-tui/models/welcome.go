package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type WelcomeModel struct {
	content  string
	quitting bool
}

func NewWelcomeModel() *WelcomeModel {

	model := WelcomeModel{
		content:  "welcome",
		quitting: false,
	}
	return &model
}

func (m WelcomeModel) Init() tea.Cmd {
	return nil
}

func (m WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m WelcomeModel) View() string {
	if m.quitting {
		return ""
	}

	output := fmt.Sprintf(
		"%s",
		m.content)

	return output
}
