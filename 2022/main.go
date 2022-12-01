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
		partOne, err := day01.PartOne()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Day 01 - Part 1:", partOne)

		partTwo, err := day01.PartTwo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Day 01 - Part 2:", partTwo)
	default:
		fmt.Println("Day not found")
	}
}
