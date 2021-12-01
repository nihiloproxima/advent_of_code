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
	sum := 0

	for i := 1; i < len(arr); i++ {
		a, err := strconv.Atoi(arr[i-1])
		if err != nil {
			fmt.Println(err)
		}

		b, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println(err)
		}

		if b > a {
			sum++
		}
	}

	fmt.Printf("%d\n", sum)
}
