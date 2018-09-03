package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	// "time"

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

func GetTodoCycles(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.QueryParam("userId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		jwtUserId := claims["id"].(string)

		if userId != jwtUserId {
			return c.JSON(http.StatusUnauthorized, "")
		}
		tcs := models.GetTodosFromDB(db, userId).Todos
		return c.JSON(http.StatusOK, todos)
	}
}

func MarkCycleTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := new(struct {
			TodoId int  `json:"todoId"`
			Done   bool `json:"done"`
		})
		c.Bind(&request)

		id, err := models.MarkCycleTodo(db, strconv.Itoa(request.TodoId), request.Done)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, id)
	}
}

func PutTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		// TODO get todo id in url
		todo := new(models.Todo)
		c.Bind(&todo)

		_, err := models.UpdateTodo(db, todo)
		if err == nil {
			return c.NoContent(http.StatusCreated)
		} else {
			return err
		}

	}
}

// PutTask endpoint
func PostTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		todo := new(models.Todo)

		c.Bind(&todo)

		if todo.Type != "NORMAL" && todo.Type != "HABIT" {
			return c.NoContent(http.StatusBadRequest)
		}

		todo.CreaterId = userId

		id, err := models.CreateTodo(db, todo)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"id": id,
			})
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		id, _ := strconv.Atoi(c.Param("id"))
		// Use our new model to delete a task
		_, err := models.DeleteTodo(db, id, userId)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{})
			// Handle errors
		} else {
			return err
		}
	}
}
