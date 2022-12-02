package day03

import (
	"strconv"
	"strings"
)

func PartTwo(lines []string) int64 {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, parseLine(line))
	}

	oxigenRatings := grid
	co2Ratings := grid
	for i := 0; i < len(grid[0]); i++ {
		if len(oxigenRatings) > 1 {
			mostCommonBit := FindMostCommonBit(oxigenRatings, i)
			for _, line := range oxigenRatings {
				if charAt := line[i]; charAt != mostCommonBit {
					oxigenRatings = removeFromArray(oxigenRatings, line)
				}
			}
		}

		if len(co2Ratings) > 1 {
			mostCommonBit := FindMostCommonBit(co2Ratings, i)
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
