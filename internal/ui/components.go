package ui

import "github.com/charmbracelet/lipgloss"

func Canvas(width, height int, content string) string {
	renderedCanvas := CanvasStyles(width, height).Render(content)

	return lipgloss.JoinVertical(lipgloss.Left, "canvas", renderedCanvas)
}

func CanvasStyles(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(GreenColor()).
		MarginRight(2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}
