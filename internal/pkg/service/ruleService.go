package service

import (
	"context"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	ruleengine "github.com/IamNirvan/grubgo-rule-engine/internal/pkg/rule_engine"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/rule_engine/facts"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/rule_engine/library"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RuleEngineService interface {
	EvaluateRule(*facts.DishDetails, context.Context) (*[]interface{}, *ServiceError)
}

type RuleEngineServiceV1 struct {
	Config   *config.Config
	Database *gorm.DB
}

func NewRuleEngineServiceV1(config *config.Config, db *gorm.DB) *RuleEngineService {
	var service RuleEngineService = &RuleEngineServiceV1{
		Config:   config,
		Database: db,
	}
	return &service
}

func (rs *RuleEngineServiceV1) EvaluateRule(fact *facts.DishDetails, ctx context.Context) (*[]interface{}, *ServiceError) {
	// Get instance of library
	lib := library.New()

	// Fetch knowledge base
	knowledgeBase, knowledgeBaseErr := lib.NewKnowledgeBaseInstance(library.KNOWLEDGE_BASE_NAME, library.VERSION)
	if knowledgeBaseErr != nil {
		log.Errorf("error when obtaining instance of KnowledgeBase: %v", knowledgeBaseErr.Error())
		return nil, &ServiceError{Error: knowledgeBaseErr.Error(), Status: 500}
	}

	// Create new instance of rule engine
	engine := ruleengine.New(rs.Config)

	// Create a data context
	dataCtx := ast.NewDataContext()
	if dataCtxErr := dataCtx.Add(facts.DISH_DETAILS_FACT_ALIAS, fact); dataCtxErr != nil {
		log.Errorf("error when adding fact to data context: %v", dataCtxErr.Error())
		return nil, &ServiceError{Error: dataCtxErr.Error(), Status: 500}
	}

	// Evaluate rule(s)
	if engineErr := engine.Engine.Execute(dataCtx, knowledgeBase); engineErr != nil {
		log.Errorf("error when evaluating rule: %v", engineErr.Error())
		return nil, &ServiceError{Error: engineErr.Error(), Status: 500}
	}

	// Fetch the response from the fact
	response := &fact.Responses

	return response, nil
}
