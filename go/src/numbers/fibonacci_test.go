package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var result []int
	var expected []int

	result = Fibonacci(2)
	expected = []int{0, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	result = Fibonacci(3)
	expected = []int{0, 1, 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	result = Fibonacci(10)
	expected = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
