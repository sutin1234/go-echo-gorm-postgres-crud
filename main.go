package main

import (
	"fmt"
	"net/http"
	"os"
	database "postgres_api/database"
	routeuser "postgres_api/users/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

func main() {
	fmt.Println("Golang GORM Started!")

	e := echo.New()
	if ok := database.IsDevMode(); ok {
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		fmt.Println("Appication Run On Development Mode")
	}

	e.GET("", func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, "/")
		return nil
	})
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello Echo Framework Go!")
	})

	// User Group
	uGroup := e.Group("/api/user")
	routeuser.RegisterRoute(uGroup)

	// Start Services
	p := os.Getenv("PORT")
	if p == "" {
		p = database.GetPort()
	}
	port := ":" + p

	e.Logger.Fatal(e.Start(port))
}
