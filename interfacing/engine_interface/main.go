package main

type Speeder interface {
	Speed() int
}

type Engine interface {
	Speeder
	Accel()
	Decel()
}

type engine struct {
	speed int
}

func (e *engine) Speed() int {
	return e.speed
}

func (e *engine) Accel() {
	e.speed = e.speed + 10
}

func (e *engine) Decel() {
	e.speed = e.speed - 10
}

func NewEngine() Engine {
	return &engine{}
}

func main() {

}

/*
engine interface of Engine
				and Engine interfaceof Speeder too.

engine have 3 behaviours speed,accel,decel

*/
