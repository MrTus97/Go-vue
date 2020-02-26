package main

import (
	service "echo-example-crud"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.POST("/users", service.CreateUser)
	e.GET("/users/:id", service.GetUser)
	e.PUT("/users/:id", service.UpdateUser)
	e.DELETE("/users/:id", service.DeleteUser)
	e.GET("/", service.TestIndex)

	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}
