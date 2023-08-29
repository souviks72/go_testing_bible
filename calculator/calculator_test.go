package calculator_test

import (
	"testing"

	"github.com/souviks72/go_testing_bible/calculator"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("Testing 3 digit Armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{
				name:     "Testing for 153",
				value:    153,
				expected: true,
			},
			{
				name:     "Testing for 351",
				value:    351,
				expected: true,
			},
			{
				name:     "Testing for 350",
				value:    350,
				expected: true,
			},
			{
				name:     "Testing for 470",
				value:    470,
				expected: true,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.actual = calculator.CalculateIsArmstrong(test.value)
				if test.actual != test.expected {
					t.Fail()
				}
			})
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("Testing negative for Armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			{
				name:     "Testing for 152",
				value:    152,
				expected: false,
			},
			{
				name:     "Testing for 251",
				value:    251,
				expected: false,
			},
			{
				name:     "Testing for 250",
				value:    250,
				expected: false,
			},
			{
				name:     "Testing for 270",
				value:    270,
				expected: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.actual = calculator.CalculateIsArmstrong(test.value)
				if test.actual != test.expected {
					t.Fail()
				}
			})
		}
	})
}

/*
Test Coverage-------
go test ./ --cover  -> will show percentage of code covered in tests

To see proper report in HTML file:
go test ./ -coverprofile=coverage.out
go tool cover -html=coverage.out
*/

// running benchmarks: go test ./calculator -run=Benchmark -bench=.
func benchmarkCalculateIsArmstrong(input int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.CalculateIsArmstrong(input)
	}
}
func BenchmarkCalculateIsAmrstrong370(b *testing.B) {
	benchmarkCalculateIsArmstrong(370, b)
}

func BenchmarkCalculateIsAmrstrong470(b *testing.B) {
	benchmarkCalculateIsArmstrong(470, b)
}

func BenchmarkCalculateIsAmrstrong371(b *testing.B) {
	benchmarkCalculateIsArmstrong(371, b)
}
