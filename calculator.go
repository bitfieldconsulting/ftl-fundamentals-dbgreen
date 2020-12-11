// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"math"
)

// Add takes two or more numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	sum := a + b
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Subtract takes two or more numbers and returns the result of subtracting
// the second from the first.
func Subtract(a, b float64, nums ...float64) float64 {
	total := a - b
	for _, num := range nums {
		total -= num
	}
	return total
}

// Multiply takes two or more numbers, multiplies the first by the second,
// and then multiplies the result by any subsequent numbers, and returns
// the result.
func Multiply(a, b float64, nums ...float64) float64 {
	total := a * b
	for _, num := range nums {
		total *= num
	}
	return total
}

// Divide takes two or more numbers, divides the first by the second,
// and then divides the result by any subsequent numbers, and returns
// the result, or an error if any but the first number is a zero.
func Divide(a, b float64, nums ...float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	total := a / b
	for _, num := range nums {
		if num == 0 {
			return 0, fmt.Errorf("divide by Zero")
		}
		total /= num
	}
	return total, nil
}

// Sqrt takes a number and returns either its square root or
// an error if the number is negative.
func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is undefined", n)
	}
	return math.Sqrt(n), nil
}

// EvaluateExpression takes two numbers separated by a binary operator and evaluates it returning
// the result, or an error if the expression is invalid.
func EvaluateExpression(expression string) (float64, error) {
	var a, b float64
	var operator string
	_, err := fmt.Sscanf(expression, "%f%1s%f", &a, &operator, &b)
	if err != nil {
		return 0, err
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
		return 0, fmt.Errorf("unknown operation in %s", expression)
	}
}
