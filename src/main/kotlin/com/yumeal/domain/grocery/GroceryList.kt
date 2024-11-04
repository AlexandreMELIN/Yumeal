package com.yumeal.domain.grocery

import com.yumeal.domain.algorithm.Classic3MealsForADay
import com.yumeal.domain.food.Food
import com.yumeal.domain.food.Meal
import com.yumeal.domain.food.PositiveQuantity

class GroceryList(mealsPlanned: List<Meal>) {
    constructor(mealsPlanned: Classic3MealsForADay) : this(listOf(mealsPlanned.breakfast, mealsPlanned.lunch, mealsPlanned.dinner))
    private val groceries: MutableMap<Food, PositiveQuantity>  = mutableMapOf()

    init {
        for (meal in mealsPlanned) {
           for (ingredient in meal.ingredients.keys){
               if (groceries.containsKey(ingredient)){
                   groceries[ingredient] = groceries[ingredient]!! + meal.ingredients[ingredient]!!
               }else{
                   groceries[ingredient] = meal.ingredients[ingredient]!!
               }
           }
        }
    }

    fun getGroceries(): Map<Food, PositiveQuantity> {
        return groceries
    }

    override fun toString(): String {
        return """
            ${
                groceries.map {
                    "${it.key.name}: ${it.value}g"
                }.joinToString()
        }         
        """.trimIndent()
    }

}
