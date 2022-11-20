package handler

import (
	"github.com/jaimera/poc-services/domain/contract"
	"github.com/jaimera/poc-services/domain/dto"
	"github.com/jaimera/poc-services/server/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// HandleGetPorts fetch for a port by it's code
func HandleGetPort(portService contract.PortService) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		ct := ctx.Request().Context()

		code := ctx.Param("code")
		port, err := portService.GetByCode(ct, code)
		if err != nil {
			return ctx.String(http.StatusNotFound, "Not found")
		}

		// parse to dto

		return ctx.JSON(http.StatusOK, port)
	}
}

// HandleInsertPort receives a ports.json and attempts to queue
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

		go portService.Queue(ct, ports)

		return ctx.String(http.StatusOK, "Json Enqueued")
	}
}
