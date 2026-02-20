package internal

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// if detected a key press
	case tea.KeyMsg:
		key := msg.String()

		switch m.state {
		case listView:
			switch key {
			// These keys should exit the program.
			case "ctrl+c", "q":
				return m, tea.Quit

			// handle up and down selections
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}

			// The "enter" key and the spacebar (a literal space) toggle
			// the selected state for the item that the cursor is pointing at.
			case "enter", " ":
				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}

			// delete marked tasks
			case "backspace", "d":

			// editing an existing task switches the body view
			case "e":

			// creating a new task switches to the title view
			case "n":
				m.state = titleView
			}

			// case titleView:
			// case bodyView:
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
