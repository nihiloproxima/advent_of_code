package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}

func firstPart(arr []string) int {
	sum := 0

	for i := 1; i < len(arr); i++ {
		a := toInteger(arr[i-1])
		b := toInteger(arr[i])

		if b > a {
			sum++
		}
	}

	return sum
}

func secondPart(arr []string) int {
	totalIncreases := -1 // first iteration does not count
	previous := 0

	for i := 2; i < len(arr); i++ {
		a := toInteger(arr[i])
		b := toInteger(arr[i-1])
		c := toInteger(arr[i-2])

		current := a + b + c

		if current > previous {
			totalIncreases++
		}

		previous = current
	}

	return totalIncreases
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data), "\n")

	fmt.Printf("First part result: %d\n", firstPart(arr))
	fmt.Printf("Second part result: %d\n", secondPart(arr))
}
