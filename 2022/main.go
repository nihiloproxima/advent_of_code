package main

import (
	day01 "aoc_2022/day_01"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <day>")
		os.Exit(1)
	}

	day := args[0]

	switch day {
	case "1":
		fmt.Println("Day 01")
		fmt.Println("Part 1:", day01.PartOne())
		fmt.Println("Part 2:", day01.PartTwo())
	default:
		fmt.Println("Day not found")
	}
}
