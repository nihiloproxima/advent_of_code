package day02

import (
	"aoc_2021/utils"
	"strings"
)

func PartOne(lines []string) int {
	var position, depth int

	for _, line := range lines {
		instruction := strings.Split(line, " ")[0]
		value := utils.ToInteger(strings.Split(line, " ")[1])

		switch instruction {
		case "up":
			depth -= value
		case "down":
			depth += value
		case "forward":
			position += value
		default:
			panic("Unknown instruction")
		}
	}

	return position * depth
}
