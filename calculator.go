// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(nums ...float64) float64 {
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(nums ...float64) float64 {
	var total float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		total -= nums[i]
	}
	return total
}

// Multiply takes two numbers and returns the result of multiplying them together
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing a by b
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by Zero")
	}
	return a / b, nil
}

// Sqrt returns the square root of n
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("sqrt of negative number")
	}
	return math.Sqrt(n), nil
}
