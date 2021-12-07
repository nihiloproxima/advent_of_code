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

func parseDrawNumber(str string) []int {
	arr := strings.Split(str, ",")

	var result []int
	for i := 0; i < len(arr); i++ {
		result = append(result, toInteger(arr[i]))
	}

	return result
}

func parseLine(str string) []int {
	arr := strings.Split(str, " ")

	var result []int

	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			integer := toInteger(arr[i])
			result = append(result, integer)
		}
	}

	return result
}

func parseGrids(lines []string) [][][]int {
	var grids [][][]int
	for i := 2; i < len(lines); i += 6 {
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
		if lineIsValid(grid[i]) {
			return true
		}
		for j := 0; j < 5; j++ {
			if colIsValid(grid, j) {
				return true
			}
		}
	}

	return false
}

func lineIsValid(grid []int) bool {
	found := 0
	for i := 0; i < 5; i++ {
		if grid[i] == 1000 {
			found++
		}
	}

	return found == 5
}

func colIsValid(grid [][]int, colPosition int) bool {
	found := 0
	for i := 0; i < 5; i++ {
		if grid[i][colPosition] == 1000 {
			found++
		}
	}

	return found == 5
}

func soluce(grid [][]int, drawed int) int {
	sum := 0
	for _, line := range grid {
		for i := 0; i < 5; i++ {
			if line[i] != 1000 {
				sum += line[i]
			}
		}
	}

	return sum * drawed
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

func findPartOne(grids [][][]int, drawNumbers []int) int {

	for i := 0; i < len(drawNumbers); i++ {
		for index, grid := range grids {
			grids[index] = findAndReplace(grid, drawNumbers[i])
			if gridIsValid(grids[index]) {
				return soluce(grids[index], drawNumbers[i])
			}
		}
	}

	return 0
}

func findPartTwo(grids [][][]int, drawNumbers []int) int {
	var validGridsIndexes []int
	var lastValidDraw int

	for i := 0; i < len(drawNumbers); i++ {
		for index, grid := range grids {
			if !gridIsValid((grids[index])) {
				grids[index] = findAndReplace(grid, drawNumbers[i])
				if gridIsValid(grids[index]) && !arrayContains(validGridsIndexes, index) {
					validGridsIndexes = append(validGridsIndexes, index)
					lastValidDraw = drawNumbers[i]
				}
			}
		}
	}

	return soluce(grids[validGridsIndexes[len(validGridsIndexes)-1]], lastValidDraw)
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	drawNumbers := parseDrawNumber(lines[0])
	grids := parseGrids(lines)

	fmt.Printf("Part 1: %d\n", findPartOne(grids, drawNumbers))
	fmt.Printf("Part 2: %d\n", findPartTwo(grids, drawNumbers))
}
