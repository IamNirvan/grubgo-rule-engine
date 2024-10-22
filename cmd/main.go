package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grubgo "github.com/IamNirvan/grubgo-rule-engine/internal/app"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/handlers"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/service"
	webserver "github.com/IamNirvan/grubgo-rule-engine/internal/pkg/web_server"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create context that listens for stop signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := config.LoadConfig()

	// Inialize service instance
	service := service.New(&service.Options{
		Config: config,
	})

	// Initialize handler instance
	handler := handlers.New(&handlers.Options{
		Config:  config,
		Service: service,
	})

	// Initialize web server
	webServer := webserver.New(&webserver.Options{
		Config:  config,
		Handler: handler,
	})

	// Create instance of app
	app := grubgo.New(config, webServer)
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to run application: %s", err)
	}
}
