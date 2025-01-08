package internal

import "math/rand/v2"
import "math"

func PlanADay(breakfastPreference MealPreference) {

}

func selectFromSlice[T any](items []T) T {
	if len(items) == 1 {
		return items[0]
	}
	return items[rand.IntN(len(items)-1)]
}

func PlanBreakfast(preference MealPreference, targetMacro NutritionalPanel) Meal {
	targetProtein := targetMacro.Protein
	targetCarbs := targetMacro.Carbs
	meal := Meal{}

	hasProtein := len(preference.Protein) > 0
	hasDairy := len(preference.Dairy) > 0
	selectedCarb := selectFromSlice(preference.Starch)
	quantity := determineQuantityToMatchMacro(selectedCarb.NutritionalPanel.Carbs, targetCarbs)
	targetCarbs -= computeMacro(quantity, selectedCarb.NutritionalPanel.Carbs)
	meal.AddItem(selectedCarb, quantity)

	if hasProtein && hasDairy {
		selectedProtein := selectFromSlice(preference.Protein)
		quantity := determineQuantityToMatchMacro(selectedProtein.NutritionalPanel.Protein, int(float64(targetProtein)*0.7))
		targetCarbs -= computeMacro(quantity, selectedProtein.NutritionalPanel.Carbs)
		targetProtein -= computeMacro(quantity, selectedProtein.NutritionalPanel.Protein)
		meal.AddItem(selectedProtein, quantity)

		selectedDairy := selectFromSlice(preference.Dairy)
		quantity = determineQuantityToMatchMacro(selectedDairy.NutritionalPanel.Protein, int(float64(targetProtein)*0.3))
		targetCarbs -= computeMacro(quantity, selectedDairy.NutritionalPanel.Carbs)
		targetProtein -= computeMacro(quantity, selectedDairy.NutritionalPanel.Protein)
		meal.AddItem(selectedDairy, quantity)
	}
	return meal
}

func PlanAMeal(preference MealPreference, targetMacro NutritionalPanel) Meal {
	targetProtein := targetMacro.Protein
	targetCarbs := targetMacro.Carbs
	meal := Meal{}

	shouldIncludePlants := len(preference.Plant) > 0 && rand.IntN(1) == 1
	if shouldIncludePlants {
		selectedPlant := selectFromSlice(preference.Plant)
		quantity := determineQuantityToMatchMacro(selectedPlant.NutritionalPanel.Carbs, targetCarbs/2)
		targetCarbs -= computeMacro(quantity, selectedPlant.NutritionalPanel.Carbs)
		targetProtein -= computeMacro(quantity, selectedPlant.NutritionalPanel.Protein)
		meal.AddItem(selectedPlant, quantity)
	}

	selectedStarch := selectFromSlice(preference.Starch)
	quantity := 0
	if shouldIncludePlants {
		quantity = determineQuantityToMatchMacro(selectedStarch.NutritionalPanel.Carbs, targetCarbs/2)
	} else {
		quantity = determineQuantityToMatchMacro(selectedStarch.NutritionalPanel.Carbs, targetCarbs)
	}
	targetCarbs -= computeMacro(quantity, selectedStarch.NutritionalPanel.Carbs)
	targetProtein -= computeMacro(quantity, selectedStarch.NutritionalPanel.Protein)
	meal.AddItem(selectedStarch, quantity)

	selectedProtein := selectFromSlice(preference.Protein)
	quantity = determineQuantityToMatchMacro(selectedProtein.NutritionalPanel.Protein, targetProtein)
	targetCarbs -= computeMacro(quantity, selectedProtein.NutritionalPanel.Carbs)
	targetProtein -= computeMacro(quantity, selectedProtein.NutritionalPanel.Protein)
	meal.AddItem(selectedProtein, quantity)

	selectedVegetable := selectFromSlice(preference.Vegetables)
	quantity = 80 * (rand.IntN(4) + 1)
	meal.AddItem(selectedVegetable, quantity)

	if len(preference.Fruits) > 0 && rand.IntN(1) == 1 {
		selectedFruit := selectFromSlice(preference.Fruits)
		quantity = 80 * (rand.IntN(4) + 1)
		meal.AddItem(selectedFruit, quantity)
	}
	return meal
}

func determineQuantityToMatchMacro(macroFor100g int, target int) int {
	return int(math.Round(float64(target) / float64(macroFor100g) * 100))
}

func computeMacro(quantity int, reference int) int {
	return int(math.Round(float64(quantity) / float64(100) * float64(reference)))
}
