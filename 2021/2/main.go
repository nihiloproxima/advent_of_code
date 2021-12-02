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

func findDepthPart1(lines []string) int {
	var position, depth int

	for _, line := range lines {
		instruction := strings.Split(line, " ")[0]
		value := toInteger(strings.Split(line, " ")[1])

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

func findDepthPart2(lines []string) int {
	var position, depth, aim int

	for _, line := range lines {
		instruction := strings.Split(line, " ")[0]
		value := toInteger(strings.Split(line, " ")[1])

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

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Printf("Part 1 depth: %d\n", findDepthPart1(lines))
	fmt.Printf("Part 2 depth: %d\n", findDepthPart2(lines))
}
