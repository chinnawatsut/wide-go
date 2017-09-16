package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, Wide-Go! ")
}

func fizzbuzz(n int) string {
	if n%3 == 0 {
		return "fizz"
	}

	if n == 5 {
		return "buzz"
	}
	return strconv.Itoa(n)
}
