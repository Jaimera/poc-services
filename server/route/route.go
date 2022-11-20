package route

import (
	"github.com/jaimera/poc-services/application"
	"github.com/jaimera/poc-services/server/handler"
	"github.com/labstack/echo/v4"
)

// Register HTTP API Routes
func Register(app application.App, server *echo.Echo) {

	server.GET("/health", handler.HandleHealthCheck())

	root := server.Group("/api/v1")

	port := root.Group("/ports")
	port.GET("", handler.HandleGetPorts(app.Services().Port()))
	port.POST("", handler.HandleInsertPort(app.Services().Port()))
}
