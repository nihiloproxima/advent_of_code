package day01

import (
	"aoc_2021/utils"
)

func PartOne(arr []string) int {
	sum := 0

	for i := 1; i < len(arr); i++ {
		a := utils.ToInteger(arr[i-1])
		b := utils.ToInteger(arr[i])

		if b > a {
			sum++
		}
	}

	return sum
}
