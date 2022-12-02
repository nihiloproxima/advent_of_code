package main

import (
	"aoc_2021/day01"
	"aoc_2021/day02"
	"aoc_2021/day03"
	"aoc_2021/day04"
	"aoc_2021/day05"
	"aoc_2021/day06"
	"aoc_2021/day07"
	"aoc_2021/utils"
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
		data := utils.ParseFile("day01/data", "\n")
		fmt.Println("Day 01 - Part 1:", day01.PartOne(data))
		fmt.Println("Day 01 - Part 2:", day01.PartTwo(data))
	case "2":
		data := utils.ParseFile("day02/data", "\n")
		fmt.Println("Day 02 - Part 1:", day02.PartOne(data))
		fmt.Println("Day 02 - Part 2:", day02.PartTwo(data))
	case "3":
		data := utils.ParseFile("day03/data", "\n")
		fmt.Println("Day 03 - Part 1:", day03.PartOne(data))
		fmt.Println("Day 03 - Part 2:", day03.PartTwo(data))
	case "4":
		data := utils.ParseFile("day04/data", "\n")
		fmt.Println("Day 04 - Part 1:", day04.PartOne(data))
		fmt.Println("Day 04 - Part 2:", day04.PartTwo(data))
	case "5":
		data := utils.ParseFile("day05/data", "\n")
		fmt.Println("Day 05 - Part 1:", day05.PartOne(data))
		fmt.Println("Day 05 - Part 2:", day05.PartTwo(data))
	case "6":
		data := utils.ParseFile("day06/data", ",")
		fmt.Println("Day 06 - Part 1:", day06.PartOne(data))
		fmt.Println("Day 06 - Part 2:", day06.PartTwo(data))
	case "7":
		data := utils.ParseFile("day07/data", ",")
		fmt.Println("Day 07 - Part 1:", day07.PartOne(data))
		fmt.Println("Day 07 - Part 2:", day07.PartTwo(data))
	default:
		fmt.Println("Day not found")
	}
}
