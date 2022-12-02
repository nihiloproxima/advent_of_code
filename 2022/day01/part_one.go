package day01

import (
	"aoc_2022/utils"
)

func PartOne(data []string) int {
	inventories := []int{}
	tmp := 0
	max := 0

	for _, line := range data {
		if len(line) == 0 {
			inventories = append(inventories, tmp)
			if tmp > max {
				max = tmp
			}

			tmp = 0
		} else {
			value := utils.ToInteger(line)
			tmp += value
		}
	}

	return max
}
