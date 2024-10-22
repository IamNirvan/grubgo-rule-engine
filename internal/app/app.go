package grubgo

import (
	"context"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/util"
	webserver "github.com/IamNirvan/grubgo-rule-engine/internal/pkg/web_server"
	"golang.org/x/sync/errgroup"
)

type Grubgo struct {
	Config    *config.Config
	WebServer *webserver.WebServer
}

func New(config *config.Config, webServer *webserver.WebServer) *Grubgo {
	return &Grubgo{
		Config:    config,
		WebServer: webServer,
	}
}

func (g *Grubgo) Start(ctx context.Context) error {
	errorGroup, ctx := errgroup.WithContext(ctx)

	// Create list of disposable resources
	disposables := []util.Disposable{g.WebServer}

	// Run the web server in a goroutine
	errorGroup.Go(func() error {
		return g.WebServer.Start()
	})

	// Create error message channel
	errChan := make(chan error)
	go func() {
		if err := errorGroup.Wait(); err != nil {
			errChan <- err
		}
	}()

	for {
		select {
		case <-ctx.Done():
			// Dispose all resources
			for index := range disposables {
				if err := disposables[index].Dispose(ctx); err != nil {
					return err
				}
			}
			return nil
		case err := <-errChan:
			return err
		}
	}
}
