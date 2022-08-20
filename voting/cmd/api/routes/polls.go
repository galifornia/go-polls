package routes

import (
	"github.com/galifornia/go-polls-voting/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetupPollsRoutes(e *echo.Echo) {
	// get all polls
	e.GET("/polls", handlers.GetAllPolls)
	// create poll
	e.POST("/polls", handlers.NewPoll)
	// get poll info
	e.GET("/polls/:id", handlers.GetPoll)
	// delete poll
	e.DELETE("/polls/:id", handlers.DeletePoll)
	// Update poll info
	e.PUT("/polls/:id", handlers.UpdatePoll)
}
