package library

import (
	"sync"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	log "github.com/sirupsen/logrus"
)

var (
	instance    *ast.KnowledgeLibrary
	libraryOnce sync.Once
)

const (
	KNOWLEDGE_BASE_NAME = "dishDetails"
	VERSION             = "0.1"
)

func New() *ast.KnowledgeLibrary {
	// TODO: define response functions...
	libraryOnce.Do(func() {
		RULES := `
		rule sample "Just a sample rule" salience 10 {
			when
				(DDF.StringListsHaveMatchingItem(DDF.Dish.Ingredients, DDF.Customer.Allergens))
			then
				DDF.AddResponse("Customer is allergic to dish!!!");
				Retract("sample");
		}
		`

		instance = ast.NewKnowledgeLibrary()

		ruleBuilder := builder.NewRuleBuilder(instance)
		// TODO: load the rules from the correct location....
		if builderErr := ruleBuilder.BuildRuleFromResource(KNOWLEDGE_BASE_NAME, VERSION, pkg.NewBytesResource([]byte(RULES))); builderErr != nil {
			log.Fatalf("an error occurred when creating rules: %s", builderErr.Error())
		}
	})

	return instance
}
