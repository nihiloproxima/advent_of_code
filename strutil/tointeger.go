package strutil

import "strconv"

// ToInteger converts a string to integer
func ToInteger(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return intValue
}
