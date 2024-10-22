package service

import (
	"sync"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
)

var (
	instance    *Service
	ServiceOnce sync.Once
)

type ServiceError struct {
	Error  string
	Status int
}

type Service struct {
	RuleEngineService *RuleEngineService
}

type Options struct {
	Config *config.Config
}

func New(options *Options) *Service {
	ServiceOnce.Do(func() {
		instance = &Service{
			RuleEngineService: NewRuleEngineServiceV1((*options).Config),
		}
	})
	return instance
}
