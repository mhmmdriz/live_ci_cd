package main

import (
	"ci_cd/database"
	"ci_cd/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
