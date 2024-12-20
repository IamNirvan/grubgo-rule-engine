package webserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WebServer struct {
	Config  *config.Config
	Handler *handlers.Handler
	Server  *http.Server
}

type Options struct {
	Config  *config.Config
	Handler *handlers.Handler
}

func New(options *Options) *WebServer {
	return &WebServer{
		Config:  (*options).Config,
		Handler: (*options).Handler,
	}
}

func (ws *WebServer) Start() error {
	log.Debug("starting web server")

	r := gin.Default()

	r.POST("/v1/evaluate/rule", ws.Handler.HandleRuleEvaluationRequest)
	r.GET("/v1/specification", ws.Handler.HandleGetSpec)

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
