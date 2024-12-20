package facts

import (
	"strings"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/constants"
	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/models/types"
)

const (
	DISH_DETAILS_FACT_ALIAS = "DDF"
)

type Customer struct {
	Id        uint64   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Allergens []string `json:"allergens"`
}

type Dish struct {
	Id          uint64   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
}

type DishDetails struct {
	Dish      Dish          `json:"dish"`
	Customer  Customer      `json:"customer"`
	Responses []interface{} `json:"responses"`
}

func NewDishDetailsFact(dish Dish, customer Customer) *DishDetails {
	return &DishDetails{
		Dish:      dish,
		Customer:  customer,
		Responses: []interface{}{},
	}
}

func (dd *DishDetails) AddResponse(response interface{}) {
	dd.Responses = append(dd.Responses, response)
}

func (dd *DishDetails) IsResponseEmpty() bool {
	return len(dd.Responses) == 0
}

func (dd *DishDetails) StringListsHaveMatchingItem(listA []string, listB []string) bool {
	for _, itemA := range listA {
		for _, itemB := range listB {
			if strings.EqualFold(itemA, itemB) {
				return true
			}
		}
	}
	return false
}

func (dd *DishDetails) AddResponseComponent(componentType string, status string, text string) {
	dd.Responses = append(dd.Responses, &types.RuleEngineResponse{
		Type: constants.RULE_ENGINE_RESPONSE_TYPE_COMPONENT,
		Payload: &types.Component{
			Type:   componentType,
			Status: status,
			Text:   text,
		},
	})
}

// func (dd *DishDetails) AddResponseComponent(componentType string, moodType string, text string) {
// 	dd.Responses = append(dd.Responses, &types.RuleEngineResponse{
// 		Type: constants.RULE_ENGINE_RESPONSE_TYPE_COMPONENT,
// 		Payload: &types.Component{
// 			Type: componentType,
// 			Mood: moodType,
// 			Text: text,
// 		},
// 	})
// }
