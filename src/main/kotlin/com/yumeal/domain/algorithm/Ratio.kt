package com.yumeal.domain.algorithm

class Ratio(val value: Double) {
    constructor(value: Int): this(value.toDouble())
    init {
        require(value in 0.0..1.0){"Ratio must be between 0.0 and 1.0"}
    }
}