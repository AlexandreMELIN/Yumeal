package nutrition

import "fmt"

type NutritionalComposition struct {
	Calories      float64
	Proteins      NutritionalFacts
	Fat           NutritionalFacts
	Carbohydrates NutritionalFacts
}

func (n NutritionalComposition) String() string {
	return fmt.Sprintf("proteins: %f, fat: %f, carbohydrates: %f, calories: %f", n.Proteins.Quantity, n.Fat.Quantity, n.Carbohydrates.Quantity, n.Calories)
}

func NewNutritionalComposition(proteinsQuantity float64, fatQuantity float64, carbohydratesQuantity float64) *NutritionalComposition {
	nutritionalComposition := NutritionalComposition{Proteins: NutritionalFacts{Quantity: proteinsQuantity, Nutrient: Protein},
		Fat:           NutritionalFacts{Quantity: fatQuantity, Nutrient: Fat},
		Carbohydrates: NutritionalFacts{Quantity: carbohydratesQuantity, Nutrient: Carbohydrate}}
	nutritionalComposition.computeCalories()
	return &nutritionalComposition
}
func (n *NutritionalComposition) computeCalories() {
	n.Calories = n.Fat.ComputeCalories() + n.Proteins.ComputeCalories() + n.Carbohydrates.ComputeCalories()
}
