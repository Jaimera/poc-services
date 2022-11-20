package handler

import (
	"context"
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/domain/dto"
	"github.com/jaimera/poc-services/server/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Handles a fetch of port by code
func HandleGetPorts(portService contract.PortService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		ct := ctx.(context.Context)

		code := ctx.Param("code")
		port, err := portService.GetByCode(ct, code)
		if err != nil {
			return ctx.String(http.StatusNotFound, "")
		}

		return ctx.JSON(http.StatusOK, port)
	}
}

// Handles ports.json insert
func HandleInsertPort(portService contract.PortService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var ports []dto.PortDto
		ct := ctx.Request().Context()

		var jsonMap map[string]map[string]interface{}
		err := ctx.Bind(&jsonMap)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "")
		}

		ports, _, err = utils.ConvertPortJson(jsonMap)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "")
		}

		err = portService.Insert(ct, ports)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "")
		}

		return ctx.String(http.StatusCreated, "OK")
	}
}
