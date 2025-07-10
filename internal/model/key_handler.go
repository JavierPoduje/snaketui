package model

import (
	"snaketui/internal/game"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	// Game actions
	case key.Matches(msg, m.keys.Quit):
		return m.handleQuit()
	case key.Matches(msg, m.keys.Pause):
		return m.handlePause()
	case key.Matches(msg, m.keys.Restart):
		return m.handleRestart()
	case key.Matches(msg, m.keys.Help):
		return m.handleHelp()

	// directions
	case key.Matches(msg, m.keys.Up):
		return m.handleKeyUp()
	case key.Matches(msg, m.keys.Right):
		return m.handleKeyRight()
	case key.Matches(msg, m.keys.Down):
		return m.handleKeyDown()
	case key.Matches(msg, m.keys.Left):
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

func (m *Model) handleHelp() (Model, tea.Cmd) {
	m.help.ShowAll = !m.help.ShowAll
	return *m, nil
}
