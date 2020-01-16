package numbers

import (
	"reflect"
	"testing"
)

func TestPrimes(t *testing.T) {
	var result []int
	var expected []int

	result = Primes(1)
	expected = []int{2}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	result = Primes(2)
	expected = []int{2, 3}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	result = Primes(10)
	expected = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
