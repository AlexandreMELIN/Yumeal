package com.yumeal.algorithm

import com.yumeal.food.Meal

data class MealsForADay(val breakfast: Meal, val lunch: Meal, val dinner: Meal){
    override fun toString(): String {
        return """
            Breakfast: ${breakfast.toString()}
            Lunch: ${lunch.toString()}
            Dinner: ${dinner.toString()}
        """.trimIndent()
    }
}
