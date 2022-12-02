package day07

import "aoc_2021/utils"

func PartTwo(data []string) int {
	var crabs []int
	for _, e := range data {
		crabs = append(crabs, utils.ToInteger(e))
	}

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
