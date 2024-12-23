package com.yumeal.domain.meals

import com.yumeal.domain.food.Meal

data class MealsForADay(val meals: List<Meal>) {
    override fun toString(): String {
        return """
            ${
            meals.joinToString("\n") { it.toString() }
        }
        """.trimIndent()
    }
}
