package main

import (
	"testing"
)

func expectedVal(t *testing.T, expected int, meter meter) {
	if meter.val() != expected {
		t.Fatal("expected", expected, "got", meter.val())
	}
}
func TestDefaultValue(t *testing.T) {
	meter := NewMeter()
	expectedVal(t, 0, meter)
}

func TestIncMeter(t *testing.T) {
	meter := NewMeter()
	meter.inc()
	expectedVal(t, 1, meter)
}

func TestDecMeter(t *testing.T) {
	meter := NewMeter()
	meter.dec()
	expectedVal(t, -1, meter)
}
