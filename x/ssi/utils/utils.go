package utils

// FindInSlice checks if an element is present in a Slice
func FindInSlice(slice []string, elementToFind string) bool {
	for _, elem := range slice {
		if elem == elementToFind {
			return true
		}
	}
	return false
}
