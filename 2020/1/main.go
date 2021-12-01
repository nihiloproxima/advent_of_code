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
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	arr := strings.Split(string(data), "\n")

	for i := 0; i < len(arr); i++ {
		a := toInteger(arr[i])
		for j := 0; j < len(arr); j++ {
			b := toInteger(arr[j])
			for k := 0; k < len(arr); k++ {
				c := toInteger(arr[k])
				if a+b+c == 2020 {
					fmt.Printf("Found: %d\n", a*b*c)
					return
				}
			}
		}
	}
}
