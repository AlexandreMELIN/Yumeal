package nutrition

import "fmt"

type Nutrient struct {
	Name                string
	CaloriesPer100Grams uint
}

func (n Nutrient) String() string {
	return fmt.Sprintf("%v has %d calories per 100 grams", n.Name, n.CaloriesPer100Grams)
}

// Enum like Nutrients
var (
	Protein      = Nutrient{Name: "Protein", CaloriesPer100Grams: 4}
	Carbohydrate = Nutrient{Name: "Carbohydrate", CaloriesPer100Grams: 4}
	Fat          = Nutrient{Name: "Fat", CaloriesPer100Grams: 9}
)
