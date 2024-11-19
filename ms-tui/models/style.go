package models

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

func footer(width, height int) string {
	// Create the ribbon style
	ribbonStyleLeft := lipgloss.NewStyle().
		Background(lipgloss.Color("240")). // Light grey background (you can adjust the color code)
		Width(width / 3).                  // Full terminal width
		PaddingLeft(1).
		AlignHorizontal(lipgloss.Left).
		SetString("widget:sentiment")

	ribbonStyleCenter := lipgloss.NewStyle().
		Background(lipgloss.Color("240")). // Light grey background (you can adjust the color code)
		Width(width / 3).                  // Full terminal width
		PaddingLeft(1).
		AlignHorizontal(lipgloss.Center).
		SetString(time.Now().Format("02/01/2006 - 15:04"))

	ribbonStyleRight := lipgloss.NewStyle().
		Background(lipgloss.Color("240")). // Light grey background (you can adjust the color code)
		Width(width / 3).                  // Full terminal width
		PaddingRight(1).
		AlignHorizontal(lipgloss.Right).
		SetString("MarketSentinel v0.0.1")

	// Create some content above the ribbon (optional)
	// content := "This is some content above the ribbon.\nYou can add more lines here."
	// contentStyle := lipgloss.NewStyle().
	// 	Height(20 - 3). // Adjust height to leave space for the ribbon (and a small buffer)
	// 	MaxWidth(100).
	// 	Align(lipgloss.Center)

	// styledContent := contentStyle.Render(content)

	// Create the ribbon text (optional)
	// ribbonText := "This is the ribbon at the bottom"
	return lipgloss.JoinHorizontal(lipgloss.Right, ribbonStyleLeft.Render(), ribbonStyleCenter.Render(), ribbonStyleRight.Render())

}
