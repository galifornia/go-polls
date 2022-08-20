package handlers

import (
	"net/http"

	"github.com/galifornia/go-polls-voting/models"
	"github.com/labstack/echo/v4"
)

// e.GET("/polls", handlers.GetAllPolls)
func GetAllPolls(ctx echo.Context) error {
	// TODO: call database
	return ctx.String(http.StatusOK, "")
}

// e.POST("/polls", handlers.NewPoll)
func NewPoll(ctx echo.Context) error {
	poll := new(models.Poll)
	if err := ctx.Bind(poll); err != nil {
		return err
	}
	// TODO: call database
	return ctx.JSON(http.StatusCreated, poll)
}

// e.GET("/polls/:id", handlers.GetPoll)
func GetPoll(ctx echo.Context) error {
	id := ctx.Param("id")
	// TODO: call database
	return ctx.String(http.StatusOK, id)
}

// e.DELETE("/polls/:id", handlers.DeletePoll)
func DeletePoll(ctx echo.Context) error {
	id := ctx.Param("id")
	// TODO: call database
	return ctx.String(http.StatusOK, id)
}

// e.UPDATE("/polls/:id", handlers.UpdatePoll)
func UpdatePoll(ctx echo.Context) error {
	poll := new(models.Poll)
	if err := ctx.Bind(poll); err != nil {
		return err
	}
	// TODO: call database
	return ctx.JSON(http.StatusAccepted, poll)
}
