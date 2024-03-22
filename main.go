package main

import (
	"Echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/users", routes.GetAllUsers)
	e.GET("/users/:id", routes.GetUserByID)
	e.POST("/users", routes.InsertUser)
	e.PUT("/users", routes.UpdateUser)
	e.DELETE("/users/:id", routes.DeleteUser)

	e.Logger.Fatal(e.Start(":8888"))
}
