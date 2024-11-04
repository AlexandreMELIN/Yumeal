package com.yumeal.algorithm

import com.yumeal.domain.algorithm.MealPlanner
import com.yumeal.domain.food.Meal
import com.yumeal.domain.food.PositiveQuantity
import org.junit.jupiter.api.Test

import org.junit.jupiter.api.Assertions.*
import kotlin.math.exp

class MealPlannerTest {
    var mealPlanner = MealPlanner()
    @Test
    fun calculateQuantityFor100gram() {
        val expected = PositiveQuantity(22.5)
        val proteinPer100g = PositiveQuantity(15)
        val quantity = PositiveQuantity(150)
        assertEquals(expected.quantity, mealPlanner.calculateQuantityFor100gram(quantity, proteinPer100g).quantity)
    }

    @Test
    fun calculateRatioInGram() {
        val numerator = PositiveQuantity(1)
        val denominator = PositiveQuantity(3)
        val expected = PositiveQuantity(33)
        assertEquals(expected.quantity, mealPlanner.calculateRatioInGram(numerator, denominator).quantity)
    }
}