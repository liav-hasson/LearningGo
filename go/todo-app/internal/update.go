package internal

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// a random date, go only cares about the format
const timeFormat = "Jan 02, 2006 15:04"

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)

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
				if len(m.selected) > 0 {
					var toDelete []Tasks
					for idx := range m.selected {
						toDelete = append(toDelete, m.tasks[idx])
					}

					if err := m.store.DeleteTask(toDelete); err != nil {
						return m, tea.Quit
					}

					var err error
					m.tasks, err = m.store.GetTodoList()
					if err != nil {
						return m, tea.Quit
					}

					// reset selections and clamp cursor
					m.selected = make(map[int]struct{})
					if m.cursor >= len(m.tasks) && len(m.tasks) > 0 {
						m.cursor = len(m.tasks) - 1
					} else if len(m.tasks) == 0 {
						m.cursor = 0
					}
				}

			// editing an existing task switches the body view
			case "e":
				// handle edit attempt when no tasks exist
				if len(m.tasks) == 0 {
					break
				}
				m.currTask = m.tasks[m.cursor]
				m.textarea.SetValue(m.currTask.Body)
				m.textarea.Focus()
				m.textarea.CursorEnd()
				m.state = bodyView

			// creating a new task switches to the title view
			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.textinput.CursorEnd()
				m.currTask = Tasks{}
				m.state = titleView
			}
		
		case titleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.currTask.Title = title

					m.textarea.SetValue("")
					m.textarea.Focus()
					m.textarea.CursorEnd()

					m.state = bodyView
				}

			case "esc":
				m.state = listView
			}

		case bodyView:
			switch key {
			case "ctrl+s":
				body := m.textarea.Value()
				m.currTask.Body = body

				// set start time for new tasks
				if m.currTask.ID == 0 {
					m.currTask.StartTime = time.Now().Format(timeFormat)
				}

				// save task to db
				var err error
				if err = m.store.SaveTask(m.currTask); err != nil {
					// TODO handle errors here
					return m, tea.Quit
				}

				m.tasks, err = m.store.GetTodoList()
				if err != nil {
					// TODO handle errors here
					return m, tea.Quit
				}

				// reset the current task
				m.currTask = Tasks{}

				// go back to the list view
				m.state = listView

			case "esc":
				// exit to list view
				m.state = listView
			}	
		}
	}

	return m, tea.Batch(cmds...)
}
