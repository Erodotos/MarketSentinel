package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type WelcomeModel struct {
	content  string
	tWidth   int
	tHeight  int
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
	case tea.WindowSizeMsg:
		m.tWidth, m.tHeight = msg.Width, msg.Height
		return m, tea.Batch(cmds...)
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
		"%s \n %s",
		m.content, footer(m.tWidth, m.tHeight))

	return output
}
