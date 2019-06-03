package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/fwchen/jellyfish/models"
	"github.com/fwchen/jellyfish/repository/todo"

	"github.com/labstack/echo"
)

// GetUserTodos :
func GetUserTodos(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.QueryParam("userId")

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		jwtUserID := claims.ID

		if userID != jwtUserID {
			return c.JSON(http.StatusUnauthorized, "")
		}
		todos := todo_repository.GetUserTodos(userID).Items
		return c.JSON(http.StatusOK, todos)
	}
}

// UpdateTodo :
func UpdateTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		todo := new(models.Todo)
		c.Bind(&todo)

		_, err := todo_repository.UpdateTodo(todo)
		if err == nil {
			return c.NoContent(http.StatusCreated)
		}
		return err
	}
}

// CreateTodo :
func CreateTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		userID := claims.ID

		todo := new(models.Todo)

		c.Bind(&todo)

		todo.CreatorID = userID

		id, err := todo_repository.CreateTodo(todo)

		if err == nil {
			return c.JSON(http.StatusCreated, map[string]string{
				"id": strconv.FormatInt(id, 10),
			})
		}
		return err
	}
}

// DeleteTodo :
func DeleteTodo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		userID := claims.ID

		id, _ := strconv.Atoi(c.Param("id"))
		_, err := todo_repository.DeleteTodo(id, userID)
		if err == nil {
			return c.JSON(http.StatusOK, map[string]string{})
		} else {
			return err
		}
	}
}
