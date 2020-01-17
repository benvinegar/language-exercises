package numbers

import "testing"

func TestValidateCreditCard(t *testing.T) {
	var result bool

	result = ValidateCreditCard("4111111111111111") // visa
	if result != true {
		t.Errorf("Expected %t, got %t", result, true)
	}

	result = ValidateCreditCard("3566002020360505") // jcb
	if result != true {
		t.Errorf("Expected %t, got %t", result, true)
	}

	result = ValidateCreditCard("5555555555554444") // mastercard
	if result != true {
		t.Errorf("Expected %t, got %t", result, true)
	}

	result = ValidateCreditCard("30569309025904") // diner's club
	if result != true {
		t.Errorf("Expected %t, got %t", result, true)
	}

	result = ValidateCreditCard("4111111111111112") // bad
	if result != false {
		t.Errorf("Expected %t, got %t", result, false)
	}
}
