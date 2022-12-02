package day04

import (
	"aoc_2021/utils"
	"strings"
)

func parseRolls(line string) []int {
	rolls := strings.Split(line, ",")

	var result []int
	for _, roll := range rolls {
		result = append(result, utils.ToInteger(roll))

	}

	return result
}

func parseLine(str string) []int {
	arr := strings.Split(str, " ")

	var result []int

	for _, e := range arr {
		if len(e) > 0 {
			result = append(result, utils.ToInteger(e))
		}
	}

	return result
}

func parseGrids(lines []string) [][][]int {
	var grids [][][]int

	for i := 0; i < len(lines); i += 6 {
		var grid [][]int
		for j := 0; j < 5; j++ {
			grid = append(grid, parseLine(lines[i+j]))
		}
		grids = append(grids, grid)
	}

	return grids
}

func gridIsValid(grid [][]int) bool {
	for i := 0; i < 5; i++ {
		if horizontallyValid(grid[i]) || verticallyCorrect(grid, i) {
			return true
		}
	}

	return false
}

func horizontallyValid(line []int) bool {
	for _, num := range line {
		if num != 1000 {
			return false
		}
	}

	return true
}

func verticallyCorrect(grid [][]int, x int) bool {
	for _, line := range grid {
		if line[x] != 1000 {
			return false
		}
	}
	return true
}

func computeScore(grid [][]int, roll int) int {
	sum := 0
	for _, line := range grid {
		for _, num := range line {
			if num != 1000 {
				sum += num
			}
		}
	}

	return sum * roll
}

func findAndReplace(grid [][]int, num int) [][]int {
	var newGrid = make([][]int, 5)
	for i := range newGrid {
		newGrid[i] = make([]int, 5)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == num {
				newGrid[y][x] = 1000
			} else {
				newGrid[y][x] = grid[y][x]
			}
		}
	}

	return newGrid
}

func PartOne(data []string) int {
	rolls := parseRolls(data[0])
	grids := parseGrids(data[2:])

	for _, roll := range rolls {
		for index, grid := range grids {
			if grids[index] = findAndReplace(grid, roll); gridIsValid(grids[index]) {
				return computeScore(grids[index], roll)
			}
		}
	}

	return 0
}
