package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Todo struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	Detail    string `json:"detail"`
	Deadline  int64  `json:"deadline"`
	Done      bool   `json:"done"`
	Status    string `json:"status"`
	CreaterId string `json:"createrId"`
	CreatedAt int64  `json:"createdAt"`
}

// TaskCollection is collection of Tasks
type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func GetTodosFromDB(db *sql.DB, userId string) TodoCollection {
	sql := "SELECT id, content, detail, deadline, status, done, created_at FROM todos where creater_id = ?"
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
		var deadline time.Time
		var createdAt time.Time
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &deadline, &todo.Status, &todo.Done, &createdAt)
		todo.Deadline = deadline.UnixNano()
		todo.CreatedAt = createdAt.UnixNano()

		if err2 != nil {
			panic(err2)
		}
		todoCollection.Todos = append(todoCollection.Todos, todo)
	}
	return todoCollection
}

func GetTodo(db *sql.DB, todoId string) Todo {
	sql := "SELECT id, content, detail, deadline, status, creater_id, created_at FROM todos where id = ?"
	row := db.QueryRow(sql, todoId)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Deadline, &todo.Status, &todo.CreaterId, &todo.CreatedAt)
	if err != nil {
		panic(err)
	}
	return todo
}

func UpdateTodo(db *sql.DB, todo *Todo) (int64, error) {
	sql := "UPDATE todos set content = ?, detail = ?, done = ?, deadline = ?, status = ?, updated_at = ? where id = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.Done, todo.Deadline, todo.Status, time.Now().UnixNano(), todo.ID)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

// PutTask into DB
func CreateTodo(db *sql.DB, todo *Todo) (int64, error) {
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
	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.CreaterId, todo.Deadline, todo.Status, time.Now().UnixNano())

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteTodo(db *sql.DB, id int, userId string) (int64, error) {
	sql := "DELETE FROM todos WHERE id = ? and creater_id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id, userId)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
