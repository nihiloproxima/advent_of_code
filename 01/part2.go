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

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data), "\n")
	totalIncreases := 0
	previousSum := 0

	for i := 0; i < len(arr)-2; i++ {
		a := toInteger(arr[i])
		b := toInteger(arr[i+1])
		c := toInteger(arr[i+2])

		currentSum := a + b + c

		// First iteration should not count as an increase
		if i > 0 && currentSum > previousSum {
			totalIncreases++
		}

		previousSum = currentSum
	}

	fmt.Printf("%d\n", totalIncreases)
}
