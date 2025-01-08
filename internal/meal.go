package internal

import (
	"math"
	"strings"
)
import "fmt"

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

func (m *Meal) PrintInConsole() {
	fmt.Println("Meal")

	// Convert meal items to strings
	var foodStrings []string
	for _, item := range m.items {
		foodString := fmt.Sprintf("%dg of %s", item.quantity, strings.ToLower(item.food.Name))
		foodStrings = append(foodStrings, foodString)
	}

	// Join all portions with appropriate separators
	switch len(foodStrings) {
	case 0:
		fmt.Println("Empty meal")
	case 1:
		fmt.Println(foodStrings[0])
	case 2:
		fmt.Printf("%s and %s\n", foodStrings[0], foodStrings[1])
	default:
		lastItem := foodStrings[len(foodStrings)-1]
		otherItems := foodStrings[:len(foodStrings)-1]
		fmt.Printf("%s, and %s\n", strings.Join(otherItems, ", "), lastItem)
	}
}
