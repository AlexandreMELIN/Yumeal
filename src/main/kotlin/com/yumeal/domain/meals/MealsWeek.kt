package com.yumeal.domain.meals

import com.yumeal.domain.grocery.GroceryList

data class MealsWeek(
    val day1: Classic3MealsForADay,
    val day2: Classic3MealsForADay,
    val day3: Classic3MealsForADay,
    val day4: Classic3MealsForADay,
    val day5: Classic3MealsForADay,
    val day6: Classic3MealsForADay,
    val day7: Classic3MealsForADay,
){
    override fun toString(): String {
        return """
            Day1:
                $day1
            Day2:
                $day2
            Day3:
                $day3
            Day4:
                $day4
            Day5:
                $day5
            Day6:
                $day6
            Day7:
                $day7
        """.trimIndent()
    }

    fun getGroceries(): GroceryList{
        var allMeals = listOf(day1, day2, day3, day4, day5, day6, day7)
            .map { it.toMealsForADay().meals }
            .reduce{
                meals, current -> meals + current
            }
            .toList()
        return GroceryList(allMeals)
    }
}
