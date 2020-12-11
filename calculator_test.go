package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func closeEnough(a, b, epsilon float64) bool {
	if math.Abs(a-b) < epsilon {
		return true
	}
	return false
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestAddRand(t *testing.T) {
	t.Parallel()
	count := 0
	for count < 100 {
		nums := make([]float64, random.Intn(10))
		for i := range nums {
			nums[i] = random.NormFloat64()
		}
		arg1 := random.NormFloat64()
		arg2 := random.NormFloat64()
		want := arg1 + arg2
		for _, val := range nums {
			want += val
		}
		got := calculator.Add(arg1, arg2, nums...)
		if !closeEnough(want, got, 0.0001) {
			t.Errorf("Add (%f,%f,%v): want %f, got %f", arg1, arg2, nums, want, got)
		}
		count++
	}
}
func TestAddMultSub(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		operation func(float64, float64, ...float64) float64
		name      string
		arg1      float64
		arg2      float64
		nums      []float64
		want      float64
	}{
		{
			operation: calculator.Add,
			name:      "Add: multiple numbers that sum to a positive",
			arg1:      5,
			arg2:      3,
			nums:      []float64{2, 3, 4},
			want:      17,
		},
		{
			operation: calculator.Add,
			name:      "Add: multiple numbers that sum to a negative",
			arg1:      5,
			arg2:      -3,
			nums:      []float64{-2, -3, -4},
			want:      -7,
		},
		{
			operation: calculator.Multiply,
			name:      "Negative number multiplied with a positive results in negative",
			arg1:      -1,
			arg2:      -1,
			nums:      []float64{-1, 5},
			want:      -5,
		},
		{
			operation: calculator.Multiply,
			name:      "Negative numbers multiplied result in a postive",
			arg1:      1,
			arg2:      1,
			nums:      []float64{-3, -9},
			want:      27,
		},
		{
			operation: calculator.Multiply,
			name:      "Number multiplied by zero results in zero",
			arg1:      0,
			arg2:      5,
			nums:      []float64{5, 0},
			want:      0,
		},
		{
			operation: calculator.Subtract,
			name:      "Numbers that subtract to a negative",
			arg1:      -5,
			arg2:      -5,
			nums:      []float64{-1, 2},
			want:      -1,
		},
		{
			operation: calculator.Subtract,
			name:      "Numbers that subtract to a positive",
			arg1:      1000,
			arg2:      800,
			nums:      []float64{100, 99},
			want:      1,
		},
		{
			operation: calculator.Subtract,
			name:      "Numbers that subtract to zero",
			arg1:      9,
			arg2:      3,
			nums:      []float64{3, 3},
			want:      0,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.operation(tc.arg1, tc.arg2, tc.nums...)
			if !closeEnough(tc.want, got, 0.0001) {
				t.Errorf("with args %f,%f,%v want %f, got %f", tc.arg1, tc.arg2, tc.nums, tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		arg1        float64
		arg2        float64
		nums        []float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Zero in optional parameters results in an error",
			arg1:        5,
			arg2:        7,
			nums:        []float64{6, 0},
			want:        0,
			errExpected: true,
		},
		{
			name:        "Zero as second argument results in an error",
			arg1:        5,
			arg2:        0,
			want:        0,
			errExpected: true,
		},
		{
			name: "Zero as first argument returns a zero",
			arg1: 0,
			arg2: 5,
			want: 0,
		},
		{
			name: "Number divided by another number results in an integer result",
			arg1: 12,
			arg2: 2,
			nums: []float64{2, 3},
			want: 1,
		},
		{
			name: "Number divided by another number results in a rational result",
			arg1: 1000,
			arg2: 10,
			nums: []float64{10, 4},
			want: 2.5,
		},
		{
			name: "Irrational result is 'close enough'",
			arg1: 1,
			arg2: 3,
			want: .33333,
		},
	}

	for _, tc := range tcs {
		got, err := calculator.Divide(tc.arg1, tc.arg2, tc.nums...)

		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%f,%f,%v): unexpected error status %v in TestCase: %s", tc.arg1, tc.arg2, tc.nums, err, tc.name)
		}
		expectedResult := closeEnough(tc.want, got, 0.0001)
		if !tc.errExpected && !expectedResult {
			t.Errorf("Divide(%f,%f,%v): wanted %f, got %f in TestCase: %s", tc.arg1, tc.arg2, tc.nums, tc.want, got, tc.name)
		}

	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Sqrt of negative results in an error",
			a:           -1,
			want:        0,
			errExpected: true,
		},
		{
			name: "Sqrt of zero results in a zero",
			a:    0,
			want: 0,
		},
		{
			name: "Sqrt of positive number results in positive number",
			a:    16,
			want: 4,
		},
	}

	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.a)

		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Sqrt(%f): unexpected error status %v in TestCase: %s", tc.a, err, tc.name)
		}
		expectedResult := closeEnough(tc.want, got, 0.0001)
		if !tc.errExpected && !expectedResult {
			t.Errorf("Sqrt(%f): wanted %f, got %f in TestCase: %s", tc.a, tc.want, got, tc.name)
		}

	}
}

func TestEvaluateExpression(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		expr        string
		want        float64
		errExpected bool
	}{
		{
			name: "Expression with no whitespace",
			expr: "5*2",
			want: 10,
		},
		{
			name: "Expression with standard whitespace",
			expr: "5 + 2",
			want: 7,
		},
		{
			name: "Expression with extra whitespace",
			expr: "7  -   2",
			want: 5,
		},
		{
			name: "Division",
			expr: "9 / 3",
			want: 3,
		},
		{
			name:        "Incomplete Expression",
			expr:        "5*",
			errExpected: true,
		},
		{
			name:        "Malformed Expression (postfix)",
			expr:        "5 7 *",
			want:        35,
			errExpected: true,
		},
		{
			name:        "Unknown operation",
			expr:        "5 & 7",
			errExpected: true,
		},
	}
	for _, tc := range tcs {
		got, err := calculator.EvaluateExpression(tc.expr)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("EvaluateExpression(%s): unexpected error status %v in TestCase: %s", tc.expr, err, tc.name)
		}
		expectedResult := closeEnough(tc.want, got, 0.0001)
		if !tc.errExpected && !expectedResult {
			t.Errorf("EvaluateExpression(%s): wanted %f, got %f in TestCase: %s", tc.expr, tc.want, got, tc.name)
		}
	}
}
