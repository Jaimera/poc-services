package server

import (
	"github.com/jaimera/poc-services/application"
	"github.com/jaimera/poc-services/server/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Run setup and executes the server
func Run(app application.App) {
	server := newServer(app)
	route.Register(app, server)

	err := server.Start(":8080")
	if err != nil {
		app.Logger.Errorf("Could not start the server. Error: %s.", err.Error())
	}
}

// Setup configures the service's Echo (REST) server
func newServer(app application.App) *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HideBanner = true
	e.Debug = true

	return e
}
