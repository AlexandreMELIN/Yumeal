package com.yumeal.food

data class NutritionalPanel (val protein: PositiveQuantity, val carbs: PositiveQuantity, val sugar: PositiveQuantity, val fat: PositiveQuantity, val saturatedFat: PositiveQuantity, val fiber: PositiveQuantity){
    val calories: Double = 4 * protein.quantity + 4 * carbs.quantity + 4 * fiber.quantity + 9 * fat.quantity
}