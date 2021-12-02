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

func extractPolicy(str string) (int, int) {
	minMax := strings.Split(str, "-")

	return toInteger(minMax[0]), toInteger(minMax[1])
}

func countFirstPartValidPasswords(arr []string) int {
	validPasswordsTotal := 0

	for _, entry := range arr {
		elements := strings.Split(entry, " ")
		min, max := extractPolicy(elements[0])
		char := strings.Split(elements[1], ":")[0]
		password := elements[2]
		occurences := strings.Count(password, char)

		if occurences >= min && occurences <= max {
			validPasswordsTotal++
		}
	}

	return validPasswordsTotal
}

func countSecondPartValidPasswords(arr []string) int {
	validPasswordsTotal := 0

	for _, entry := range arr {
		elements := strings.Split(entry, " ")
		first, second := extractPolicy(elements[0])
		char := strings.Split(elements[1], ":")[0]
		password := elements[2]

		v1 := password[first-1 : first]
		v2 := password[second-1 : second]

		count := 0

		if v1 == char {
			count++
		}
		if v2 == char {
			count++
		}

		if count == 1 {
			validPasswordsTotal++
		}
	}

	return validPasswordsTotal
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(data), "\n")

	fmt.Printf("1st part valid passwords total: %d\n", countFirstPartValidPasswords(arr))
	fmt.Printf("2nd part valid passwords total: %d\n", countSecondPartValidPasswords(arr))
}
