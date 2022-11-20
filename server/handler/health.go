package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Health check needed to monitoring
func HandleHealthCheck() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	}
}
