package main

import (
	"github.com/labstack/echo/v4"
)

func (app *App) setupRoutes(e *echo.Echo) {
	// get all polls
	e.GET("/polls", GetAllPolls)
	// create poll
	e.POST("/polls", NewPoll)
	// get poll info
	e.GET("/polls/:id", GetPoll)
	// delete poll
	e.DELETE("/polls/:id", DeletePoll)
	// Update poll info
	e.PUT("/polls/:id", UpdatePoll)

	// Cast vote
	e.POST("/vote", CastVote)
	// Delete vote
	e.DELETE("/vote/:id", DeleteVote)
	// Update vote if poll allows it
	e.PUT("/vote/:id", UpdateVote)
	// Get all the votes for a concrete poll_id
	e.GET("/vote/:poll_id", GetVotesForPoll)

}
