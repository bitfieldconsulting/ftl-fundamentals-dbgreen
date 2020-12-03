package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
	"time"
)

const testcount = 100
const testcaseMaxsize = 10

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
		nums := make([]float64, r.Intn(testcaseMaxsize))
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
		nums []float64
		want float64
	}{
		{name: "Negative number multiplied with a positive results in negative", nums: []float64{-1, 5}, want: -5},
		{name: "Two negative numbers multiplied result in a postive", nums: []float64{-3, -9}, want: 27},
		{name: "Number multiplied by zero results in zero", nums: []float64{5, 0}, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.nums...)
		if tc.want != got {
			t.Errorf("Multiply(%f): want %f, got %f in TestCase: %s", tc.nums, tc.want, got, tc.name)
		}
	}
}
func TestDivide(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		nums        []float64
		want        float64
		errExpected bool
	}{
		{name: "Number divided by zero results in an error", nums: []float64{6, 0}, want: 0, errExpected: true},
		{name: "Number divided by another number results in an integer result", nums: []float64{6, 3}, want: 2},
		{name: "Number divided by another number results in a rational result", nums: []float64{10, 4}, want: 2.5},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.nums...)

		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%f): unexpected error status %v in TestCase: %s", tc.nums, err, tc.name)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f): wanted %f, got %f in TestCase: %s", tc.nums, tc.want, got, tc.name)
		}

	}
}
func TestSqrt(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}{
		{name: "Sqrt of negative results in an error", a: -1, want: 0, errExpected: true},
		{name: "Sqrt of zero results in a zero", a: 0, want: 0},
		{name: "Sqrt of positive number results in positive number", a: 16, want: 4},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)

		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("Sqrt(%f): unexpected error status %v in TestCase: %s", tc.a, err, tc.name)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("Sqrt(%f): wanted %f, got %f in TestCase: %s", tc.a, tc.want, got, tc.name)
		}

	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	var testCases = []struct {
		name string
		oper string
		want float64
	}{
		{name: "Multiplication", oper: "5*2", want: 10},
		{name: "Addition", oper: "5 + 2", want: 7},
		{name: "Subtraction", oper: "7  -   2", want: 5},
		{name: "Division", oper: "9 / 3", want: 3},
	}
	for _, tc := range testCases {
		got, _ := calculator.Parse(tc.oper)
		if tc.want != got {
			t.Errorf("Parse(%s): wanted %f, got %f", tc.oper, tc.want, got)
		}
	}
}
