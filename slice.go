// Package slice set of functions for working with slices
package slice

import (
	"errors"
	"math"
)

// NewNativeArray creates a new own array for the result
func NewNativeArray(slice []string) []string {
	newSlice := make([]string, len(slice))
	copy(newSlice, slice)
	return newSlice
}

// Chunk splits an array into parts.
// Splits the array into multiple arrays of length elements.
// The last array retrieved may contain fewer values than length.
// The first argument to flags can be set to the boolean value true to indicate that a new native array should be created for the returned result.
func Chunk(slice []string, length int, flags ...bool) ([][]string, error) {
	// argument checking
	if length <= 0 {
		return nil, errors.New("length must be greater than zero")
	}
	if len(flags) > 0 && flags[0] {
		slice = NewNativeArray(slice)
	}
	capacity := int(math.Ceil(float64(len(slice)) / float64(length)))
	result := make([][]string, 0, capacity)
	for part := 0; part < len(slice); part += length {
		step := part + length
		if step <= len(slice) {
			result = append(result, slice[part:step])
		} else {
			result = append(result, slice[part:])
		}
	}
	return result, nil
}

// CountValues counts the number of all slice values.
// Returns a map with indexes of slice values and values indicating the number of items found.
func CountValues(slice []string) map[string]int {
	result := make(map[string]int)
	for _, item := range slice {
		result[item]++
	}
	return result
}
