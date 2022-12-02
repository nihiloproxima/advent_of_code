package day01

import (
	"aoc_2021/utils"
)

func PartTwo(arr []string) int {
	totalIncreases := -1 // first iteration does not count
	previous := 0

	for i := 2; i < len(arr); i++ {
		a := utils.ToInteger(arr[i])
		b := utils.ToInteger(arr[i-1])
		c := utils.ToInteger(arr[i-2])

		current := a + b + c

		if current > previous {
			totalIncreases++
		}

		previous = current
	}

	return totalIncreases
}
