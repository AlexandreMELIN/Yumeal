package nutrition

type RecipeType int8

const (
	LUNCH     RecipeType = 1
	DINNER    RecipeType = 2
	BREAKFAST RecipeType = 3
	SNACK     RecipeType = 4
)

type Recipe struct {
	name        string
	ingredients map[Ingredient]float64
	recipeType  RecipeType
}

func (r Recipe) GetNutritionalComposition() NutritionalComposition {
	result := NewNutritionalComposition(0, 0, 0)
	for ingredient, quantity := range r.ingredients {
		ingredientNutritionalComposition := ingredient.nutritionalPanel
		ratioToApply := quantity / float64(ingredient.nutritionalPanel.referenceQuantityInGram)
		result.Carbohydrates.Quantity += ratioToApply * ingredientNutritionalComposition.composition.Carbohydrates.Quantity
		result.Fat.Quantity += ratioToApply * ingredientNutritionalComposition.composition.Fat.Quantity
		result.Proteins.Quantity += ratioToApply * ingredientNutritionalComposition.composition.Proteins.Quantity
	}
	result.computeCalories()
	return *result
}
