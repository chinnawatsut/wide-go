package main

import "testing"
import "strconv"

func TestCaseWithNormal(t *testing.T) {
	testSuit := []int{1, 2, 4, 8, 11, 13, 14}
	for _, v := range testSuit {
		expected := strconv.Itoa(v)
		result := fizzbuzz(v)

		errorExpeted(result, expected, t)
	}
}
func TestCaseModBy3(t *testing.T) {
	testSuit := []int{3, 6, 9, 12}
	for _, v := range testSuit {
		expected := "fizz"
		result := fizzbuzz(v)

		errorExpeted(result, expected, t)

	}
}

func TestCaseModBy5(t *testing.T) {
	testSuit := []int{5, 10}
	for _, v := range testSuit {
		result := fizzbuzz(v)
		expected := "buzz"

		errorExpeted(result, expected, t)
	}
}

func TestCase15(t *testing.T) {
	errorExpeted(fizzbuzz(15), "fizzbuzz", t)
}

func errorExpeted(result, expeted string, t *testing.T) {
	if result != expeted {
		t.Error("result should be " + expeted + " but got " + result)

	}
}
