package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Todo struct {
	ID        int     `json:"id"`
	Content   string  `json:"content"`
	Detail    string  `json:"detail"`
	Deadline  string  `json:"deadline"`
	Status    string  `json:"status"`
	CreatedAt float64 `json:"created_at"`
}

// TaskCollection is collection of Tasks
type TodoCollection struct {
	Todos []Todo `json:"items"`
}

// GetTasks from the DB
func GetTodos(db *sql.DB) TodoCollection {
	sql := "SELECT * FROM todos"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TodoCollection{}
	for rows.Next() {
		todo := Todo{}
		err2 := rows.Scan(&todo.ID, &todo.Content)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Todos = append(result.Todos, todo)
	}
	return result
}

// PutTask into DB
func PostTodo(db *sql.DB, todo *Todo) (int64, error) {
	sql := "INSERT INTO todos(content, detail, deadline, status, created_at) VALUES(?, ?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.Deadline, todo.Status, time.Now().Unix())
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteTodo(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM todos WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
