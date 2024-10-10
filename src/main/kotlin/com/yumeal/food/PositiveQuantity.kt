package com.yumeal.food

class PositiveQuantity(val quantity: Double) {

    constructor(quantity: Int): this(quantity.toDouble())

    operator fun times(d: Double): PositiveQuantity {
        require(d >= 0) { "Cannot multiply a positive quantity by a negative double" }
        return PositiveQuantity(quantity * d)
    }

    operator fun times(p: PositiveQuantity): PositiveQuantity{
        return PositiveQuantity(quantity * p.quantity)
    }

    override fun toString(): String {
        return quantity.toString()
    }

    init {
        require(quantity >= 0){ "Quantity must be greater than 0"}
    }



}
