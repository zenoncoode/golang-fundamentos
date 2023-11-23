package entity

import "testing"

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}

	if order.Validate() == nil {
		t.Error("ID is required")
	}
}
