package main

import (
	"github.com/labstack/echo/v4"
)

const PORT = "80"

type App struct {
}

func main() {
	app := App{}

	e := echo.New()
	app.setupRoutes(e)

	e.Logger.Fatal(e.Start(":" + PORT))
}
