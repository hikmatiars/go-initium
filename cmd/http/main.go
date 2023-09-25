package main

import (
	"context"
	"go-initium/cmd/server"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go-initium/internal/config"
	"go-initium/internal/deliveries/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		log.Panic(err)
	}

	httpServ := http.NewHttp(e, cfg)
	server.New(ctx, cfg)

	go func() {
		httpServ.Start()
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServ.Stop(ctx)
	}()

	wg.Wait()
	log.Printf("Server shutting down")
}
