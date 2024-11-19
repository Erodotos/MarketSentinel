package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type mainSessionState int

const (
	welcome mainSessionState = iota
	welcometwo
)

type MainModel struct {
	state      int
	router     textinput.Model
	showRouter bool
	models     []tea.Model
	quitting   bool
	loaded     bool
}

func NewMainModel() *MainModel {
	ti := textinput.New()
	ti.Placeholder = "Select widget here ...."
	ti.Focus()

	model := MainModel{
		state:  1,
		router: ti,
		models: []tea.Model{NewWelcomeModel(), NewWelcomeTwoModel()},
	}
	return &model
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.showRouter {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch key := msg.String(); key {
			case "enter":
				m.state, _ = strconv.Atoi(m.router.Value())
				m.router.SetValue("")
				m.showRouter = false
			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			default:
				m.router, cmd = m.router.Update(msg)
			}

		}
	} else {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.loaded = true
		case tea.KeyMsg:
			switch key := msg.String(); key {
			case "/":
				m.showRouter = true
				return m, tea.Batch(cmds...)
			}

		}

		fmt.Println(m.state - 1)

		newModel, newCmd := m.models[m.state-1].Update(msg)
		m.models[m.state-1] = newModel
		cmd = newCmd
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	if m.quitting {
		return ""
	}
	if !m.loaded {
		return "loading..."
	}
	if m.showRouter {
		return m.router.View()
	} else {
		return m.models[m.state-1].View()
	}
}
