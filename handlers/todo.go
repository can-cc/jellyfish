package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"jellyfish/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTodos(db))
	}
}

// PutTask endpoint
func PostTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		todo := new(models.Todo)

		c.Bind(&todo)

		id, err := models.PostTodo(db, todo)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteTodo(db, id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
