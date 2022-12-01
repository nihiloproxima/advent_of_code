package day01

import (
	"aoc_2022/utils"
)

func PartOne() (int, error) {
	data, err := utils.ParseFile("./day_01/data")
	if err != nil {
		return 0, err
	}

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

	return max, nil
}
