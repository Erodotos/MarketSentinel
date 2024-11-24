package models

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
)

// Global Variables
var terminalWidth, terminalHeight int
var currentTime string

type tickMsg time.Time

type RootModel struct {
	dump         io.Writer
	currentModel string
	router       textinput.Model
	showRouter   bool
	models       map[string]tea.Model
	quitting     bool
	loaded       bool
}

func NewRootModel() *RootModel {

	var dump *os.File

	var err error
	dump, err = os.OpenFile("messages.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
	if err != nil {
		os.Exit(1)
	}

	ti := textinput.New()
	ti.Placeholder = "Select widget here ...."
	ti.Focus()

	model := RootModel{
		dump:         dump,
		currentModel: "home",
		router:       ti,
		models: map[string]tea.Model{
			"home": NewHomeModel(),
			"news": NewNewsWidgetModel()},
	}
	return &model
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *RootModel) Init() tea.Cmd {
	return tickCmd()
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	if m.dump != nil {
		spew.Fdump(m.dump, msg)
	}

	currentTime = time.Now().Format("02/01/2006 - 15:04")

	if m.showRouter {
		switch msg := msg.(type) {
		case tickMsg:
			cmd = tickCmd()
			cmds = append(cmds, cmd)
		case tea.KeyMsg:
			switch key := msg.String(); key {
			case "enter":
				m.currentModel = m.router.Value()
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
		case tickMsg:
			cmd = tickCmd()
			cmds = append(cmds, cmd)
		case tea.WindowSizeMsg:
			terminalWidth = msg.Width
			terminalHeight = msg.Height
			m.loaded = true
		case tea.KeyMsg:
			switch key := msg.String(); key {
			case "/":
				m.showRouter = true
				return m, tea.Batch(cmds...)
			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			}

		}

		newModel, newCmd := m.models[m.currentModel].Update(msg)
		m.models[m.currentModel] = newModel
		cmds = append(cmds, newCmd)
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *RootModel) View() string {
	if m.quitting {
		return ""
	}
	if !m.loaded {
		return "loading..."
	}
	if m.showRouter {
		return m.router.View()
	} else {
		return m.models[m.currentModel].View()
	}
}
