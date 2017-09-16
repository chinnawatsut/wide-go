package main

import "testing"
import "strconv"

func TestCaseWithoutFizz(t *testing.T) {
	testSuit := []int{1, 2, 4}
	for _, v := range testSuit {
		expected := strconv.Itoa(v)
		result := fizzbuzz(v)
		if expected != result {
			errorExpeted(result, expected, t)
		}
	}
}
func TestCase3(t *testing.T) {
	result := fizzbuzz(3)
	expecting := "fizz"
	errorExpeted(result, expecting, t)
}

func errorExpeted(result, expeted string, t *testing.T) {
	if result != expeted {
		t.Error("result should be " + expeted + "but got " + result)

	}
}
