package com.yumeal.food

data class NutritionalPanel (val protein: PositiveQuantity, val carbs: PositiveQuantity, val sugar: PositiveQuantity, val fat: PositiveQuantity, val saturatedFat: PositiveQuantity, val fiber: PositiveQuantity){
    val calories: Double = 4 * protein.quantity + 4 * carbs.quantity + 4 * fiber.quantity + 9 * fat.quantity


    constructor(): this(PositiveQuantity(0), PositiveQuantity(0),PositiveQuantity(0),PositiveQuantity(0),PositiveQuantity(0),PositiveQuantity(0))
    operator fun plus(other: NutritionalPanel): NutritionalPanel {
        return NutritionalPanel(
            protein + other.protein,
            carbs + other.carbs,
            sugar + other.sugar,
            fat + other.fat,
            saturatedFat + other.saturatedFat,
            fiber + other.fiber)
    }

}