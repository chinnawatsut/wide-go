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
	fmt.Println("Normal : ", normal, "Percentage:", percentOf(normal, max), "%")
	fmt.Println("Fizz : ", countFizz, "Percentage:", percentOf(countFizz, max), "%")
	fmt.Println("Buzz : ", countBuzz, "Percentage:", percentOf(countBuzz, max), "%")
	fmt.Println("Fizzbuzz : ", countFizzBuzz, "Percentage:", percentOf(countFizzBuzz, max), "%")

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

func percentOf(current int, all int) float64 {
	percent := (float64(current) * float64(100)) / float64(all)
	return percent
}
