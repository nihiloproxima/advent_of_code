package day06

import "aoc_2021/utils"

func PartTwo(data []string) int {
	var firstGen []int
	for _, e := range data {
		firstGen = append(firstGen, utils.ToInteger(e))
	}

	return resolve(firstGen, 256)
}
