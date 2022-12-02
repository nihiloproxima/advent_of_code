package day03

import (
	"strconv"
	"strings"
)

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

func FindMostCommonBit(grid [][]string, position int) string {
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

func PartOne(lines []string) int64 {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, parseLine(line))
	}

	var gamaRateBits []string
	var epsilonRateBits []string
	for i := 0; i < len(grid[0]); i++ {
		mostCommonBit := FindMostCommonBit(grid, i)
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
