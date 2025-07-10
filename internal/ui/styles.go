package ui

import "github.com/charmbracelet/lipgloss"

func StatHeaderStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(SecondaryTextColor()).
		Align(lipgloss.Left).
		PaddingRight(3).
		Height(1)
}

func StatValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PrimaryTextColor()).
		Align(lipgloss.Right).
		Height(1)
}

func StatsStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(PrimaryTextColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(GrayColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		MarginTop(1).
		Height(6)
}

func TitleStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(SecondaryTextColor()).
		Align(lipgloss.Center).
		Height(1).
		MarginBottom(1)
}
