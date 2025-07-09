package model

import (
	"snaketui/internal/game"
	"snaketui/internal/ui"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	DefaultTerminalWidth  = 40
	DefaultTerminalHeight = 24
	CanvasWidth           = 20
	CanvasHeight          = 20
)

const DefaultSnakeDir = game.Right
const DefaultSnakeSpeed = float64(3)

const (
	SnakeChar   = "S"
	NeutralChar = "."
)

type TickMsg time.Time

type Model struct {
	game           *game.Game
	nextSnakeMove  game.Direction
	terminalHeight int
	terminalWidth  int
}

func NewModel() Model {
	return Model{
		game:           game.NewGame(CanvasWidth, CanvasHeight),
		terminalHeight: DefaultTerminalHeight,
		nextSnakeMove:  DefaultSnakeDir,
		terminalWidth:  DefaultTerminalWidth,
	}
}

func (m Model) Init() tea.Cmd {
	return m.tick(DefaultSnakeSpeed)
}

func (m Model) tick(snakeSpeed float64) tea.Cmd {
	interval := time.Second / time.Duration(snakeSpeed)
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.HandleWindowResize(msg)
	case tea.KeyMsg:
		return m.HandleKeyPressed(msg)
	case TickMsg:
		return m.HandleTick()
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

func (m Model) BuildNextCanvasContent() string {
	strCanvas := strings.Builder{}

	width := m.game.Canvas.Width
	height := m.game.Canvas.Height
	snake := m.game.Snake

	for Y := range width {
		for X := range height {
			coord := game.Coord{X: X, Y: Y}

			if snake.Contains(coord) {
				snakeComponent := ui.Snake(SnakeChar, snake.IsHead(coord))
				strCanvas.WriteString(snakeComponent)
			} else {
				strCanvas.WriteString(NeutralChar)
			}
		}
	}

	return strCanvas.String()
}

func (m *Model) HandleTick() (Model, tea.Cmd) {
	m.game.Tick(m.nextSnakeMove)

	return *m, m.tick(m.game.Snake.Speed)
}
