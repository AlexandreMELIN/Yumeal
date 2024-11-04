package com.yumeal.domain.algorithm

import com.yumeal.domain.food.Food
import com.yumeal.domain.food.FoodPreference
import com.yumeal.domain.food.Meal
import com.yumeal.domain.food.PositiveQuantity
import kotlin.math.floor
import kotlin.random.Random

class MealPlanner {
    private val VEGETABLE_PORTION_IN_GRAM = 80
    fun plan(userProfile: UserProfile, breakfastPreference: FoodPreference, mealPreference: FoodPreference): Classic3MealsForADay {
        val targetFiber = (14 * (userProfile.targetCalories.quantity / 1000)).toInt()
        var breakfast = planAMeal(
            targetCalories = userProfile.targetCalories * 0.2,
            targetProteins = userProfile.targetProtein * 0.2,
            targetCarbs = userProfile.targetCarbs * 0.2,
            preference = breakfastPreference,
        )
        var lunch = planAMeal(
            targetCalories = userProfile.targetCalories * 0.3,
            targetProteins = userProfile.targetProtein * 0.3,
            targetCarbs = userProfile.targetCarbs * 0.3,
            preference = mealPreference,
        )
        var dinner = planAMeal(
            targetCalories = userProfile.targetCalories * 0.5,
            targetProteins = userProfile.targetProtein * 0.5,
            targetCarbs = userProfile.targetCarbs * 0.5,
            preference = mealPreference,
        )

        return Classic3MealsForADay(breakfast, lunch, dinner)
    }

    fun plan(userProfile: UserProfile, mealOrganization: MealOrganization): MealsForADay{
        var meals: MealsForADay = MealsForADay(mealOrganization.mealPreference.map {
            planAMeal(
                targetCalories = userProfile.targetCalories * it.ratio.value,
                targetProteins = userProfile.targetProtein * it.ratio.value,
                targetCarbs = userProfile.targetCarbs * it.ratio.value,
                preference = it.foodPreference)
        })
        return meals
    }

    private fun planAMeal(targetCalories: PositiveQuantity, targetProteins: PositiveQuantity, targetCarbs: PositiveQuantity, preference: FoodPreference): Meal {
        val result = mutableMapOf<Food, PositiveQuantity>()
        var remainingProteins = targetProteins.quantity
        var remainingCarbs = targetCarbs.quantity
        var remainingCalories = targetCalories.quantity * 0.9 // Leaving 10 % of calories for seasoning / fat during the cooking process
        val shouldIncludePlants = preference.plants.isNotEmpty() && Random.nextBoolean()
        var carbs = preference.carbs.random()

        // We first include plants if needed. We substract the amount of carbs, calories, and proteins they bring
        if (shouldIncludePlants){
            var plant = preference.plants.random()
            var plantQuantity = calculateRatioInGram(PositiveQuantity(remainingCarbs / 2), plant.nutritionalPanel.carbs)
            remainingCarbs -= calculateQuantityFor100gram(plant.nutritionalPanel.carbs, plantQuantity).quantity
            remainingProteins -= calculateQuantityFor100gram(plant.nutritionalPanel.protein, plantQuantity).quantity
            remainingCalories -= calculateQuantityFor100gram(PositiveQuantity(plant.nutritionalPanel.calories), plantQuantity).quantity
            result[plant] = plantQuantity
        }
        // Calculating either on the full carbs or carbs left after adding plant
        result[carbs] = calculateRatioInGram(PositiveQuantity(remainingCarbs), carbs.nutritionalPanel.carbs)
        remainingCalories -= calculateQuantityFor100gram(PositiveQuantity(carbs.nutritionalPanel.calories), PositiveQuantity(result[carbs]?.quantity ?: 0.0)).quantity

        // Computing protein and substracting calories, starting with dairy
        if (preference.dairy.isNotEmpty()){
            var dairy = preference.dairy.random()
            // TODO: those ratios need to be in the preference as for example it's ok for the breakfast to have more calories coming from dairy (like white cheese)
            val targetProteinsForDairy = remainingProteins * 0.2
            result[dairy] = calculateRatioInGram(PositiveQuantity(targetProteinsForDairy), dairy.nutritionalPanel.protein)
            remainingProteins -= targetProteinsForDairy
            remainingCalories -= calculateQuantityFor100gram(PositiveQuantity(dairy.nutritionalPanel.calories), PositiveQuantity(result[dairy]?.quantity ?: 0.0)).quantity
        }
        var protein = preference.proteins.random()
        result[protein] = calculateRatioInGram(PositiveQuantity(remainingProteins), protein.nutritionalPanel.protein)
        remainingCalories -= calculateQuantityFor100gram(PositiveQuantity(protein.nutritionalPanel.calories), PositiveQuantity(result[protein]?.quantity ?: 0.0)).quantity
        //Ugly but ok for now
        if (remainingCalories < 0)
            return Meal(result)
        // Adding veggies and fruits on the leftover calories. If none are set, will be in a calorie deficit but a diet without fruit and  vegetable is not healthy anyway
        var hasVegetables = preference.vegetables.isNotEmpty()
        var hasFruits = preference.fruits.isNotEmpty()
        when {
            hasVegetables && hasFruits -> {
                val fruit = preference.fruits.random()
                val vegetable = preference.vegetables.random()
                result[vegetable] = PositiveQuantity(VEGETABLE_PORTION_IN_GRAM)
                result[fruit] = PositiveQuantity(VEGETABLE_PORTION_IN_GRAM)
            }
            hasVegetables -> {
                val vegetable = preference.vegetables.random()
                result[vegetable] = PositiveQuantity(2 * VEGETABLE_PORTION_IN_GRAM)
            }
            hasFruits -> {
                val fruit = preference.fruits.random()
                result[fruit] = PositiveQuantity(2 * VEGETABLE_PORTION_IN_GRAM)

            }

            else -> {
                // default we do nothing
            }
        }


        return Meal(result)
    }
    fun calculateQuantityFor100gram(referentialQuantity: PositiveQuantity, quantity: PositiveQuantity): PositiveQuantity {
        return PositiveQuantity((referentialQuantity.quantity / 100) * quantity.quantity)
    }
    fun calculateRatioInGram(numerator: PositiveQuantity, denominator: PositiveQuantity): PositiveQuantity = PositiveQuantity(
        floor(numerator.quantity / denominator.quantity * 100)
    )
}