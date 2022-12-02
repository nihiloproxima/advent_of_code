package day01

import (
	"aoc_2022/utils"
	"sort"
)

func PartTwo(data []string) int {
	inventories := []int{}
	tmp := 0

	for _, line := range data {
		if len(line) == 0 {
			inventories = append(inventories, tmp)

			tmp = 0
		} else {
			value := utils.ToInteger(line)
			tmp += value
		}
	}

	sort.Ints(inventories)

	sum := 0
	for _, n := range inventories[len(inventories)-3:] {
		sum += n
	}

	return sum
}
