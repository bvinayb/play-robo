package direction

import "testing"

func TestIsValidDirection(t *testing.T) {
	if !IsValidDirection("NORTH") {
		t.Errorf("This is a valid direction")
	}
}

func TestIsNotValidDirection(t *testing.T) {
	if IsValidDirection("MELBOURNE") {
		t.Errorf("This is not a valid direction")
	}
}
