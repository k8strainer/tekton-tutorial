package main

import (
	"math"
	"testing"
)

func TestCalculatePi(t *testing.T) {
	cases := []struct {
		iterations int
		expected   float64
		tolerance  float64
	}{
		{1, 4.0, 0.0000000001},
		{10, 3.0418396189, 0.0000000001},
		{100, 3.1315929036, 0.0000000001},
		{1000, 3.1405926538, 0.0000000001},
		{10000, 3.1414926536, 0.0000000001},
		{100000, 3.1415826536, 0.0000000001},
		{1000000, 3.1415916536, 0.0000000001},
		{20000000, math.Pi, 0.0000001}, // High-precision test
	}

	for _, c := range cases {
		result := calculatePi(c.iterations)
		if math.Abs(result-c.expected) > c.tolerance {
			t.Errorf("calculatePi(%d) = %.10f, expected approximately %.10f", c.iterations, result, c.expected)
		}
	}
}
