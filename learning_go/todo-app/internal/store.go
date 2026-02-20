package internal

import (
	"database/sql"
	"time"

	// We use the underscore import because we only need the side effects (driver registration)
	_ "github.com/mattn/go-sqlite3"
)

type Tasks struct {
	ID        int64
	Title     string
	Body      string
	StartTime time.Time
	EndTime   time.Time
}

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error

	s.conn, err = sql.Open("sqlite3", "./notes.db")
	if err != nil {
		return err
	}

	createTableStmt := `
	CREATE TABLE IF NOT EXISTS todo (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		body TEXT,
		startTime TEXT NOT NULL, 
		endTime TEXT
	);`

	if _, err = s.conn.Exec(createTableStmt); err != nil {
		return err
	}

	return nil
}

func (s *Store) GetTodoList() ([]Tasks, error) {
	query := `SELECT * FROM todo`
	rows, err := s.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []Tasks{}

	for rows.Next() {
		var task Tasks
		rows.Scan(&task.ID, &task.Title, &task.Body, &task.StartTime, &task.EndTime)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// saves a new task or creates a new task
func (s *Store) SaveTask(task Tasks) error {
	// if the task is new, auto increment id
	if task.ID == 0 {
		insertQuery := `INSERT INTO todo (title, body, startTime, endTime) VALUES (?, ?, ?, ?, ?)`
		_, err := s.conn.Exec(insertQuery, task.Title, task.Body, task.StartTime, task.EndTime)
		if err != nil {
			return err
		}
	}

	// task id must already exists, therefore modify by passing the task id
	updateQuery := `UPDATE todo SET title = ?, body = ?, startTime = ?, endTime = ? WHERE id = ?`
	_, err := s.conn.Exec(updateQuery, task.Title, task.Body, task.StartTime, task.EndTime, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteTask(tasks []Tasks) error {
	updateQuery := `DELETE FROM todo WHERE id = ?`

	// handle multiple deletions (probably can be optimized for a single query)
	for i := range tasks {
		_, err := s.conn.Exec(updateQuery, tasks[i].ID)
		if err != nil {
			return err
		}
	}
	return nil
}
