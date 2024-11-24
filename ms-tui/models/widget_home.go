package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	widgetName string
	content    string
	quitting   bool
}

func NewHomeModel() *HomeModel {

	model := HomeModel{
		widgetName: "home",
		content:    "Home Screen",
		quitting:   false,
	}
	return &model
}

func (m *HomeModel) Init() tea.Cmd {
	return nil
}

func (m *HomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *HomeModel) View() string {
	if m.quitting {
		return ""
	}

	output := fmt.Sprintf(
		"%s",
		applyStyle(m.widgetName, m.content, terminalWidth, terminalHeight),
	)

	return output
}
