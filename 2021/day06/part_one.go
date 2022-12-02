package day06

import (
	"aoc_2021/utils"
)

func resolve(fishs []int, days int) int {
	ages := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, fish := range fishs {
		ages[fish]++
	}

	for day := 0; day < days; day++ {
		newBorns := 0
		for i := 0; i < 9; i++ {
			if ages[i] > 0 {
				if i == 0 {
					newBorns = ages[i]
					ages[i] = 0
				} else {
					ages[i-1] = ages[i]
					if i < 8 {
						ages[i] = ages[i+1]
					} else {
						ages[i] = 0
					}
				}
			}
		}

		ages[6] += newBorns
		ages[8] += newBorns

	}

	sum := 0
	for _, e := range ages {
		sum += e
	}

	return sum
}

func PartOne(data []string) int {
	var firstGen []int
	for _, e := range data {
		firstGen = append(firstGen, utils.ToInteger(e))
	}

	return resolve(firstGen, 80)
}
