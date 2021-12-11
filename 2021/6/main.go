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

func main() {
	input, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println(err)
	}
	str := strings.Split(string(input), ",")

	var firstGen []int
	for _, e := range str {
		firstGen = append(firstGen, toInteger(e))
	}

	fmt.Printf("Part 1: %d\n", resolve(firstGen, 80))
	fmt.Printf("Part 2: %d\n", resolve(firstGen, 256))
}
