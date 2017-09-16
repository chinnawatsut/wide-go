package main

import (
	"testing"
)

func exceptSpeed(t *testing.T, speeder Speeder, speed int) {
	actual := speeder.Speed()
	if actual != speed {
		t.Fatal("expected", speed, "got", actual)
	}
}

func TestEngine(t *testing.T) {
	e := NewEngine()
	exceptSpeed(t, e, 0)
	e.Accel()
	exceptSpeed(t, e, 10)
	e.Decel()
	exceptSpeed(t, e, 0)

}
