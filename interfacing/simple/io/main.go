package main

import (
	"fmt"
)

type Reader interface {
	Read()
}

type Writer interface {
	Write(data string)
}

type ReadWrite interface {
	Reader
	Writer
}

type printer struct{
	data string
}

func (p *printer) Read(){
	fmt.Println(p.data)
}

func (p *printer) Write(data string){
	p.data = data
}

func newReadWrite() ReadWrite{
	return &printer{}
}

func main(){
	rw := newReadWrite()

	rw.Write("Hello")
	rw.Read()
	callerInterface(rw)
	rw.Read()
	
}

func callerInterface(w interface{}){
	w.Write("somo")
}