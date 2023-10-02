package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthHandler struct{}

func NewHealthCheckHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) HttpHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "OK")
}
