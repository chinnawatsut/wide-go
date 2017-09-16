package main

import "testing"

func TestCase1(t *testing.T) {
	result := fizzbuzz(1)
	expecting := "1"

	errorExpeted(result, expecting, t)
}

func TestCase2(t *testing.T) {
	result := fizzbuzz(2)
	expecting := "2"
	errorExpeted(result, expecting, t)
}

func errorExpeted(result, expeted string, t *testing.T) {
	if result != expeted {
		t.Error("result should be " + expeted + "but got " + result)

	}
}
