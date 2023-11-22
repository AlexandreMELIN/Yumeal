package nutrition

import "errors"

type GroceryList struct {
	groceries map[Ingredient]float64
}

func NewGroceryList() *GroceryList {
	return &GroceryList{make(map[Ingredient]float64)}
}

func (gl *GroceryList) AddGrocery(ingredient Ingredient, quantity float64) error {
	if quantity < 0 {
		return errors.New("1d6b8175-23ff-441a-86b2-8ffd4d411b6e - Cannot add grocery with negative quantity")
	}
	if _, exist := gl.groceries[ingredient]; exist {
		gl.groceries[ingredient] += quantity
	} else {
		gl.groceries[ingredient] = quantity
	}
	return nil
}
