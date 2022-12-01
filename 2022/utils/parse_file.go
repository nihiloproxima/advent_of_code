package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParseFile(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	arr := strings.Split(string(data), "\n")

	return arr
}
