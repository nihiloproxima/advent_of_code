package day01

import (
	"aoc_2022/utils"
)

func PartOne() int {
	data := utils.ParseFile("./day_01/data")
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
