package main

type Reader interface {
	Read(string)
}

type Writer interface {
	Write(src *string,data string)
}

func main(){

}