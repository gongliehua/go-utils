package utils

func InSlice[T comparable](elem T, slice []T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
