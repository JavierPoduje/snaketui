package model

import (
	"snaketui/internal/game"

	tea "github.com/charmbracelet/bubbletea"
	"slices"
)

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	// Game actions
	case slices.Contains([]string{"q", "esc", "ctrl+c"}, msg.String()):
		return m.handleQuit()
	case slices.Contains([]string{"p"}, msg.String()):
		return m.handlePause()
	case slices.Contains([]string{"r"}, msg.String()):
		return m.handleRestart()

	// directions
	case slices.Contains([]string{"up", "k"}, msg.String()):
		return m.handleKeyUp()
	case slices.Contains([]string{"right", "l"}, msg.String()):
		return m.handleKeyRight()
	case slices.Contains([]string{"down", "j"}, msg.String()):
		return m.handleKeyDown()
	case slices.Contains([]string{"left", "h"}, msg.String()):
		return m.handleKeyLeft()
	default:
		return *m, nil
	}
}

func (m *Model) handlePause() (Model, tea.Cmd) {
	switch m.game.State {
	case game.Running:
		m.game.State = game.Paused
	case game.Paused:
		m.game.State = game.Running
		return *m, m.tick(m.game.Snake.Speed)
	default:
		// don't do nothing here
	}
	return *m, nil
}

func (m *Model) handleRestart() (Model, tea.Cmd) {
	switch m.game.State {
	case game.Paused:
		m.game.State = game.Running
		return *m, m.tick(m.game.Snake.Speed)
	case game.GameOver:
		m.game.State = game.Running
		return RestartModel(m.terminalWidth, m.terminalHeight), m.tick(m.game.Snake.Speed)
	default:
		panic("Unreachable state during restart")
	}
}

func (m *Model) handleKeyUp() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Up) {
		m.nextSnakeMove = game.Up
	}
	return *m, nil
}

func (m *Model) handleKeyRight() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Right) {
		m.nextSnakeMove = game.Right
	}
	return *m, nil
}

func (m *Model) handleKeyDown() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Down) {
		m.nextSnakeMove = game.Down
	}
	return *m, nil
}

func (m *Model) handleKeyLeft() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Left) {
		m.nextSnakeMove = game.Left
	}
	return *m, nil
}

func (m *Model) handleQuit() (Model, tea.Cmd) {
	return *m, tea.Quit
}
