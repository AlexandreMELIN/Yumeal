package nutrition

import (
	"testing"
	"time"
)

func TestMealPlannerPlanCorrectlyADay(t *testing.T) {
	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	mealPlanner := NewPlannerForAWeek(time.Now())
	recipe := createRecipeForMealPlanner()
	mealPlanner.AddBreakfast(tomorrow, recipe)
	mealPlanner.AddMorningSnack(tomorrow, recipe)
	mealPlanner.AddLunch(tomorrow, recipe)
	mealPlanner.AddAfternoonSnack(tomorrow, recipe)
	mealPlanner.AddDinner(tomorrow, recipe)

	plan := mealPlanner.Plan()

	wantedNutritionalComposition := NewNutritionalComposition(29*5, 32*5, 68*5)
	if plan.NutritionalCompositionPerDay[tomorrow] != *wantedNutritionalComposition {
		t.Errorf("Nutritional Composition not computed correctly. Expected %s, got %s", wantedNutritionalComposition, plan.NutritionalCompositionPerDay[tomorrow])
	}

	wantedGroceryList := GroceryList{groceries: map[Ingredient]float64{
		Ingredient{name: "pasta", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(8, 0, 60)}}:   500,
		Ingredient{name: "cheese", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(21, 32, 8)}}: 500,
	}}
	gottenGroceryList := plan.GroceryList

	for ingredient, quantity := range wantedGroceryList.groceries {
		if gottenGroceryList.groceries[ingredient] != quantity {
			t.Errorf("4aa62b17-d908-4cc7-84a2-20f3da368d83 expected: %f , got: %f, ingredient: %s", quantity, gottenGroceryList.groceries[ingredient], ingredient.name)
		}
	}

}

func createRecipeForMealPlanner() Recipe {
	recipe := Recipe{name: "Recipe for test", recipeType: LUNCH, ingredients: map[Ingredient]float64{
		Ingredient{name: "pasta", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(8, 0, 60)}}:   100,
		Ingredient{name: "cheese", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(21, 32, 8)}}: 100,
	}}

	return recipe
}
