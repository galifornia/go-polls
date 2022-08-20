package routes

import (
	"github.com/galifornia/go-polls-voting/cmd/api/handlers"
	"github.com/labstack/echo/v4"
)

func SetupUsersRoutes(e *echo.Echo) {
	// Log in user
	e.POST("/users/login", handlers.LogIn)
	// Sign out user
	e.POST("/users/logout", handlers.LogOut)
	// Sign up user
	e.POST("/users/signup", handlers.Register)
	// Get user info
	e.GET("/users/:id", handlers.GetUser)
	// Delete user
	e.DELETE("/users/:id", handlers.RemoveUser)
	// Update user info
	e.PUT("/users/:id", handlers.UpdateUserInfo)
}
