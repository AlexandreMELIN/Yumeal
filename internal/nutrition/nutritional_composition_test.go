package nutrition

import "testing"

func TestConstructorComputesTheRightAmountOfCalories(t *testing.T) {
	got := NewNutritionalComposition(60, 10, 100).Calories
	want := 730.0
	if got != want {
		t.Errorf("Computed calories %f mistmach from the expected amount of calories: %f", got, want)
	}
}
