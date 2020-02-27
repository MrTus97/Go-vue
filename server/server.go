package main

import (
	"echo-example-crud/server/controller"
	"net/http"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
	// e.Use(middleware.CORS())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Routes
	e.GET("/users", controller.GetAllUser)
	e.POST("/users", controller.AddEmployee)
	e.GET("/users/:id", controller.GetEmployee)
	e.PUT("/users/:id", controller.EditEmployee)
	e.DELETE("/users/:id", controller.DeleteEmployee)
	e.GET("/", controller.TestIndex)

	// Start server at localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}
