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
		ingredientNutritionalComposition := ingredient.nutritionalComposition
		ratioToApply := quantity / REFEREENCE_QUANTITY_IN_GRAM
		result.CarbohydratesQuantity.Quantity += ratioToApply * ingredientNutritionalComposition.CarbohydratesQuantity.Quantity
		result.FatQuantity.Quantity += ratioToApply * ingredientNutritionalComposition.FatQuantity.Quantity
		result.ProteinsQuantity.Quantity += ratioToApply * ingredientNutritionalComposition.ProteinsQuantity.Quantity
	}
	return *result
}
