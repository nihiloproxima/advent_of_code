package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type speed struct {
	y int
	x int
}

func toInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}

func isTree(line string, x int) bool {
	return []rune(line)[x] == rune('#')
}

func countTreesEncounter(lines []string, speed speed) int {
	treeEncounters := 0

	x := 0
	for y := speed.y; y < len(lines); y += speed.y {
		lineLen := len(lines[y])
		x += speed.x
		if x >= lineLen {
			x -= lineLen
		}

		if isTree(lines[y], x) {
			treeEncounters++
		}
	}

	return treeEncounters
}

func countEncourters2(lines []string) int {
	configs := []speed{
		{
			y: 1,
			x: 1,
		},
		{
			y: 1,
			x: 3,
		},
		{
			y: 1,
			x: 5,
		},
		{
			y: 1,
			x: 7,
		},
		{
			y: 2,
			x: 1,
		},
	}

	total := 1
	for _, config := range configs {
		total *= countTreesEncounter(lines, config)
	}

	return total
}

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part 1: trees encounters =  %d\n", countTreesEncounter(lines, speed{
		y: 1,
		x: 3,
	}))
	fmt.Printf("Part 2: trees encoubters = %d\n", countEncourters2(lines))
}
