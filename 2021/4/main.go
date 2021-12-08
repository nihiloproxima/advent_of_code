package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}

func arrayContains(arr []int, value int) bool {
	for _, e := range arr {
		if e == value {
			return true
		}
	}

	return false
}

func parseRolls(line string) []int {
	rolls := strings.Split(line, ",")

	var result []int
	for _, roll := range rolls {
		result = append(result, toInteger(roll))

	}

	return result
}

func parseLine(str string) []int {
	arr := strings.Split(str, " ")

	var result []int

	for _, e := range arr {
		if len(e) > 0 {
			result = append(result, toInteger(e))
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

func findPartOne(grids [][][]int, rolls []int) int {
	for _, roll := range rolls {
		for index, grid := range grids {
			if grids[index] = findAndReplace(grid, roll); gridIsValid(grids[index]) {
				return computeScore(grids[index], roll)
			}
		}
	}

	return 0
}

func findPartTwo(grids [][][]int, rolls []int) int {
	var validGridsIndexes []int
	var lastRoll int

	for _, roll := range rolls {
		for index, grid := range grids {
			if !gridIsValid((grids[index])) {
				if grids[index] = findAndReplace(grid, roll); gridIsValid(grids[index]) && !arrayContains(validGridsIndexes, index) {
					validGridsIndexes = append(validGridsIndexes, index)
					lastRoll = roll
				}
			}
		}
	}

	return computeScore(grids[validGridsIndexes[len(validGridsIndexes)-1]], lastRoll)
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")
	rolls := parseRolls(lines[0])
	grids := parseGrids(lines[2:])

	fmt.Printf("Part 1: %d\n", findPartOne(grids, rolls))
	fmt.Printf("Part 2: %d\n", findPartTwo(grids, rolls))
}
