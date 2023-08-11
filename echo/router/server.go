package router

import (
	"net/http"
	"timtoronto634/go_web_framework_comparison/echo/handler"

	"timtoronto634/go_web_framework_comparison/echo/entity"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// users
	// e.POST("/users", saveUser)
	e.GET("/users/:id", handler.GetUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.GET("/show", handler.Show)
	e.POST("/save", handler.Save)

	e.POST("/users", func(c echo.Context) error {
		u := new(entity.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.Static("/static", "static")

	return e
}
