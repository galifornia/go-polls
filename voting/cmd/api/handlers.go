package main

import (
	"net/http"

	"github.com/galifornia/go-polls-voting/models"

	"github.com/labstack/echo/v4"
)

// e.GET("/polls", handlers.GetAllPolls)
func GetAllPolls(c echo.Context) error {
	// TODO: call database
	return c.String(http.StatusOK, "")
}

// e.POST("/polls", handlers.NewPoll)
func NewPoll(c echo.Context) error {
	poll := new(models.Poll)
	if err := c.Bind(poll); err != nil {
		return err
	}
	// TODO: call database
	return c.JSON(http.StatusCreated, poll)
}

// e.GET("/polls/:id", handlers.GetPoll)
func GetPoll(c echo.Context) error {
	id := c.Param("id")
	// TODO: call database
	return c.String(http.StatusOK, id)
}

// e.DELETE("/polls/:id", handlers.DeletePoll)
func DeletePoll(c echo.Context) error {
	id := c.Param("id")
	// TODO: call database
	return c.String(http.StatusOK, id)
}

// e.UPDATE("/polls/:id", handlers.UpdatePoll)
func UpdatePoll(c echo.Context) error {
	poll := new(models.Poll)
	if err := c.Bind(poll); err != nil {
		return err
	}
	// TODO: call database
	return c.JSON(http.StatusAccepted, poll)
}

// e.POST("/vote", handlers.CastVote)
func CastVote(c echo.Context) error {
	vote := new(models.Vote)
	if err := c.Bind(vote); err != nil {
		return err
	}
	// TODO: call database
	return c.JSON(http.StatusCreated, vote)
}

// e.DELETE("/vote/:id", handlers.DeleteVote)
func DeleteVote(c echo.Context) error {
	id := c.Param("id")
	// TODO: call database
	return c.String(http.StatusOK, id)
}

// e.UPDATE("/vote/:id", handlers.UpdateVote)
func UpdateVote(c echo.Context) error {
	vote := new(models.Vote)
	if err := c.Bind(vote); err != nil {
		return err
	}
	// TODO: call database
	return c.JSON(http.StatusAccepted, vote)
}

// e.Get("/vote/:poll_id", handlers.GetVotesForPoll)
func GetVotesForPoll(c echo.Context) error {
	poll_id := c.Param("poll_id")
	// TODO: call database
	return c.String(http.StatusOK, poll_id)
}
