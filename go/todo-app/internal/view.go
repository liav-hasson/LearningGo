package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// colors and styles
var (
	appNameStyle    = lipgloss.NewStyle().Background(lipgloss.Color("#453f9c")).Padding(0, 1)
	faintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#cacdff")).Faint(true)
	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#e0f79f")).MarginRight(1)
)

func (m model) View() string {
	// The header "Interative TODO List CLI Application\n\n"
	s := appNameStyle.Render("MY CLI TODO LIST") + "\n\n"

	if m.state == titleView {
		s += "Task title: \n\n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("⏎ - save title, esc - discard")
	}

	if m.state == bodyView {
		s += "Task description: \n\n"
		s += m.textarea.View() + "\n\n"
		s += faintStyle.Render("ctrl+s - save task, esc - discard")
	}

	if m.state == listView {
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

			// Render the task metadata
			metadata := fmt.Sprintf("[%s] Task %d: %s | Started: %s", checked, task.ID, task.Title, task.StartTime)
			s += fmt.Sprintf("%s %s\n", enumeratorStyle.Render(cursor), faintStyle.Render(metadata))

			// Render the task body (indent all lines with a tab)
			indentedBody := "\t" + strings.ReplaceAll(task.Body, "\n", "\n\t")
			s += indentedBody + "\n\n"
		}

		s += faintStyle.Render(`
Available actions: 
	⇅ - move cursor
	⏎ - mark a selected task
	n - create a new task
	e - edit an existing task
	d - delete all marked tasks
	q - quit
		`)
	}

	// Send the UI for rendering
	return s
}
