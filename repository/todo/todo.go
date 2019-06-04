package todorepository

import (
	"github.com/fwchen/jellyfish/database"
	"time"

	"github.com/fwchen/jellyfish/models"

	_ "github.com/mattn/go-sqlite3"
)

// GetUserTodos :
func GetUserTodos(userID string) models.TodoCollection {
	db := database.GetDB()
	sql := "SELECT id, content, detail, type, deadline, status, done, created_at FROM todos where creater_id = ?"
	rows, err := db.Query(sql, userID)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	todoCollection := models.TodoCollection{Items: make([]models.Todo, 0)}

	for rows.Next() {
		todo := models.Todo{}
		var deadline time.Time
		var createdAt time.Time
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Type, &deadline, &todo.Status, &todo.Done, &createdAt)
		todo.Deadline = deadline.UnixNano() / int64(time.Millisecond)
		todo.CreatedAt = createdAt.UnixNano() / int64(time.Millisecond)

		if err2 != nil {
			panic(err2)
		}
		todoCollection.Items = append(todoCollection.Items, todo)
	}
	return todoCollection
}

// GetTodo :
func GetTodo(todoID string) models.Todo {
	db := database.GetDB()
	sql := "SELECT id, content, detail, deadline, status, creater_id, created_at FROM todos where id = ?"
	row := db.QueryRow(sql, todoID)

	var todo models.Todo
	err := row.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Deadline, &todo.Status, &todo.CreatorID, &todo.CreatedAt)
	if err != nil {
		panic(err)
	}
	return todo
}

// UpdateTodo :
func UpdateTodo(todo *models.Todo) (int64, error) {
	db := database.GetDB()
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

// CheckCycleTodoStatusExist :
func CheckCycleTodoStatusExist(todoId string) bool {
	db := database.GetDB()
	t := time.Now()
	dateString := t.Format("2006-01-02")
	sql := "SELECT id FROM cycle_todo_status where todo_id = ? and date = ?"
	row := db.QueryRow(sql, todoId, dateString)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	return true
}

// CreateCycleTodoStatus :
func CreateCycleTodoStatus(todoId string, done bool) (int64, error) {
	db := database.GetDB()
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

// UpdateCycleTodoStatus :
func UpdateCycleTodoStatus(todoId string, done bool) (int64, error) {
	db := database.GetDB()
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

// CreateTodo:
func CreateTodo(todo *models.Todo) (int64, error) {
	db := database.GetDB()
	sql := "INSERT INTO todos(content, detail, type, creater_id, deadline, status, created_at) VALUES(?, ?, ?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(todo.Content, todo.Detail, todo.Type, todo.CreatorID, todo.Deadline, todo.Status, time.Now().UnixNano()/int64(time.Millisecond))

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteTodo(id int, userId string) (int64, error) {
	db := database.GetDB()
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


