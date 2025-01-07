package internal

import "testing"

func TestMealPlanner(t *testing.T) {
	// Initialize food items with their nutritional values (per 100g)
	var mealPreference = MealPreference{
		Protein: []Food{
			{
				Name: "Eggs",
				NutritionalPanel: NutritionalPanel{
					Protein: 13,
					Carbs:   1,
					Fat:     11,
					Fiber:   0,
				},
				Category: Protein,
			},
			{
				Name: "Chicken",
				NutritionalPanel: NutritionalPanel{
					Protein: 31,
					Carbs:   0,
					Fat:     3,
					Fiber:   0,
				},
				Category: Protein,
			},
			{
				Name: "Turkey",
				NutritionalPanel: NutritionalPanel{
					Protein: 29,
					Carbs:   0,
					Fat:     2,
					Fiber:   0,
				},
				Category: Protein,
			},
			{
				Name: "Salmon",
				NutritionalPanel: NutritionalPanel{
					Protein: 25,
					Carbs:   0,
					Fat:     13,
					Fiber:   0,
				},
				Category: Protein,
			},
			{
				Name: "Tofu",
				NutritionalPanel: NutritionalPanel{
					Protein: 8,
					Carbs:   2,
					Fat:     4,
					Fiber:   1,
				},
				Category: Protein,
			},
		},
		Plant: []Food{
			{
				Name: "Lentils",
				NutritionalPanel: NutritionalPanel{
					Protein: 9,
					Carbs:   20,
					Fat:     1,
					Fiber:   8,
				},
				Category: Plant,
			},
			{
				Name: "Chickpea",
				NutritionalPanel: NutritionalPanel{
					Protein: 9,
					Carbs:   27,
					Fat:     3,
					Fiber:   8,
				},
				Category: Plant,
			},
		},
		Vegetables: []Food{
			{
				Name: "Zucchini",
				NutritionalPanel: NutritionalPanel{
					Protein: 1,
					Carbs:   3,
					Fat:     0,
					Fiber:   1,
				},
				Category: Vegetables,
			},
			{
				Name: "Pepper",
				NutritionalPanel: NutritionalPanel{
					Protein: 1,
					Carbs:   6,
					Fat:     0,
					Fiber:   2,
				},
				Category: Vegetables,
			},
			{
				Name: "Green Beans",
				NutritionalPanel: NutritionalPanel{
					Protein: 2,
					Carbs:   7,
					Fat:     0,
					Fiber:   3,
				},
				Category: Vegetables,
			},
			{
				Name: "Leek",
				NutritionalPanel: NutritionalPanel{
					Protein: 1,
					Carbs:   14,
					Fat:     0,
					Fiber:   2,
				},
				Category: Vegetables,
			},
		},
		Starch: []Food{
			{
				Name: "Whole Pasta",
				NutritionalPanel: NutritionalPanel{
					Protein: 13,
					Carbs:   71,
					Fat:     2,
					Fiber:   7,
				},
				Category: Starch,
			},
			{
				Name: "Whole Rice",
				NutritionalPanel: NutritionalPanel{
					Protein: 7,
					Carbs:   76,
					Fat:     1,
					Fiber:   4,
				},
				Category: Starch,
			},
			{
				Name: "Potatoes",
				NutritionalPanel: NutritionalPanel{
					Protein: 2,
					Carbs:   17,
					Fat:     0,
					Fiber:   2,
				},
				Category: Starch,
			},
		},
		Dairy: []Food{
			{
				Name: "Greek Yogurt",
				NutritionalPanel: NutritionalPanel{
					Protein: 10,
					Carbs:   4,
					Fat:     0,
					Fiber:   0,
				},
				Category: Dairy,
			},
			{
				Name: "Comte Cheese",
				NutritionalPanel: NutritionalPanel{
					Protein: 29,
					Carbs:   0,
					Fat:     32,
					Fiber:   0,
				},
				Category: Dairy,
			},
		},
		Fruits: []Food{
			{
				Name: "Orange",
				NutritionalPanel: NutritionalPanel{
					Protein: 1,
					Carbs:   12,
					Fat:     0,
					Fiber:   2,
				},
				Category: Fruits,
			},
			{
				Name: "Banana",
				NutritionalPanel: NutritionalPanel{
					Protein: 1,
					Carbs:   23,
					Fat:     0,
					Fiber:   3,
				},
				Category: Fruits,
			},
			{
				Name: "Apple",
				NutritionalPanel: NutritionalPanel{
					Protein: 0,
					Carbs:   14,
					Fat:     0,
					Fiber:   2,
				},
				Category: Fruits,
			},
		},
	}
	target := NutritionalPanel{Protein: 55, Carbs: 132}
	meal := PlanAMeal(mealPreference, target)
	expectedProtein := 55
	expectedCarbs := 132
	gotProtein := meal.GetProtein()
	gotCarbs := meal.GetCarbs()
	if gotProtein < expectedProtein {
		t.Errorf("got protein <= %d, want %d", gotProtein, expectedProtein)
	}
	if gotCarbs < expectedCarbs {
		t.Errorf("got carbs <= %d, want %d", gotCarbs, expectedCarbs)
	}
}
