package day04

import "aoc_2021/utils"

func PartTwo(data []string) int {
	rolls := parseRolls(data[0])
	grids := parseGrids(data[2:])

	var validGridsIndexes []int
	var lastRoll int

	for _, roll := range rolls {
		for index, grid := range grids {
			if !gridIsValid((grids[index])) {
				if grids[index] = findAndReplace(grid, roll); gridIsValid(grids[index]) && !utils.ArrayContains(validGridsIndexes, index) {
					validGridsIndexes = append(validGridsIndexes, index)
					lastRoll = roll
				}
			}
		}
	}

	return computeScore(grids[validGridsIndexes[len(validGridsIndexes)-1]], lastRoll)
}
