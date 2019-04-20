// Package grains does the thing
package grains

import "fmt"

// method Square returns the number of grains on a given square
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Input[%d] invalid. Input should be between 1 and 64 (inclusive)", n)
	}
	return 1 << (uint(n) - 1), nil
}

// method Total returns the total number of grains
func Total() uint64 {
	return (1 << 64) -1
}