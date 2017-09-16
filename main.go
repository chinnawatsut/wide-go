package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, Wide-Go! ")
}

func fizzbuzz(n int) string {
	if n == 3 {
		return "fizz"
	}
	return strconv.Itoa(n)
}
