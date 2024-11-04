package com.yumeal.domain.food

data class Meal(val ingredients: Map<Food, PositiveQuantity>){
    override fun toString(): String {
        return """
            ${
                ingredients.map { 
                    "${it.key.name}: ${it.value}" 
                }.joinToString()
            }
            Nutritional panel : ${getNutritionalPanel()}
        """.trimIndent()
    }

    fun getNutritionalPanel(): NutritionalPanel {
        var result = NutritionalPanel()
        for (food in ingredients.keys){
            var foodNutritionalPanel = food.nutritionalPanel
            var quantity = ingredients[food]
            assert(quantity != null)
            var ratio = quantity?.quantity?.div(100)
            assert(ratio != null)
            result += foodNutritionalPanel * PositiveQuantity(ratio!!)
        }
        return result
    }
}
