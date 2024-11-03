package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type WelcomeModel struct {
	viewport viewport.Model
}

func NewWelcomeModel() WelcomeModel {
	vp := viewport.New(30, 5)
	vp.SetContent("Welcome to MarketSentinel Terminal")

	return WelcomeModel{
		viewport: vp,
	}
}

func (m WelcomeModel) Init() tea.Cmd {
	return nil
}

func (m WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = (msg.Width)
		fmt.Println(m.viewport.Width)
		return m, nil
	default:
		return m, nil
	}
}

func (m WelcomeModel) View() string {

	return fmt.Sprintf(
		"%s",
		m.viewport.View(),
	)
}
