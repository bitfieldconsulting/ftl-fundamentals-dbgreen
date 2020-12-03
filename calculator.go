// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
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
func Multiply(nums ...float64) float64 {
	var total float64 = nums[0]
	for i := 1; i < len(nums); i++ {
		total *= nums[i]
	}
	return total
}

// Divide takes two numbers and returns the result of dividing a by b
func Divide(nums ...float64) (float64, error) {
	if nums[0] == 0 {
		return 0, errors.New("divide by zero")
	}
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, errors.New("divide by Zero")
		}
		total /= nums[i]
	}
	return total, nil
}

// Sqrt returns the square root of n
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("sqrt of negative number")
	}
	return math.Sqrt(n), nil
}

// Parse performs the computation passed in string
func Parse(oper string) (float64, error) {
	var a, b float64
	var operator string
	count, err := fmt.Sscanf(oper, "%f%1s%f", &a, &operator, &b)
	if err != nil {
		return 0, err
	}
	if count != 3 {
		return 0, errors.New("wrong format")
	}
	nums := []float64{a, b}
	switch operator {
	case "+":
		return Add(nums...), nil
	case "-":
		return Subtract(nums...), nil
	case "*":
		return Multiply(nums...), nil
	case "/":
		return Divide(nums...)
	default:
		return 0, errors.New("unknown operation")
	}

}
