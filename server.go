package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users []User

func createUser(user User) {
	users = append(users, user)
}

func getUser(id int) (User, error) {
	for _, user := range users {
		fmt.Println(user)
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func updateUser(id int, newUser User) error {
	for i, user := range users {
		if user.ID == id {
			users[i] = newUser
			return nil
		}
	}
	return errors.New("user not found")
}

func deleteUser(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil {
			return err
		}
		createUser(user)
		return c.JSON(http.StatusCreated, user)
	})

	e.GET("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		user, err := getUser(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	})

	e.PUT("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var user User
		if err := c.Bind(&user); err != nil {
			return err
		}
		user.ID = id
		if err := updateUser(id, user); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, user)
	})

	e.DELETE("/users/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := deleteUser(id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
