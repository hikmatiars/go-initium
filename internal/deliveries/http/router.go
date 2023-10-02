package http

import (
	"github.com/labstack/echo/v4"
	"go-initium/internal/deliveries/http/v1/healthcheck"
	"go-initium/internal/deliveries/http/v1/user"
)

func (h *Http) Route(e *echo.Echo) {
	//Register all group route
	group := e.Group("/api/v1")
	user.Route(group.Group("/user"))
	healthcheck.Route(group.Group("/health-check"))
}
