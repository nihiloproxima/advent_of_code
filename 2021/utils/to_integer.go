package utils

import (
	"strconv"
)

func ToInteger(str string) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return result
}
