package internal

import "testing"

func TestMeal_GetProtein(t *testing.T) {
	food := Food{
		Name: "Whole Pasta",
		NutritionalPanel: NutritionalPanel{
			Protein: 13,
			Carbs:   71,
			Fat:     2,
			Fiber:   7,
		},
		Category: Starch}
	meal := Meal{}
	meal.AddItem(food, 150)
	expectedProtein := 20
	got := meal.GetProtein()
	if got != expectedProtein {
		t.Errorf("meal.GetProtein() = %d, want %d", got, expectedProtein)
	}
}

func TestMeal_GetCarbs(t *testing.T) {
	food := Food{
		Name: "Whole Pasta",
		NutritionalPanel: NutritionalPanel{
			Protein: 13,
			Carbs:   71,
			Fat:     2,
			Fiber:   7,
		},
		Category: Starch}
	meal := Meal{}
	meal.AddItem(food, 150)
	expectedCarbs := 107
	got := meal.GetCarbs()
	if got != expectedCarbs {
		t.Errorf("meal.GetCarbs() = %d, want %d", got, expectedCarbs)
	}
}
