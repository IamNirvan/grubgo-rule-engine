package library

import (
	"strings"
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	instance    *ast.KnowledgeLibrary
	libraryOnce sync.Once
)

const (
	KNOWLEDGE_BASE_NAME = "dishDetails"
	VERSION             = "0.1"
)

func New(args ...*gorm.DB) *ast.KnowledgeLibrary {
	// libraryOnce.Do(func() {
	if len(args) == 0 || len(args) > 1 {
		log.Fatal("a single database instance is required to load all the rules")
	}

	db := *(args[0])

	// Load the loadedRules from the database
	var loadedRules []string
	if err := db.Table("rules").Select("rule").Find(&loadedRules).Error; err != nil {
		log.Fatalf("failed to load rules from database: %v", err)
	}
	log.Debugf("loaded %d rule(s)", len(loadedRules))

	// Use rules in the knowledge base
	var sb strings.Builder
	for _, rule := range loadedRules {
		sb.WriteString(rule)
		sb.WriteString("\n")
	}
	rules := sb.String()

	instance = ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(instance)
	if builderErr := ruleBuilder.BuildRuleFromResource(KNOWLEDGE_BASE_NAME, VERSION, pkg.NewBytesResource([]byte(rules))); builderErr != nil {
		log.Warnf("an error occurred when creating rules: %s", builderErr.Error())
	}
	// })

	return instance
}
