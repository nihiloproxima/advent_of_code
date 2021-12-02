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

func part1(arr []string) int {
	for i := 0; i < len(arr); i++ {
		a := toInteger(arr[i])
		for j := 0; j < len(arr); j++ {
			b := toInteger(arr[j])
			if a+b == 2020 {
				return a * b
			}
		}
	}

	panic(`No solution has been found.`)
}

func part2(arr []string) int {
	for i := 0; i < len(arr); i++ {
		a := toInteger(arr[i])
		for j := 0; j < len(arr); j++ {
			b := toInteger(arr[j])
			for k := 0; k < len(arr); k++ {
				c := toInteger(arr[k])
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	panic(`No solution has been found.`)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(data), "\n")

	fmt.Printf("Part 1 solution: %d\n", part1(arr))
	fmt.Printf("Part 2 solution: %d\n", part2(arr))
}
