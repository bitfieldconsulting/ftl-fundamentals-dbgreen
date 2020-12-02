package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

const testcount = 100
const testcase_maxsize = 10

type testCases struct {
	name string
	nums []float64
	want float64
}

func TestAddRand(t *testing.T) {
	t.Parallel()
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	count := 0
	// run the test addcount times
	for count < testcount {
		// get a random size for the number of values
		// to add together
		nums := make([]float64, r.Intn(testcase_maxsize))
		// initialize nums slice to contain random values
		var want float64 = 0
		for i := range nums {
			nums[i] = r.NormFloat64()
			want += nums[i]
		}

		tc := testCases{name: "Random value data", nums: nums, want: want}
		// call Add to get the sum of nums
		got := calculator.Add(tc.nums...)
		// did we get what we expected?
		if tc.want != got {
			t.Errorf("Add (%v): want %f, got %f", tc.nums, tc.want, got)
		}
		count++
	}
}
func TestAdd(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name string
		nums []float64
		want float64
	}{
		{name: "Two numbers that sum to zero", nums: []float64{2, 3, 4}, want: 9},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.nums...)
		if tc.want != got {
			t.Errorf("Add (%v): want %f, got %f with TestCase: %s", tc.nums, tc.want, got, tc.name)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		nums []float64
		want float64
	}{
		{name: "Two numbers that subtract to a negative", nums: []float64{-1, 2}, want: -3},
		{name: "Two numbers that subtract to a positive", nums: []float64{100, 99}, want: 1},
		{name: "Two numbers that subtract to zero", nums: []float64{3, 3}, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.nums...)
		if tc.want != got {
			t.Errorf("Subtract(%v): want %f, got %f with TestCase: %s", tc.nums, tc.want, got, tc.name)
		}
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		a, b float64
		want float64
	}{
		{name: "Negative number multiplied with a positive results in negative", a: -1, b: 5, want: -5},
		{name: "Two negative numbers multiplied result in a postive", a: -9, b: -3, want: 27},
		{name: "Number multiplied by zero results in zero", a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f,%f): want %f, got %f in TestCase: %s", tc.a, tc.b, tc.want, got, tc.name)
		}
	}
}
func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}{
		{name: "Number divided by zero results in an error", a: 6, b: 0, want: 0, errExpected: true},
		{name: "Number divided by another number results in an integer result", a: 6, b: 3, want: 2},
		{name: "Number divided by another number results in a rational result", a: 10, b: 4, want: 2.5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%f,%f): unexpected error status %v in TestCase: %s", tc.a, tc.b, err, tc.name)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide (%f,%f): wanted %f, got %f in TestCase: %s", tc.a, tc.b, tc.want, got, tc.name)
		}

	}
}
