package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Widget interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type CoreModel struct {
	command textinput.Model
	widget  Widget
}

func NewCoreModel() CoreModel {
	ti := textinput.New()
	ti.Placeholder = "Type your command"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return CoreModel{
		command: ti,
		widget:  NewWelcomeModel(),
	}
}

func (m CoreModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Executes given command
		case "enter":
			m.command.SetValue("")
		// Exits the application
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.command, cmd = m.command.Update(msg)

	return m, cmd
}

func (m CoreModel) View() string {
	return fmt.Sprintf(
		"%s\n%s",
		m.command.View(),
		m.widget.View(),
	)
}

func main() {
	if _, err := tea.NewProgram(NewCoreModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
