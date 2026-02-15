// SIMPLE TO-DO CLI APP

// when running:
// welcome message
// your tasks (list)
// available actions

// available actions commands:
// add - create a new task
// edit - edit an existing task
// delete - remove a task
// list - lists all tasks

// task structure
// title, body, data-added, date-deadline

package main

import "todo-app/internal"

func main() {
	internal.InitTasks()
	internal.StartRepl()
}
