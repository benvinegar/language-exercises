package numbers

import (
	"strconv"
)

// ValidateCreditCard returns true if the given credit card number is valid
// using Luhn's algorithm (https://en.wikipedia.org/wiki/Luhn_algorithm)
func ValidateCreditCard(num string) bool {

	var doubled [16]int

	for i := len(num) - 1; i >= 0; i-- {
		// ignore error; assume string is well-formed and without errors (e.g. missing digits)
		digit, _ := strconv.Atoi(string(num[i]))
		if i%2 == 0 {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}
		doubled[i] = digit
	}

	// sum digits
	sum := 0
	for i := 0; i < len(doubled); i++ {
		sum += doubled[i]
	}

	return sum%10 == 0
}
