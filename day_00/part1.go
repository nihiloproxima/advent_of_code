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
	sum := 0

	for i := 1; i < len(arr); i++ {
		a := toInteger((arr[i-1]))
		b := toInteger((arr[i]))

		if b > a {
			sum++
		}
	}

	fmt.Printf("%d\n", sum)
}
