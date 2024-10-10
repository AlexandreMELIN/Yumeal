package com.yumeal.food

data class Meal(val ingredients: Map<Food, PositiveQuantity>){
    override fun toString(): String {
        return """
            ${
                ingredients.map { 
                    "${it.key.name}: ${it.value}"
                }.joinToString()
            }
        """.trimIndent()
    }
}
