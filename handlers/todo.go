package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fwchen/jellyfish/models"
	todorepository "github.com/fwchen/jellyfish/repository/todo"

	"github.com/labstack/echo"
)

// GetUserTodos
func GetUserTodos() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		jwtUserID := claims.ID

		todos := todorepository.GetUserTodos(jwtUserID, "All").Items
		return c.JSON(http.StatusOK, todos)
	}
}

// GetUserDoingTodos :
func GetUserDoingTodos() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		jwtUserID := claims.ID

		todos := todorepository.GetUserTodos(jwtUserID, "Doing").Items
		return c.JSON(http.StatusOK, todos)
	}
}

// GetUserDoneTodos :
func GetUserDoneTodos() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		jwtUserID := claims.ID

		todos := todorepository.GetUserTodos(jwtUserID, "Done").Items
		return c.JSON(http.StatusOK, todos)
	}
}

// UpdateTodo :
func UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		todo := new(models.Todo)
		c.Bind(&todo)

		todorepository.UpdateTodo(todo)
		return c.NoContent(http.StatusCreated)
	}
}

// CreateTodo :
func CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		userID := claims.ID

		todo := new(models.Todo)

		c.Bind(&todo)

		todo.CreatorID = userID
		if todo.Status == "" {
			todo.Status = "VALID"
		}
		if todo.Type == "" {
			todo.Type = "NORMAL"
		}

		id, err := todorepository.CreateTodo(todo)

		if err == nil {
			return c.JSON(http.StatusCreated, map[string]string{
				"id": id,
			})
		}
		panic(err)
	}
}

// DeleteTodo :
func DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtAppClaims)
		userID := claims.ID

		todoID := c.Param("id")
		error := todorepository.DeleteTodo(todoID, userID)

		if error == nil {
			return c.JSON(http.StatusOK, map[string]string{})
		}
		return error
	}
}
