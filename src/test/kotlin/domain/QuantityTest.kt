package domain

import org.junit.jupiter.api.assertDoesNotThrow
import org.junit.jupiter.api.assertThrows
import kotlin.test.Test

class QuantityTest{
    @Test
    fun testProportionMustBePositiveOrZero(){
        assertThrows<IllegalArgumentException> { Quantity(proportion = -1.0, unit = MeasureUnit.GRAM) }
        assertDoesNotThrow { Quantity(proportion = 0.0, unit = MeasureUnit.GRAM) }
        assertDoesNotThrow { Quantity(proportion = 0.1, unit = MeasureUnit.GRAM) }
        assertDoesNotThrow { Quantity(proportion = 1.1, unit = MeasureUnit.GRAM) }
    }
}