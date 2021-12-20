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

func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func computeFuelCost(n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += 1 + i
	}

	return result
}

func resolvePart1(crabs []int) int {
	min, max := findMinAndMax(crabs)

	fmt.Println(min, max)

	previousSum := 10000000
	for i := min; i < max; i++ {
		sum := 0
		for _, crab := range crabs {
			if crab > i {
				sum += crab - i
			} else {
				sum += i - crab
			}
		}

		if sum < previousSum {
			previousSum = sum
		}
	}

	return previousSum
}

func resolvePart2(crabs []int) int {
	min, max := findMinAndMax(crabs)

	previousSum := 1000000000
	for i := min; i < max; i++ {
		sum := 0
		for _, crab := range crabs {
			var travelUnits int
			if crab > i {
				travelUnits = crab - i
			} else {
				travelUnits = i - crab
			}
			sum += computeFuelCost(travelUnits)
		}

		if sum < previousSum {
			previousSum = sum
		}
	}

	return previousSum
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}
	str := strings.Split(string(input), ",")

	var crabs []int
	for _, e := range str {
		crabs = append(crabs, toInteger(e))
	}

	fmt.Printf("Part 1: %d\n", resolvePart1(crabs))
	fmt.Printf("Part 2: %d\n", resolvePart2(crabs))
}
