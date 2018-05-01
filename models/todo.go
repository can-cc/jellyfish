package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Todo struct {
	ID        int        `json:"id"`
	Content   string     `json:"content"`
	Detail    string     `json:"detail"`
	Deadline  *time.Time `json:"deadline"`
	Status    string     `json:"status"`
	CreaterId string     `json:"createrId"`
	CreatedAt *time.Time `json:"createdAt"`
}

// TaskCollection is collection of Tasks
type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func GetTodosFromDB(db *sql.DB, userId string) TodoCollection {
	sql := "SELECT id, content, detail, deadline, status, created_at FROM todos where creater_id = ?"
	rows, err := db.Query(sql, userId)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	todoCollection := TodoCollection{Todos: make([]Todo, 0)}

	for rows.Next() {
		todo := Todo{}
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Deadline, &todo.Status, &todo.CreatedAt)

		if err2 != nil {
			panic(err2)
		}
		todoCollection.Todos = append(todoCollection.Todos, todo)
	}
	return todoCollection
}

// PutTask into DB
func PostTodo(db *sql.DB, todo *Todo) (int64, error) {
	sql := "INSERT INTO todos(content, detail, creater_id, deadline, status, created_at) VALUES(?, ?, ?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.CreaterId, todo.Deadline, todo.Status, time.Now().Unix())

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
