package main

import "fmt"

type vanko struct {
	count int
	name string
}

func (vanko)say() {
	fmt.Println("Hi")
}

type marine struct {

}

func (marine)say() {
	fmt.Println("Hola")
}

func main() {
	//v := vanko{2,"vanko"}
	//vdd(v)
	m := marine{}
	vdd(m)
}

func vdd(v interface{}){
	v.(marine).say()
	fmt.Printf("%v", v)
}
