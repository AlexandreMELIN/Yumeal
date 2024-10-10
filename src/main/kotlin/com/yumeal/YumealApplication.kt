package com.yumeal

import com.yumeal.algorithm.MealPlanner
import com.yumeal.algorithm.UserProfile
import com.yumeal.food.*

//@SpringBootApplication
class YumealApplication

fun main(args: Array<String>) {
// Ingredients and their nutritional panels using FoodQuantity
    val pasta = Food(
        name = "Whole Pasta",
        category = FoodCategory.Carbs,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(14.0),
            carbs = PositiveQuantity(73.0),
            sugar = PositiveQuantity(3.0),
            fat = PositiveQuantity(3.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(8.0)
        )
    )

    val rice = Food(
        name = "Rice",
        category = FoodCategory.Carbs,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(8.0),
            carbs = PositiveQuantity(76.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(3.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val oat = Food(
        name = "Oat",
        category = FoodCategory.Carbs,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(13.0),
            carbs = PositiveQuantity(70.0),
            sugar = PositiveQuantity(3.0),
            fat = PositiveQuantity(6.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(7.0)
        )
    )

    val chicken = Food(
        name = "Chicken",
        category = FoodCategory.Protein,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(27.0),
            carbs = PositiveQuantity(0.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(6.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val egg = Food(
        name = "Eggs",
        category = FoodCategory.Protein,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(13.0),
            carbs = PositiveQuantity(1.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(11.0),
            saturatedFat = PositiveQuantity(3.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val salmon = Food(
        name = "Salmon",
        category = FoodCategory.Protein,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(22.0),
            carbs = PositiveQuantity(0.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(5.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val shrimps = Food(
        name = "Shrimps",
        category = FoodCategory.Protein,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(24.0),
            carbs = PositiveQuantity(0.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(0.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val comteCheese = Food(
        name = "Comte Cheese",
        category = FoodCategory.Dairy,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(25.0),
            carbs = PositiveQuantity(0.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(35.0),
            saturatedFat = PositiveQuantity(21.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val mozarella = Food(
        name = "Mozarella",
        category = FoodCategory.Dairy,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(28.0),
            carbs = PositiveQuantity(3.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(17.0),
            saturatedFat = PositiveQuantity(11.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val whiteCheese = Food(
        name = "White Cheese",
        category = FoodCategory.Dairy,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(8.0),
            carbs = PositiveQuantity(3.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(3.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val pepper = Food(
        name = "Pepper",
        category = FoodCategory.Vegetable,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(1.0),
            carbs = PositiveQuantity(6.0),
            sugar = PositiveQuantity(4.0),
            fat = PositiveQuantity(0.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(2.0)
        )
    )

    val eggplant = Food(
        name = "Eggplant",
        category = FoodCategory.Vegetable,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(1.0),
            carbs = PositiveQuantity(6.0),
            sugar = PositiveQuantity(4.0),
            fat = PositiveQuantity(0.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(3.0)
        )
    )

    val zucchini = Food(
        name = "Zucchini",
        category = FoodCategory.Vegetable,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(1.0),
            carbs = PositiveQuantity(3.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(1.0),
            saturatedFat = PositiveQuantity(0.0),
            fiber = PositiveQuantity(0.0)
        )
    )

    val chickpea = Food(
        name = "Chickpea",
        category = FoodCategory.Plant,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(19.0),
            carbs = PositiveQuantity(61.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(6.0),
            saturatedFat = PositiveQuantity(1.0),
            fiber = PositiveQuantity(17.0)
        )
    )

    val lentil = Food(
        name = "Lentil",
        category = FoodCategory.Plant,
        nutritionalPanel = NutritionalPanel(
            protein = PositiveQuantity(25.0),
            carbs = PositiveQuantity(61.0),
            sugar = PositiveQuantity(0.0),
            fat = PositiveQuantity(1.0),
            saturatedFat = PositiveQuantity(1.0),
            fiber = PositiveQuantity(30.0)
        )
    )



    val breakfastPreference = FoodPreference(
        proteins = listOf(egg),
        carbs = listOf(oat),
        dairy = listOf(whiteCheese),
        fruits = listOf(),
        vegetables = listOf(),
        plants = listOf()
    )

    val mealPreference = FoodPreference(
        proteins = listOf(egg, salmon, shrimps, chicken),
        carbs = listOf(rice, pasta),
        dairy = listOf(comteCheese, mozarella),
        fruits = listOf(),
        vegetables = listOf(zucchini, eggplant, pepper),
        plants = listOf(lentil, chickpea)
    )
    var mealPlanner = MealPlanner()
    var userProfile = UserProfile(
        targetCalories = PositiveQuantity(2900),
        targetBodyweightInKilogram = PositiveQuantity(80),
        proteinPerBodyweightKilogram = PositiveQuantity(1.6),
        carbsPerBodyweightKilogram = PositiveQuantity(4),
        fatPerBodyweightKilogram = PositiveQuantity(0.75)
    )
    val mealForADay = mealPlanner.plan(userProfile, breakfastPreference, mealPreference)
    println(mealForADay.toString())
}

