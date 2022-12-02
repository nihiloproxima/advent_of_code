package day07

import (
	"aoc_2021/utils"
	"fmt"
)

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

func PartOne(data []string) int {
	var crabs []int
	for _, e := range data {
		crabs = append(crabs, utils.ToInteger(e))
	}

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
