package webserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WebServer struct {
	Config *config.Config
	Server *http.Server
}

func New(config *config.Config) *WebServer {
	return &WebServer{
		Config: config,
	}
}

func (ws *WebServer) Start() error {
	log.Debug("starting web server")

	r := gin.Default()

	// TODO: add correct route here to evaluate a rule(s) using a fact
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ws.Server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", ws.Config.WebServer.Host, ws.Config.WebServer.Port),
		Handler: r,
	}

	return ws.Server.ListenAndServe()
}

func (ws *WebServer) Dispose(ctx context.Context) error {
	log.Debug("stopping web server")
	return ws.Server.Shutdown(ctx)
}
