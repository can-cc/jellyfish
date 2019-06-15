package todorepository

import (
	"github.com/fwchen/jellyfish/database"

	. "github.com/fwchen/jellyfish/models"
)

// GetUserTodos :
func GetUserTodos(userID string) TodoCollection {
	db := database.GetDB()

	sql := `SELECT id, TRIM(content), TRIM(detail), TRIM(type), deadline, TRIM(status), done, created_at FROM todos where creater_id = $1`
	rows, err := db.Query(sql, userID)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	todoCollection := TodoCollection{Items: make([]Todo, 0)}

	for rows.Next() {
		todo := Todo{}
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Type, &todo.Deadline, &todo.Status, &todo.Done, &todo.CreatedAt)

		if err2 != nil {
			panic(err2)
		}
		todoCollection.Items = append(todoCollection.Items, todo)
	}
	return todoCollection
}

// GetTodo :
func GetTodo(todoID string) Todo {
	db := database.GetDB()
	sql := "SELECT id, content, detail, deadline, status, creater_id, created_at FROM todos where id = ?"
	row := db.QueryRow(sql, todoID)

	var todo Todo
	err := row.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Deadline, &todo.Status, &todo.CreatorID, &todo.CreatedAt)
	if err != nil {
		panic(err)
	}
	return todo
}

// UpdateTodo :
func UpdateTodo(todo *Todo) {
	db := database.GetDB()
	sql := `UPDATE todos set content = $1, detail = $2, done = $3, deadline = $4, updated_at = now() where id = $5`

	var doneValue int
	if todo.Done {
		doneValue = 1
	} else {
		doneValue = 0
	}

	_, err := db.Exec(sql, todo.Content, todo.Detail, doneValue, todo.Deadline, todo.ID)

	if err != nil {
		panic(err)
	}
}

// CreateTodo :
func CreateTodo(todo *Todo) (string, error) {
	db := database.GetDB()

	var id string
	err := db.QueryRow(`INSERT INTO todos(content, detail, type, creater_id, deadline, status, created_at) VALUES($1, $2, $3, $4, $5, $6, now()) RETURNING id`,
		todo.Content,
		todo.Detail,
		todo.Type,
		todo.CreatorID,
		todo.Deadline,
		todo.Status,
	).Scan(&id)

	return id, err
}

// DeleteTodo : from DB
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

	result, err2 := stmt.Exec(id, userId)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
