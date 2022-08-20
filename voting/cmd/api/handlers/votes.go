package handlers

import (
	"net/http"

	"github.com/galifornia/go-polls-voting/models"
	"github.com/labstack/echo/v4"
)

// e.POST("/vote", handlers.CastVote)
func CastVote(ctx echo.Context) error {
	vote := new(models.Vote)
	if err := ctx.Bind(vote); err != nil {
		return err
	}
	// TODO: call database
	return ctx.JSON(http.StatusCreated, vote)
}

// e.DELETE("/vote/:id", handlers.DeleteVote)
func DeleteVote(ctx echo.Context) error {
	id := ctx.Param("id")
	// TODO: call database
	return ctx.String(http.StatusOK, id)
}

// e.UPDATE("/vote/:id", handlers.UpdateVote)
func UpdateVote(ctx echo.Context) error {
	vote := new(models.Vote)
	if err := ctx.Bind(vote); err != nil {
		return err
	}
	// TODO: call database
	return ctx.JSON(http.StatusAccepted, vote)
}

// e.Get("/vote/:poll_id", handlers.GetVotesForPoll)
func GetVotesForPoll(ctx echo.Context) error {
	poll_id := ctx.Param("poll_id")
	// TODO: call database
	return ctx.String(http.StatusOK, poll_id)
}
