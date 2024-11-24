package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type NewsWidgetModel struct {
	widgetName string
	content    string
	quitting   bool
}

func NewNewsWidgetModel() *NewsWidgetModel {

	model := NewsWidgetModel{
		widgetName: "news",
		content:    "News Screen",
		quitting:   false,
	}
	return &model
}

func (m *NewsWidgetModel) Init() tea.Cmd {
	return nil
}

func (m *NewsWidgetModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *NewsWidgetModel) View() string {
	if m.quitting {
		return ""
	}

	output := fmt.Sprintf(
		"%s",
		applyStyle(m.widgetName, m.content, terminalWidth, terminalHeight),
	)

	return output
}
