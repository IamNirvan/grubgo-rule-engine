package ruleengine

import (
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	"github.com/hyperjumptech/grule-rule-engine/engine"
)

type RuleEngine struct {
	Config *config.Config
	Engine *engine.GruleEngine
}

func New(cfg *config.Config) *RuleEngine {
	return &RuleEngine{
		Config: cfg,
		Engine: engine.NewGruleEngine(),
	}
}
