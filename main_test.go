package main

import "testing"

func TestCase1(t *testing.T) {
	result := fizzbuzz(1)
	expecting := "1"

	if result != expecting {
		t.Error("result should be " + expecting + " but got " + result)
	}
}
