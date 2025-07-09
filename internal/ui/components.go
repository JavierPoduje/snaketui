package ui

import "github.com/charmbracelet/lipgloss"

func Apple(char string) string {
	return lipgloss.NewStyle().
		Foreground(RedColor()).
		Render(char)
}

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

func Snake(char string, isHead bool) string {
	snakeColor := func() lipgloss.TerminalColor {
		if isHead {
			return GreenColor()
		}
		return PrimaryTextColor()
	}()

	return lipgloss.NewStyle().
		Foreground(snakeColor).
		Render(char)
}
