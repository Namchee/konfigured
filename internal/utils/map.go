package utils

// FilterKeysByValue returns a subset of map that satisfies the specified value
func FilterKeysByValue[K comparable, V comparable](in map[K]V, val V) map[K]V {
	filtered := map[K]V{}

	for k, v := range in {
		if v == val {
			filtered[k] = v
		}
	}

	return filtered
}
