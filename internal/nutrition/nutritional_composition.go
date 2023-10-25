package nutrition

type NutritionalComposition struct {
	CaloriesQuantity      float64
	ProteinsQuantity      NutritionalFacts
	FatQuantity           NutritionalFacts
	CarbohydratesQuantity NutritionalFacts
}

func NewNutritionalComposition(proteinsQuantity float64, fatQuantity float64, carbohydratesQuantity float64) *NutritionalComposition {
	nutritionalComposition := NutritionalComposition{ProteinsQuantity: NutritionalFacts{Quantity: proteinsQuantity, Nutrient: Protein},
		FatQuantity:           NutritionalFacts{Quantity: fatQuantity, Nutrient: Fat},
		CarbohydratesQuantity: NutritionalFacts{Quantity: carbohydratesQuantity, Nutrient: Carbohydrate}}
	nutritionalComposition.computeCalories()
	return &nutritionalComposition
}
func (n *NutritionalComposition) computeCalories() {
	n.CaloriesQuantity = n.FatQuantity.ComputeCalories() + n.ProteinsQuantity.ComputeCalories() + n.CarbohydratesQuantity.ComputeCalories()
}

const (
	REFEREENCE_QUANTITY_IN_GRAM = 100
)
