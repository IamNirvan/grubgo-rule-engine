package facts

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

// func (dd *DishDetails) IsDishEmpty() bool {
// 	return dd.Dish == nil
// }

func (dd *DishDetails) AddResponse(response interface{}) {
	dd.Responses = append(dd.Responses, response)
}

func (dd *DishDetails) IsResponseEmpty() bool {
	return len(dd.Responses) == 0
}

func (dd *DishDetails) StringListsHaveMatchingItem(listA []string, listB []string) bool {
	var foundItem bool = false

	for i := range listA {
		for x := range listB {
			if listA[i] == listB[x] {
				foundItem = true
				break
			}
		}
	}
	return foundItem
}
