package com.yumeal.domain.algorithm

import com.yumeal.domain.food.PositiveQuantity

data class UserProfile(val targetBodyweightInKilogram: PositiveQuantity, val targetCalories: PositiveQuantity, val proteinPerBodyweightKilogram: PositiveQuantity, val carbsPerBodyweightKilogram: PositiveQuantity, val fatPerBodyweightKilogram: PositiveQuantity){
    val targetProtein = targetBodyweightInKilogram * proteinPerBodyweightKilogram
    val targetCarbs = targetBodyweightInKilogram * carbsPerBodyweightKilogram
}

