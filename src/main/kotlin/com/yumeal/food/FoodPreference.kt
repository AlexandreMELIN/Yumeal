package com.yumeal.food

// Food preferences
data class FoodPreference(
    val proteins: List<Food>,
    val carbs: List<Food>,
    val dairy: List<Food>,
    val fruits: List<Food>,
    val vegetables: List<Food>,
    val plants: List<Food>
){
    init {
            require(proteins.all { it.category == FoodCategory.Protein}){"Not all proteins are tagged with category protein"}
            require(carbs.all { it.category == FoodCategory.Carbs}){"Not all carbs are tagged with category carbs"}
            require(dairy.all { it.category == FoodCategory.Dairy}){"Not all dairy are tagged with category dairy"}
            require(fruits.all { it.category == FoodCategory.Fruit}){"Not all fruit are tagged with category fruit"}
            require(vegetables.all { it.category == FoodCategory.Vegetable}){"Not all vegetables are tagged with category vegetables"}
            require(plants.all { it.category == FoodCategory.Plant}){"Not all plant are tagged with category plant"}
            require(proteins.isNotEmpty())
            require(carbs.isNotEmpty())
    }
}