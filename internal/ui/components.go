package ui

import (
	"snaketui/internal/game"

	"github.com/charmbracelet/lipgloss"
)

func Apple(char string) string {
	return lipgloss.NewStyle().
		Foreground(RedColor()).
		Render(char)
}

func Canvas(width, height int, state game.GameState, content string) string {
	renderedCanvas := CanvasStyles(width, height, state).Render(content)

	return lipgloss.JoinVertical(lipgloss.Left, CanvasLabel(state), renderedCanvas)
}

func CanvasLabel(state game.GameState) string {
	var label string
	switch state {
	case game.Running:
		label = "Running"
	case game.GameOver:
		label = "Game Over"
	case game.Paused:
		label = "Paused"
	default:
		panic("Unknown game state")
	}

	return CanvasLabelStyles(state).Render(label)
}

func CanvasLabelStyles(state game.GameState) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ColorByStateStyles(state)).
		Height(1).
		MarginRight(2)
}

func ColorByStateStyles(state game.GameState) lipgloss.Color {
	var color lipgloss.Color
	switch state {
	case game.Running:
		color = GreenColor()
	case game.GameOver:
		color = RedColor()
	case game.Paused:
		color = OrangeColor()
	default:
		panic("Unknown game state")
	}
	return color
}

func CanvasStyles(width, height int, state game.GameState) lipgloss.Style {
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
