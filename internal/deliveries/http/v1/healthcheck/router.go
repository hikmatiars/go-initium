package healthcheck

import "github.com/labstack/echo/v4"

func Route(g *echo.Group) {
	h := NewHealthCheckHandler()
	g.GET("/status", h.HttpHealth)
}
