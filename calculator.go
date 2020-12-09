// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	var sum float64 = a + b
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Subtract takes two or more numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, nums ...float64) float64 {
	var total float64 = a - b
	for _, num := range nums {
		total -= num
	}
	return total
}

// Multiply takes two or more numbers and returns the result of multiplying them together
func Multiply(a, b float64, nums ...float64) float64 {
	var total float64 = a * b
	for _, num := range nums {
		total *= num
	}
	return total
}

// Divide takes two or more numbers and returns the result of dividing a by b
func Divide(a, b float64, nums ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	// if the first number is zero then the result is zero.
	if a == 0 {
		return 0, nil
	}

	total := a / b
	for _, num := range nums {
		if num == 0 {
			return 0, errors.New("divide by Zero")
		}
		total /= num
		// if total is ever zero we can immediately return
		// zero.
		if total == 0 {
			return 0, nil
		}
	}
	return total, nil
}

// Sqrt returns the square root of n
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is undefined", n)
	}
	return math.Sqrt(n), nil
}

// ExpressionProcessor performs the computation passed in string
func ExpressionProcessor(expression string) (float64, error) {
	var a, b float64
	var operator string

	count, err := fmt.Sscanf(expression, "%f%1s%f", &a, &operator, &b)

	if err != nil {
		return 0, err
	}

	if count != 3 {
		return 0, fmt.Errorf("ExpressionProcessor was passed an expression with %d arguments when 3 are required", count)
	}

	switch operator {
	case "+":
		return Add(a, b), nil
	case "-":
		return Subtract(a, b), nil
	case "*":
		return Multiply(a, b), nil
	case "/":
		return Divide(a, b)
	default:
		return 0, fmt.Errorf("ExpressionProcessor was passed an unknown operation in expression %s", expression)
	}

}
