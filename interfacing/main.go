package main

import "fmt"

type animal interface {
	run()
	eat()
	walk()
}

type dog struct {
	name string
}

func (d dog) run() {
	fmt.Println("RUN")

}

func (d dog) eat() {
	fmt.Println("EAT")
}

func (d dog) walk() {
	fmt.Println("WALK")

}

func allTheyCanDo(a animal) {
	a.eat()
	a.run()
	a.walk()
}

func main() {
	rogers := dog{"roger"}
	allTheyCanDo(rogers)
}
