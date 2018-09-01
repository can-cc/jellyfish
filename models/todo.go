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
	Type      string `json:"type"`
	CreaterId string `json:"createrId"`
	CreatedAt int64  `json:"createdAt"`
}

type CycleTodoStatus struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	date      string `json:"date"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// TaskCollection is collection of Tasks
type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func GetTodosFromDB(db *sql.DB, userId string) TodoCollection {
	sql := "SELECT id, content, detail, type, deadline, status, done, created_at FROM todos where creater_id = ?"
	rows, err := db.Query(sql, userId)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	todoCollection := TodoCollection{Todos: make([]Todo, 0)}

	for rows.Next() {
		todo := Todo{}
		var deadline time.Time
		var createdAt time.Time
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Type, &deadline, &todo.Status, &todo.Done, &createdAt)
		todo.Deadline = deadline.UnixNano() / int64(time.Millisecond)
		todo.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)

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

	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.Done, todo.Deadline, todo.Status, time.Now().UnixNano()/int64(time.Millisecond), todo.ID)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func CheckCycleTodoStatusExist(db *sql.DB, todoId string) bool {
	sql := "SELECT id FROM cycle_todo_status where todo_id = ? and date = ?"
	row := db.QueryRow(sql, todoId)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	return true
}

func CreateCycleTodoStatus(db *sql.DB, todoId string, done bool) (int64, error) {
	sql := "INSERT INTO cycle_todo_status(todo_id, status, date, created_at, updated_at) VALUES(?, ?, ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	t := time.Now()
	dateString := t.Format("2006-01-02")

	var status string
	if done {
		status = "DONE"
	} else {
		status = "UNDONE"
	}

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(todoId, status, dateString, time.Now().UnixNano()/int64(time.Millisecond), time.Now().UnixNano()/int64(time.Millisecond))

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func UpdateCycleTodoStatus(db *sql.DB, todoId string, done bool) (int64, error) {
	sql := "UPDATE cycle_todo_status set status = ?, updated_at = ? where todo_id = ? and date = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	var status string
	if done {
		status = "DONE"
	} else {
		status = "UNDONE"
	}

	t := time.Now()
	dateString := t.Format("2006-01-02")

	result, err2 := stmt.Exec(status, time.Now().UnixNano()/int64(time.Millisecond), todoId, dateString)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func MarkCycleTodo(db *sql.DB, todoId string, done bool) (int64, error) {
	cycleTodoExist := CheckCycleTodoStatusExist(db, todoId)
	if cycleTodoExist {
		return UpdateCycleTodoStatus(db, todoId, done)
	} else {
		return CreateCycleTodoStatus(db, todoId, done)
	}
}

// PutTask into DB
func CreateTodo(db *sql.DB, todo *Todo) (int64, error) {
	sql := "INSERT INTO todos(content, detail, type, creater_id, deadline, status, created_at) VALUES(?, ?, ?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.Type, todo.CreaterId, todo.Deadline, todo.Status, time.Now().UnixNano()/int64(time.Millisecond))

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
