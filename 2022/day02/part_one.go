package day02

import (
	"strings"
)

// ROCK = A X
// PAPER = B Y
// SCISSORS = C Z

func PartOne(data []string) int {
	total := 0
	for _, match := range data {
		if len(match) == 0 {
			continue
		}

		var matchResultPoints = map[string]int{
			"WIN":  6,
			"DRAW": 3,
			"LOSE": 0,
		}
		var moves = map[string]string{
			"A Y": "WIN",
			"B Z": "WIN",
			"C X": "WIN",
			"A X": "DRAW",
			"B Y": "DRAW",
			"C Z": "DRAW",
			"A Z": "LOSE",
			"B X": "LOSE",
			"C Y": "LOSE",
		}
		var coupPoints = map[string]int{"X": 1, "Y": 2, "Z": 3}

		total += matchResultPoints[moves[match]] + coupPoints[strings.Split(match, " ")[1]]

	}

	return total
}
