package utils

// Contains searches slice for a specified value
func Contains[K comparable](in []K, val K) bool {
	for _, v := range in {
		if v == val {
			return true
		}
	}

	return false
}
