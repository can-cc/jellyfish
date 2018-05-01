package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"jellyfish/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.QueryParam("userId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		if userId != jwtUserId {
			return c.JSON(http.StatusUnauthorized, "")
		}
		todos := models.GetTodosFromDB(db, userId).Todos
		return c.JSON(http.StatusOK, todos)
	}
}

// PutTask endpoint
func PostTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		todo := new(models.Todo)

		// deadlinePath := new(struct {
		// 	Deadline int64 `json:"deadline"`
		// })

		// fmt.Println(deadlinePath.Deadline)
		// var deadline time.Time
		// if deadlinePath.Deadline != 0 {
		// 	deadline = time.Unix(deadlinePath.Deadline, 0)
		// }

		c.Bind(&todo)

		todo.CreaterId = userId

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
