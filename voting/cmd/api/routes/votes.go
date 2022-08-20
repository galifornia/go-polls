package routes

import (
	"github.com/galifornia/go-polls-voting/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetupVotesRoutes(e *echo.Echo) {
	// Cast vote
	e.POST("/vote", handlers.CastVote)
	// Delete vote
	e.DELETE("/vote/:id", handlers.DeleteVote)
	// Update vote if poll allows it
	e.PUT("/vote/:id", handlers.UpdateVote)
	// Get all the votes for a concrete poll_id
	e.GET("/vote/:poll_id", handlers.GetVotesForPoll)
}
