package day02

import (
	"strings"
)

// ROCK = A
// PAPER = B
// SCISSORS = C

func PartTwo(data []string) int {
	total := 0
	for _, match := range data {
		if len(match) == 0 {
			continue
		}

		var LOSE_MOVE = map[string]string{
			"A": "C",
			"B": "A",
			"C": "B",
		}
		var WIN_MOVE = map[string]string{
			"A": "B",
			"B": "C",
			"C": "A",
		}
		var matchResultPoints = map[string]int{
			"WIN":  6,
			"DRAW": 3,
			"LOSE": 0,
		}
		var coupPoints = map[string]int{"A": 1, "B": 2, "C": 3}

		opponentCoup := strings.Split(match, " ")[0]
		expectedMatchOutcome := strings.Split(match, " ")[1]

		if expectedMatchOutcome == "X" {
			// LOST
			total += matchResultPoints["LOSE"] + coupPoints[LOSE_MOVE[opponentCoup]]

		} else if expectedMatchOutcome == "Z" {
			// WIN
			total += matchResultPoints["WIN"] + coupPoints[WIN_MOVE[opponentCoup]]

		} else {
			// DRAW
			total += matchResultPoints["DRAW"] + coupPoints[opponentCoup]
		}
	}

	return total
}
