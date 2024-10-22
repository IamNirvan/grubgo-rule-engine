package handlers

import (
	"sync"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/rule_engine/facts"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

var (
	instance    *Handler
	HandlerOnce sync.Once
)

type Handler struct {
	Config  *config.Config
	Service *service.Service
}

type Options struct {
	Config  *config.Config
	Service *service.Service
}

func New(options *Options) *Handler {
	HandlerOnce.Do(func() {
		instance = &Handler{
			Config:  (*options).Config,
			Service: (*options).Service,
		}
	})
	return instance
}

func (h *Handler) HandleRuleEvaluationRequest(c *gin.Context) {
	var fact facts.DishDetails
	if err := c.ShouldBindJSON(&fact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := (*h.Service.RuleEngineService).EvaluateRule(&fact, c)
	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Error})
		return
	}

	c.JSON(200, response)
}
