package internal

import "math"

type Meal struct {
	items []MealItem
}

type MealItem struct {
	food     Food
	quantity int
}

func (meal *Meal) AddItem(food Food, quantity int) {
	meal.items = append(meal.items, MealItem{food, quantity})
}

func (meal *Meal) GetProtein() int {
	protein := 0
	for _, i := range meal.items {
		protein += int(math.Round(float64(i.quantity) / float64(100) * float64(i.food.NutritionalPanel.Protein)))
	}
	return protein
}

func (meal *Meal) GetCarbs() int {
	carbs := 0
	for _, i := range meal.items {
		carbs += int(math.Round(float64(i.quantity) / float64(100) * float64(i.food.NutritionalPanel.Carbs)))
	}
	return carbs
}
