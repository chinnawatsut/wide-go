package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, Wide-Go! ")
	countFizz, countBuzz, countFizzBuzz, normal := 0, 0, 0, 0
	max := 1000
	for i := 1; i < max; i++ {
		result := fizzbuzz(i)
		if result == "fizzbuzz" {
		}

		switch result {
		case "fizzbuzz":
			countFizzBuzz = countFizzBuzz + 1
		case "fizz":
			countFizz = countFizz + 1
		case "buzz":
			countBuzz = countBuzz + 1
		default:
			normal = normal + 1
		}
	}
	fmt.Println("Normal : ", normal)
	fmt.Println("Fizz : ", countFizz)
	fmt.Println("Buzz : ", countBuzz)
	fmt.Println("Fizzbuzz : ", countFizzBuzz)

}

func fizzbuzz(n int) string {
	result := ""

	if isFizz(n) {
		result = "fizz"
	}

	if isBuzz(n) {
		result += "buzz"
	}

	if len(result) == 0 {
		return strconv.Itoa(n)
	}
	return result

}

func isFizz(n int) bool {
	return n%3 == 0
}

func isBuzz(n int) bool {
	return n%5 == 0
}
