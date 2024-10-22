package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	grubgo "github.com/IamNirvan/grubgo-rule-engine/internal/app"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	webserver "github.com/IamNirvan/grubgo-rule-engine/internal/pkg/web_server"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Create context that listens for stop signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	config := config.LoadConfig()
	// log.Debugf("config loaded: %v", config)

	// Initialize web server
	webServer := webserver.New(config)

	// Create instance of app
	app := grubgo.New(config, webServer)
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to run application: %s", err)
	}
}
