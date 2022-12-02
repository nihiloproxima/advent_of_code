package utils

func ArrayContains(arr []int, value int) bool {
	for _, e := range arr {
		if e == value {
			return true
		}
	}

	return false
}
