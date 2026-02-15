package internal

import "os"

// go's json package can only parse capital letter feilds
type Tasks struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Body     string `json:"body"`
	Priority string `json:"priority"`

	// TODO: implenent task deadline
	// start_date string
	// end_date string
}

// Import and parse tasks.json
func getExistingTaskList() (error, *[]Tasks) {
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		return err, nil
	}

	// implement file parsing
}

// create new data structure
func createNewTaskList() *[]Tasks {
	return &[]Tasks{}
}
