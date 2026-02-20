package internal

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state    uint
	store    *Store
	tasks    []Tasks          // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func NewModel(store *Store) model {
	tasks, err := store.GetTodoList()
	if err != nil {
		log.Fatalf("Unable to get todo list: %v", err)
	}

	return model{
		// Our to-do list is a grocery list
		state: listView,
		store: store,
		tasks: tasks,
		// A map which indicates which tasks are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `tasks` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
