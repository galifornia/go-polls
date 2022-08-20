package main

import (
	"github.com/galifornia/go-polls-voting/cmd/api/routes"
	"github.com/labstack/echo/v4"
)

func (app *App) setupRoutes(e *echo.Echo) {
	routes.SetupPollsRoutes(e)
	routes.SetupVotesRoutes(e)
	routes.SetupUsersRoutes(e)
}
