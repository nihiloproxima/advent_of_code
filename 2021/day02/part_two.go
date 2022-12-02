package day02

import (
	"aoc_2021/utils"
	"strings"
)

func PartTwo(lines []string) int {
	var position, depth, aim int

	for _, line := range lines {
		instruction := strings.Split(line, " ")[0]
		value := utils.ToInteger(strings.Split(line, " ")[1])

		switch instruction {
		case "up":
			aim -= value
		case "down":
			aim += value
		case "forward":
			position += value
			depth += aim * value
		default:
			panic("Unknown instruction")
		}
	}

	return position * depth
}
