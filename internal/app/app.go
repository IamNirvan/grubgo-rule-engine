package grubgo

import (
	"context"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/rule_engine/library"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/util"
	webserver "github.com/IamNirvan/grubgo-rule-engine/internal/pkg/web_server"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type Grubgo struct {
	Config    *config.Config
	WebServer *webserver.WebServer
	Database  *gorm.DB
}

func New(config *config.Config, webServer *webserver.WebServer, db *gorm.DB) *Grubgo {
	return &Grubgo{
		Config:    config,
		WebServer: webServer,
		Database:  db,
	}
}

func (g *Grubgo) Start(ctx context.Context) error {
	errorGroup, ctx := errgroup.WithContext(ctx)

	// Load all the rules from the database
	library.New(g.Database)

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
