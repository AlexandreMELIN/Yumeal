package com.yumeal.domain.algorithm

import com.yumeal.domain.food.*
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test
import kotlin.test.assertFailsWith

class MealOrganizationTest{
    @Test
    fun constructorMustCheckThatRatioSumIsEqualTo1(){
        val dummyProtein = Food(
            name = "Dummy food",
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
        val dummyCarb = Food(
            name = "Dummy food",
            category = FoodCategory.Carbs,
            nutritionalPanel = NutritionalPanel(
                protein = PositiveQuantity(13.0),
                carbs = PositiveQuantity(1.0),
                sugar = PositiveQuantity(0.0),
                fat = PositiveQuantity(11.0),
                saturatedFat = PositiveQuantity(3.0),
                fiber = PositiveQuantity(0.0)
            )
        )
        val dummyDairy = Food(
            name = "Dummy food",
            category = FoodCategory.Dairy,
            nutritionalPanel = NutritionalPanel(
                protein = PositiveQuantity(13.0),
                carbs = PositiveQuantity(1.0),
                sugar = PositiveQuantity(0.0),
                fat = PositiveQuantity(11.0),
                saturatedFat = PositiveQuantity(3.0),
                fiber = PositiveQuantity(0.0)
            )
        )
        val foodPreference = FoodPreference(
            proteins = listOf(dummyProtein),
            carbs = listOf(dummyCarb),
            dairy = listOf(dummyDairy),
            fruits = listOf(),
            vegetables = listOf(),
            plants = listOf()
        )
        assertFailsWith<IllegalArgumentException> {
            MealOrganization(listOf(
                MealPreference(Ratio(0.7), foodPreference),
                MealPreference(Ratio(0.7), foodPreference)))
        }
        assertFailsWith<IllegalArgumentException> {
            MealOrganization(listOf(
                MealPreference(Ratio(0.4), foodPreference),
                MealPreference(Ratio(0.7), foodPreference)))
        }
        val mealOrganization = MealOrganization(listOf(
            MealPreference(Ratio(0.7), foodPreference),
            MealPreference(Ratio(0.3), foodPreference)
        ))
        assertNotNull(mealOrganization)
    }
}