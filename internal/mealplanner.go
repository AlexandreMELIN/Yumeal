package internal

import "math/rand/v2"
import "math"

func PlanAMeal(preference MealPreference, targetMacro NutritionalPanel) Meal {

	targetProtein := targetMacro.Protein
	targetCarbs := targetMacro.Carbs
	meal := Meal{}

	shouldIncludePlants := len(preference.Plant) > 0 && rand.IntN(1) == 1
	if shouldIncludePlants {
		selectedPlant := preference.Plant[rand.IntN(len(preference.Plant)-1)]
		quantity := determineQuantityToMatchMacro(selectedPlant.NutritionalPanel.Carbs, targetCarbs/2)
		targetCarbs -= computeMacro(quantity, selectedPlant.NutritionalPanel.Carbs)
		targetProtein -= computeMacro(quantity, selectedPlant.NutritionalPanel.Protein)
		meal.AddItem(selectedPlant, quantity)
	}

	selectedStarch := preference.Starch[rand.IntN(len(preference.Starch)-1)]
	quantity := 0
	if shouldIncludePlants {
		quantity = determineQuantityToMatchMacro(selectedStarch.NutritionalPanel.Carbs, targetCarbs/2)
	} else {
		quantity = determineQuantityToMatchMacro(selectedStarch.NutritionalPanel.Carbs, targetCarbs)
	}
	targetCarbs -= computeMacro(quantity, selectedStarch.NutritionalPanel.Carbs)
	targetProtein -= computeMacro(quantity, selectedStarch.NutritionalPanel.Protein)
	meal.AddItem(selectedStarch, quantity)

	selectedProtein := preference.Protein[rand.IntN(len(preference.Protein)-1)]
	quantity = determineQuantityToMatchMacro(selectedProtein.NutritionalPanel.Protein, targetProtein)
	targetCarbs -= computeMacro(quantity, selectedProtein.NutritionalPanel.Carbs)
	targetProtein -= computeMacro(quantity, selectedProtein.NutritionalPanel.Protein)
	meal.AddItem(selectedProtein, quantity)

	return meal
}

func determineQuantityToMatchMacro(macroFor100g int, target int) int {
	return int(math.Round(float64(target) / float64(macroFor100g) * 100))
}

func computeMacro(quantity int, reference int) int {
	return int(math.Round(float64(quantity) / float64(100) * float64(reference)))
}
