package app

import (
	"fmt"

	_ "github.com/jfavo/integra-partners-assessment-backend/docs"
	"github.com/jfavo/integra-partners-assessment-backend/internal/config"
	"github.com/jfavo/integra-partners-assessment-backend/internal/controllers"
	"github.com/jfavo/integra-partners-assessment-backend/internal/database"
	"github.com/jfavo/integra-partners-assessment-backend/internal/errors"
	"github.com/jfavo/integra-partners-assessment-backend/internal/logging"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/swaggo/echo-swagger"
)

// StartServer will create a new server instance and all dependent resources.
//
// Will throw panic if the DB repository fails to initialize.
func StartServer() {
	e := echo.New()

	// Configure middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	// Add swagger documentation page
	e.GET("/docs/*", echoSwagger.WrapHandler)

	config := config.New()

	// Initialize Database client
	repo, err := database.CreateNewRepo(config.Database)
	if err != nil {
		code := errors.DBRepoFailedToInitialize
		errMessage := errors.GetErrorMessage(code)
		logging.ErrorWithCode(
			code,
			errMessage,
			err)
		panic(errMessage)
	}

	// Initialize Controllers
	// This will create a new UserController struct, attach our DB repository
	// to it, and register its routes
	controllers.Initialize[controllers.UserController](repo, e)

	// Start the HTTP server, if it returns an error we will log it
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.Server.Port)))
}
