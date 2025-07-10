package model

import (
	"snaketui/internal/db"
	"snaketui/internal/game"
	"snaketui/internal/ui"
	"strconv"
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
	AppleChar   = "A"
	NeutralChar = "."
	SnakeChar   = "S"
)

type TickMsg time.Time

type Model struct {
	db             db.DB
	game           *game.Game
	nextSnakeMove  game.Direction
	scores         []int
	terminalHeight int
	terminalWidth  int
}

func NewModel() Model {
	db := db.NewDB()
	return Model{
		db:             db,
		game:           game.NewGame(CanvasWidth, CanvasHeight),
		nextSnakeMove:  DefaultSnakeDir,
		scores:         db.GetScores(),
		terminalHeight: DefaultTerminalHeight,
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
	stats := m.BuildNextStatsContent()

	// components
	canvas := ui.Canvas(CanvasWidth, CanvasHeight, m.game.State, canvasContent)
	statsCard := ui.StatsCard(stats)
	historicScoresCard := ui.HistoricScoresCard(m.scores)

	infoCards := lipgloss.JoinVertical(lipgloss.Center, statsCard, historicScoresCard)
	contentSection := lipgloss.JoinHorizontal(lipgloss.Top, canvas, infoCards)
	content := lipgloss.JoinVertical(lipgloss.Right, contentSection)

	return lipgloss.Place(
		m.terminalWidth, m.terminalHeight,
		lipgloss.Center, lipgloss.Center,
		content,
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
	apple := m.game.Apple

	for Y := range width {
		for X := range height {
			coord := game.Coord{X: X, Y: Y}

			if apple.Equals(coord) {
				strCanvas.WriteString(ui.Apple(AppleChar))
			} else if snake.Contains(coord) {
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
	if m.game.State == game.Paused {
		return *m, nil
	}

	m.game.Tick(m.nextSnakeMove)

	if m.game.State == game.GameOver {
		if m.game.Stats.ScoreAsInt() > 0 {
			m.db.SaveScore(m.game.Stats.ScoreAsInt())
			m.scores = m.db.GetScores()
		}
		return *m, nil
	}

	return *m, m.tick(m.game.Snake.Speed)
}

func RestartModel(width, height int) Model {
	m := NewModel()
	m.terminalWidth = width
	m.terminalHeight = height
	return m
}

func (model Model) BuildNextStatsContent() [][]string {
	return [][]string{
		{"Eaten apples:", strconv.Itoa(model.game.Stats.EatenApples)},
		{"Score:", model.game.Stats.RoundedScoreAsString()},
	}
}
