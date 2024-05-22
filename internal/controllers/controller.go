package controllers

import (
	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	createDefault(repo database.Repo) Controller
	registerRoutes(e *echo.Echo) Controller
}

// New creates a new instance of the controller and initializes it
// with a Repo instance.
// It will also call the Controllers RegisterRoutes function to allow
// it to register any routes it may control.
func Initialize[T Controller](repo database.Repo, e *echo.Echo) {
	var controller T
	controller.
		createDefault(repo).
		registerRoutes(e)
}