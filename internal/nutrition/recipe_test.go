package nutrition

import "testing"

func TestNutritionalCompositionIsCorrect(t *testing.T) {
	recipeTested := createRecipeForTest()
	got := recipeTested.GetNutritionalComposition()
	want := NewNutritionalComposition(29, 32, 68)
	if got != *want {
		t.Errorf("Mistmatch between the expected and calculated nutriotional composition. Expected %s, got : %s", want, got)
	}
}

func createRecipeForTest() Recipe {
	recipe := Recipe{name: "Recipe for test", recipeType: LUNCH, ingredients: map[Ingredient]float64{
		Ingredient{name: "pasta", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(8, 0, 60)}}:   100,
		Ingredient{name: "cheese", nutritionalPanel: NutritionalPanel{100, *NewNutritionalComposition(21, 32, 8)}}: 100,
	}}

	return recipe
}
