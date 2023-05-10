package domain

data class Quantity(val proportion: Double, val unit: MeasureUnit){
    init {
        require(proportion >= 0){
            "The quantity proportion must be greater than or equal to 0"
        }
    }
}
