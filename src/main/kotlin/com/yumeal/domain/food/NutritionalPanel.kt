package com.yumeal.domain.food

data class NutritionalPanel (val protein: PositiveQuantity, val carbs: PositiveQuantity, val sugar: PositiveQuantity, val fat: PositiveQuantity, val saturatedFat: PositiveQuantity, val fiber: PositiveQuantity){
    val calories: Double = 4 * protein.quantity + 4 * carbs.quantity + 2 * fiber.quantity + 9 * fat.quantity


    constructor(): this(
        PositiveQuantity(0), PositiveQuantity(0),
        PositiveQuantity(0),
        PositiveQuantity(0),
        PositiveQuantity(0),
        PositiveQuantity(0)
    )
    operator fun plus(other: NutritionalPanel): NutritionalPanel {
        return NutritionalPanel(
            protein + other.protein,
            carbs + other.carbs,
            sugar + other.sugar,
            fat + other.fat,
            saturatedFat + other.saturatedFat,
            fiber + other.fiber)
    }
    operator fun times(ratio: PositiveQuantity): NutritionalPanel {
        return NutritionalPanel(
            protein * ratio,
            carbs * ratio,
            sugar * ratio,
            fat * ratio,
            saturatedFat * ratio,
            fiber * ratio
        )
    }

    override fun toString(): String {
        return """
            protein=$protein,
            carbs=$carbs,
                sugar=$sugar,
            fat=$fat,
                saturatedFat=$saturatedFat,
            fiber=$fiber,
            calories=$calories)
        """.trimIndent()
    }


}