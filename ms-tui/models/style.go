package models

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

func applyStyle(widgetName, content string, width, height int) string {
	// Create Footer
	ribbonStyleLeft := lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Width(width / 3).
		PaddingLeft(2).
		AlignHorizontal(lipgloss.Left).
		SetString("widget : " + widgetName)

	ribbonStyleCenter := lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Width(width / 3).
		AlignHorizontal(lipgloss.Center).
		SetString(time.Now().Format("02/01/2006 - 15:04"))

	ribbonStyleRight := lipgloss.NewStyle().
		Background(lipgloss.Color("240")).
		Width(width / 3).
		PaddingRight(2).
		AlignHorizontal(lipgloss.Right).
		SetString("MarketSentinel v0.0.1")

	contentStyled := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)

	footerStyled := lipgloss.JoinHorizontal(lipgloss.Right,
		ribbonStyleLeft.Render(),
		ribbonStyleCenter.Render(),
		ribbonStyleRight.Render(),
	)

	return lipgloss.JoinVertical(lipgloss.Top, contentStyled, footerStyled)

}
