package utils

import "slices"

// FindInSlice checks if an element is present in a Slice
func FindInSlice(slice []string, elementToFind string) bool {
	return slices.Contains(slice, elementToFind)
}
