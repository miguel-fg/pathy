package main

import (
	"log"

	"os"
	"path/filepath"
	"pathy/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var startDir string

	if len(os.Args) > 1 {
		startDir = os.Args[1]
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		startDir = wd
	}

	absDir, err := filepath.Abs(startDir)
	if err != nil {
		log.Fatal(err)
	}

	prog := tea.NewProgram(ui.NewModel(absDir), tea.WithAltScreen())
	if _, err := prog.Run(); err != nil {
		log.Fatal(err)
	}
}
