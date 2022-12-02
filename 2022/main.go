package main

import (
	"aoc_2022/day01"
	"aoc_2022/day02"
	"aoc_2022/utils"
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
		data := utils.ParseFile("day01/data")
		fmt.Println("Day 01 - Part 1:", day01.PartOne(data))
		fmt.Println("Day 01 - Part 2:", day01.PartTwo(data))
	case "2":
		data := utils.ParseFile("day02/data")
		fmt.Println("Day 02 - Part 1:", day02.PartOne(data))
		fmt.Println("Day 02 - Part 2:", day02.PartTwo(data))
	default:
		fmt.Println("Day not found")
	}
}
