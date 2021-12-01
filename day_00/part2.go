package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data), "\n")
	totalIncreases := 0
	previousSum := 0

	for i := 0; i < len(arr)-2; i++ {
		a, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println(err)
		}

		b, err := strconv.Atoi(arr[i+1])
		if err != nil {
			fmt.Println(err)
		}

		c, err := strconv.Atoi(arr[i+2])
		if err != nil {
			fmt.Println(err)
		}

		currentSum := a + b + c

		// First iteration should not count as an increase
		if i > 0 && currentSum > previousSum {
			totalIncreases++
		}

		previousSum = currentSum
	}
	fmt.Printf("%d\n", totalIncreases)
}
