package internal

import "fmt"

func (m model) View() string {
	// The header
	s := "Interative TODO List CLI Application\n\n"

	// Iterate over our choices
	for i, task := range m.tasks {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] Task %d: %s | Starts: %s | Ends: %s\n \t- %s\n\n",
			cursor, checked, task.ID, task.Title, task.StartTime.Format("Jan 02, 15:04"),
			task.EndTime.Format("Jan 02, 15:04"), task.Body,
		)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
