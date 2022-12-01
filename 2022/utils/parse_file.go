package utils

import (
	"os"
	"strings"
)

func ParseFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return make([]string, 0), err
	}

	arr := strings.Split(string(data), "\n")

	return arr, nil
}
