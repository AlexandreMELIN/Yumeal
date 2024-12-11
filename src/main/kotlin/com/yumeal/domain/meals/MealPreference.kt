package com.yumeal.domain.meals

import com.yumeal.domain.algorithm.Ratio
import com.yumeal.domain.food.FoodPreference

data class MealPreference(val ratio: Ratio, val foodPreference: FoodPreference) {

}