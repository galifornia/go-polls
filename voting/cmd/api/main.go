package main

import (
	"github.com/galifornia/go-polls-voting/database"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const PORT = "80"

type App struct {
	DB *pgxpool.Pool
}

func main() {
	app := App{}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	app.setupRoutes(e)

	db := database.ConnectToDB()
	app.DB = db

	e.Logger.Fatal(e.Start(":" + PORT))
}
