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

func parseDrawNumber(str string) []int {
	str = str[:len(str)-1]
	arr := strings.Split(str, ",")

	var result []int
	for i := 0; i < len(arr); i++ {
		result = append(result, toInteger(arr[i]))
	}

	return result
}

func parseLine(str string) []int {
	str = str[:len(str)-1]
	arr := strings.Split(str, " ")

	var result []int
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > 0 {
			result = append(result, toInteger(arr[i]))
		}
	}

	return result
}

func parseGrids(lines []string) [][][]int {
	var grids [][][]int
	for i := 2; i < len(lines); i += 6 {
		var grid [][]int
		for j := 0; j < 5; j++ {
			grid = append(grid, parseLine(lines[i]))
		}

		grids = append(grids, grid)
	}

	return grids
}

func findPartOne(lines []string) int {
	drawNumbers := parseDrawNumber(lines[0])
	grids := parseGrids(lines)

	fmt.Println("Draw numbers", drawNumbers)
	fmt.Println("Grids", grids)

	return 0
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part 1: %d\n", findPartOne(lines))
}
