package nutrition

import (
	"errors"
	"time"
)

type MealPlanner struct {
	start       time.Time
	end         time.Time
	plannedFor  int
	daysPlanned map[time.Time]MealPlannerDay
}

type MealPlannerDay struct {
	Breakfast      Recipe
	MorningSnack   Recipe
	Lunch          Recipe
	AfternoonSnack Recipe
	Dinner         Recipe
}

type MealPlan struct {
	GroceryList                  GroceryList
	NutritionalCompositionPerDay map[time.Time]NutritionalComposition
}

func NewMealPlan() *MealPlan {
	return &MealPlan{GroceryList: *NewGroceryList(), NutritionalCompositionPerDay: make(map[time.Time]NutritionalComposition)}
}

func NewPlannerForAWeek(start time.Time) *MealPlanner {
	mealPlanner := MealPlanner{start: start, plannedFor: 7, end: start.Add(7 * 24 * time.Hour)}
	mealPlanner.daysPlanned = make(map[time.Time]MealPlannerDay)
	return &mealPlanner
}

func (mp *MealPlanner) AddBreakfast(day time.Time, recipe Recipe) error {
	err := assertDayIsInTheMealPlannerTimeRange(day, mp)
	if err != nil {
		return err
	}
	if dayPlanned, dayAlreadyInPlanner := mp.daysPlanned[day]; dayAlreadyInPlanner {
		dayPlanned.Breakfast = recipe
		mp.daysPlanned[day] = dayPlanned
	} else {
		newDayPlanned := MealPlannerDay{Breakfast: recipe}
		mp.daysPlanned[day] = newDayPlanned
	}
	return nil
}

func (mp *MealPlanner) AddMorningSnack(day time.Time, recipe Recipe) error {
	err := assertDayIsInTheMealPlannerTimeRange(day, mp)
	if err != nil {
		return err
	}
	if dayPlanned, dayAlreadyInPlanner := mp.daysPlanned[day]; dayAlreadyInPlanner {
		dayPlanned.MorningSnack = recipe
		mp.daysPlanned[day] = dayPlanned
	} else {
		newDayPlanned := MealPlannerDay{MorningSnack: recipe}
		mp.daysPlanned[day] = newDayPlanned
	}
	return nil
}

func (mp *MealPlanner) AddLunch(day time.Time, recipe Recipe) error {
	err := assertDayIsInTheMealPlannerTimeRange(day, mp)
	if err != nil {
		return err
	}
	if dayPlanned, dayAlreadyInPlanner := mp.daysPlanned[day]; dayAlreadyInPlanner {
		dayPlanned.Lunch = recipe
		mp.daysPlanned[day] = dayPlanned
	} else {
		newDayPlanned := MealPlannerDay{Lunch: recipe}
		mp.daysPlanned[day] = newDayPlanned
	}
	return nil
}

func (mp *MealPlanner) AddAfternoonSnack(day time.Time, recipe Recipe) error {
	err := assertDayIsInTheMealPlannerTimeRange(day, mp)
	if err != nil {
		return err
	}
	if dayPlanned, dayAlreadyInPlanner := mp.daysPlanned[day]; dayAlreadyInPlanner {
		dayPlanned.AfternoonSnack = recipe
		mp.daysPlanned[day] = dayPlanned
	} else {
		newDayPlanned := MealPlannerDay{AfternoonSnack: recipe}
		mp.daysPlanned[day] = newDayPlanned
	}
	return nil
}

func (mp *MealPlanner) AddDinner(day time.Time, recipe Recipe) error {
	err := assertDayIsInTheMealPlannerTimeRange(day, mp)
	if err != nil {
		return err
	}
	if dayPlanned, dayAlreadyInPlanner := mp.daysPlanned[day]; dayAlreadyInPlanner {
		dayPlanned.Dinner = recipe
		mp.daysPlanned[day] = dayPlanned
	} else {
		newDayPlanned := MealPlannerDay{Dinner: recipe}
		mp.daysPlanned[day] = newDayPlanned
	}
	return nil
}
func (mp MealPlanner) Plan() MealPlan {
	plan := NewMealPlan()
	groceryList := NewGroceryList()

	for day, dayPlanned := range mp.daysPlanned {
		plan.NutritionalCompositionPerDay[day] = dayPlanned.getNutritionalComposition()
		for _, recipe := range dayPlanned.getRecipes() {
			for ingredient, quantity := range recipe.ingredients {
				groceryList.AddGrocery(ingredient, quantity)
			}
		}
	}
	plan.GroceryList = *groceryList
	return *plan
}

func (day MealPlannerDay) getNutritionalComposition() NutritionalComposition {
	return SumNutritionalCompositions(day.Breakfast.GetNutritionalComposition(), day.MorningSnack.GetNutritionalComposition(), day.Lunch.GetNutritionalComposition(), day.AfternoonSnack.GetNutritionalComposition(), day.Dinner.GetNutritionalComposition())
}
func (day MealPlannerDay) getRecipes() []Recipe {
	return []Recipe{day.Breakfast, day.MorningSnack, day.Lunch, day.AfternoonSnack, day.Dinner}
}

func assertDayIsInTheMealPlannerTimeRange(day time.Time, mp *MealPlanner) error {
	if day.Before(mp.start) || day.After(mp.end) {
		return errors.New("cannot add a recipe because the date given is not in the right time range")
	}
	return nil
}
