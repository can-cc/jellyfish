package repository

import (
	"github.com/fwchen/jellyfish/database"
	"github.com/fwchen/jellyfish/models"
)

func GetUserTodos(userID string, statusTag string) models.TodoCollection {
	db := database.GetDB()

	var additionSQL string = ""
	if statusTag == "Doing" {
		additionSQL = " AND done = 0"
	} else if statusTag == "Done" {
		additionSQL = " AND done = 1"
	}

	sql := `SELECT id, TRIM(content), TRIM(detail), TRIM(type), deadline, TRIM(status), done, created_at FROM todos where creater_id = $1` + additionSQL
	rows, err := db.Query(sql, userID)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	todoCollection := models.TodoCollection{Items: make([]models.Todo, 0)}

	for rows.Next() {
		todo := models.Todo{}
		err2 := rows.Scan(&todo.ID, &todo.Content, &todo.Detail, &todo.Type, &todo.Deadline, &todo.Status, &todo.Done, &todo.CreatedAt)

		if err2 != nil {
			panic(err2)
		}
		todoCollection.Items = append(todoCollection.Items, todo)
	}
	return todoCollection
}

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

func UpdateTodo(todo *models.Todo) {
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

func CreateTodo(todo *models.Todo) (string, error) {
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

func DeleteTodo(todoID string, userID string) error {
	db := database.GetDB()

	_, err := db.Exec(`DELETE FROM todos WHERE id = $1 and creater_id = $2`, todoID, userID)

	return err
}
