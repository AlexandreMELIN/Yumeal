package internal

import "math/rand/v2"
import "math"

type DailyMeals struct {
	Breakfast Meal
	Lunch     Meal
	Dinner    Meal
}

func PlanAWeek(breakfastPreference MealPreference, mealsPreference MealPreference, targetForTheDay NutritionalPanel) [7]DailyMeals {
	// Initialize week plan
	var weekPlan [7]DailyMeals

	// Plan each day
	for day := 0; day < 7; day++ {
		meals := PlanADay(breakfastPreference, mealsPreference, targetForTheDay)
		weekPlan[day] = DailyMeals{
			Breakfast: meals[0],
			Lunch:     meals[1],
			Dinner:    meals[2],
		}
	}

	return weekPlan
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

func PlanADay(breakfastPreference MealPreference, mealsPreference MealPreference, targetForTheDay NutritionalPanel) []Meal {
	// Calculate macro splits for each meal
	breakfastMacros := NutritionalPanel{
		Protein: int(float64(targetForTheDay.Protein) * 0.2),
		Carbs:   int(float64(targetForTheDay.Carbs) * 0.2),
		Fat:     int(float64(targetForTheDay.Fat) * 0.2),
		Fiber:   int(float64(targetForTheDay.Fiber) * 0.2),
	}

	lunchMacros := NutritionalPanel{
		Protein: int(float64(targetForTheDay.Protein) * 0.3),
		Carbs:   int(float64(targetForTheDay.Carbs) * 0.3),
		Fat:     int(float64(targetForTheDay.Fat) * 0.3),
		Fiber:   int(float64(targetForTheDay.Fiber) * 0.3),
	}

	dinnerMacros := NutritionalPanel{
		Protein: int(float64(targetForTheDay.Protein) * 0.5),
		Carbs:   int(float64(targetForTheDay.Carbs) * 0.5),
		Fat:     int(float64(targetForTheDay.Fat) * 0.5),
		Fiber:   int(float64(targetForTheDay.Fiber) * 0.5),
	}

	// Plan each meal
	breakfast := PlanBreakfast(breakfastPreference, breakfastMacros)
	lunch := PlanAMeal(mealsPreference, lunchMacros)
	dinner := PlanAMeal(mealsPreference, dinnerMacros)

	return []Meal{breakfast, lunch, dinner}
}

func determineQuantityToMatchMacro(macroFor100g int, target int) int {
	return int(math.Round(float64(target) / float64(macroFor100g) * 100))
}

func computeMacro(quantity int, reference int) int {
	return int(math.Round(float64(quantity) / float64(100) * float64(reference)))
}
func selectFromSlice[T any](items []T) T {
	if len(items) == 1 {
		return items[0]
	}
	return items[rand.IntN(len(items)-1)]
}
