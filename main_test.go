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
func TestCaseFizz(t *testing.T) {
	testSuit := []int{3, 6}
	for _, v := range testSuit {
		expected := "fizz"
		result := fizzbuzz(v)
		if result != expected {
			errorExpeted(result, expected, t)
		}
	}
}

func TestCase5(t *testing.T) {
	errorExpeted(fizzbuzz(5), "buzz", t)
}

func errorExpeted(result, expeted string, t *testing.T) {
	if result != expeted {
		t.Error("result should be " + expeted + " but got " + result)

	}
}
