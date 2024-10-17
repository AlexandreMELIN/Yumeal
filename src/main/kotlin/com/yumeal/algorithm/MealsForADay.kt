package com.yumeal.algorithm

import com.yumeal.food.Meal
import com.yumeal.food.NutritionalPanel

data class MealsForADay(val breakfast: Meal, val lunch: Meal, val dinner: Meal){
    override fun toString(): String {
        return """
            Breakfast: ${breakfast.toString()}
            Lunch: ${lunch.toString()}
            Dinner: ${dinner.toString()}
            Nutritional panel for the day : ${getNutritionalPanel()}
        """.trimIndent()
    }

    fun getNutritionalPanel(): NutritionalPanel{
        return breakfast.getNutritionalPanel() + lunch.getNutritionalPanel() + dinner.getNutritionalPanel()
    }
}
