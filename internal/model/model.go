package model

import (
	"snaketui/internal/ui"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	DefaultTerminalWidth  = 40
	DefaultTerminalHeight = 24
	CanvasWidth           = 20
	CanvasHeight          = 20
)

const (
	NeutralChar = "."
)

type Model struct {
	terminalHeight int
	terminalWidth  int
}

func NewModel() Model {
	return Model{
		terminalHeight: DefaultTerminalHeight,
		terminalWidth:  DefaultTerminalWidth,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.HandleWindowResize(msg)
	case tea.KeyMsg:
		return m.HandleKeyPressed(msg)
	default:
		panic("unhandled message type")
	}

	return m, nil
}

func (m Model) View() string {
	canvasContent := m.BuildNextCanvasContent()
	canvas := ui.Canvas(CanvasWidth, CanvasHeight, canvasContent)

	return lipgloss.Place(
		m.terminalWidth, m.terminalHeight,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, canvas),
	)
}

func (m *Model) HandleWindowResize(msg tea.WindowSizeMsg) {
	m.terminalWidth, m.terminalHeight = msg.Width, msg.Height
}

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	// Game actions
	case msg.String() == "q" || msg.String() == "esc" || msg.String() == "ctrl+c":
		return *m, tea.Quit
	default:
		return *m, nil
	}
}

func (m Model) BuildNextCanvasContent() string {
	strCanvas := strings.Builder{}

	for _ = range CanvasWidth {
		for _ = range CanvasHeight {
			strCanvas.WriteString(NeutralChar)
		}
	}

	return strCanvas.String()
}
