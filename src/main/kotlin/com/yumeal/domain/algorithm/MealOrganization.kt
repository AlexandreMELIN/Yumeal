package com.yumeal.domain.algorithm

class MealOrganization(val mealPreference: List<MealPreference>) {
    init {
        require(mealPreference.sumOf { it.ratio.value } == 1.0) {"Sum of all meal preference ratio must be equals to 1"}
    }
}