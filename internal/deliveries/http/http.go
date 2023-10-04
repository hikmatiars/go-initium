package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-initium/cmd/server"
	"go-initium/internal/config"
	"log"
)

type (
	RegisterHttp interface {
		Start() *echo.Echo
		Route(group *echo.Echo)
		Stop(ctx context.Context)
	}

	Http struct {
		e        *echo.Echo
		addr     string
		Contract *server.Contract
	}
)

func NewHttp(c *echo.Echo, cfg config.Config, contract *server.Contract) RegisterHttp {
	addr := fmt.Sprintf(":%d", cfg.App.HTTPPort)

	return &Http{
		e:        c,
		addr:     addr,
		Contract: contract,
	}
}

func (h *Http) Start() *echo.Echo {
	h.Route(h.e)
	if err := h.e.Start(h.addr); err != nil {
		log.Fatal("service shutting down")
	}

	return h.e
}

func (h *Http) Stop(ctx context.Context) {
	if err := h.e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
