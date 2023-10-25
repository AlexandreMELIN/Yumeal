package nutrition

type NutritionalFacts struct {
	Nutrient Nutrient
	Quantity float64
}

func (n NutritionalFacts) ComputeCalories() float64 {
	return float64(n.Nutrient.CaloriesPer100Grams) * n.Quantity
}
