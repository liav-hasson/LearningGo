package main

import (
	"log"
	"os"

	"todo-app/internal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := &internal.Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("Unable to init store: %v", err)
	}

	m := internal.NewModel(store)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Unable to run CLI app: %v", err)
		os.Exit(1)
	}
}
