package utils

import (
	"fmt"
	"os"
	"strings"
)

func ParseFile(path string, separator string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	arr := strings.Split(string(data), separator)

	return arr
}
