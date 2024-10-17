package com.yumeal.grocery

import com.yumeal.algorithm.MealsForADay
import com.yumeal.food.Food
import com.yumeal.food.Meal
import com.yumeal.food.PositiveQuantity

class GroceryList(mealsPlanned: List<Meal>) {
    constructor(mealsPlanned: MealsForADay) : this(listOf(mealsPlanned.breakfast, mealsPlanned.lunch, mealsPlanned.dinner))

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
