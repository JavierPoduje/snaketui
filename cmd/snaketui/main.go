package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"snaketui/internal/model"
)

func main() {
	p := tea.NewProgram(model.NewModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}
