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

func parseLine(line string) []string {
	var res []string
	for _, lineByte := range line {
		res = append(res, string(lineByte))
	}

	return res
}

func isEqual(val1, val2 []string) bool {
	if len(val1) != len(val2) {
		return false
	}

	for i := 0; i < len(val1); i++ {
		if val1[i] != val2[i] {
			return false
		}
	}

	return true
}

func removeFromArray(arr [][]string, toRemove []string) [][]string {
	var result [][]string
	for _, e := range arr {
		if !isEqual(e, toRemove) {
			result = append(result, e)
		}
	}
	return result
}

func findMostCommonBit(grid [][]string, position int) string {
	var one int
	for _, line := range grid {
		if line[position] == "1" {
			one++
		} else {
			one--
		}
	}

	if one >= 0 {
		return "1"
	}

	return "0"
}

func partOne(lines []string) int64 {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, parseLine(line))
	}

	var gamaRateBits []string
	var epsilonRateBits []string
	for i := 0; i < len(grid[0]); i++ {
		mostCommonBit := findMostCommonBit(grid, i)
		gamaRateBits = append(gamaRateBits, mostCommonBit)
		if mostCommonBit == "1" {
			epsilonRateBits = append(epsilonRateBits, "0")
		} else {
			epsilonRateBits = append(epsilonRateBits, "1")
		}
	}

	gamaRate, err := strconv.ParseInt(strings.Join(gamaRateBits, ""), 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonRate, err := strconv.ParseInt(strings.Join(epsilonRateBits, ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return gamaRate * epsilonRate
}

func partTwo(lines []string) int64 {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, parseLine(line))
	}

	oxigenRatings := grid
	co2Ratings := grid
	for i := 0; i < len(grid[0]); i++ {
		if len(oxigenRatings) > 1 {
			mostCommonBit := findMostCommonBit(oxigenRatings, i)
			for _, line := range oxigenRatings {
				if charAt := line[i]; charAt != mostCommonBit {
					oxigenRatings = removeFromArray(oxigenRatings, line)
				}
			}
		}

		if len(co2Ratings) > 1 {
			mostCommonBit := findMostCommonBit(co2Ratings, i)
			for _, line := range co2Ratings {
				if charAt := line[i]; charAt == mostCommonBit {
					co2Ratings = removeFromArray(co2Ratings, line)
				}
			}
		}
	}

	oxigenRating, err := strconv.ParseInt(strings.Join(oxigenRatings[0], ""), 2, 64)
	if err != nil {
		panic(err)
	}
	co2Rating, err := strconv.ParseInt(strings.Join(co2Ratings[0], ""), 2, 64)
	if err != nil {
		panic(err)
	}

	return oxigenRating * co2Rating
}

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(input), "\n")

	fmt.Println("Part 1:", partOne(lines))
	fmt.Println("Part 2:", partTwo(lines))

}
