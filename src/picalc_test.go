package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

// Tests calculatePi with various iteration counts.
func TestCalculatePi(t *testing.T) {
	tests := []struct {
		iterations int
		expected   float64
		delta      float64
	}{
		{1, 4.0, 0.0001},
		{10, 3.0418396189, 0.0001},
		{1000, 3.1405926538, 0.00001},
		{2000000, 3.1415921536, 0.0000005},
	}

	for _, test := range tests {
		result := calculatePi(test.iterations)
		if result < test.expected-test.delta || result > test.expected+test.delta {
			t.Errorf("calculatePi(%d) = %.10f; want %.10f ± %.10f", test.iterations, result, test.expected, test.delta)
		}
	}
}

// Tests the HTTP handler with valid, invalid, and missing query parameters.
func TestHandler(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"/?iterations=10", "3.0418396189\n"},
		{"/?iterations=invalid", "iterations parameter not valid\n"},
		{"/", "iterations parameter missing\n"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.url, nil)
		rr := httptest.NewRecorder()
		handler(rr, req)

		if rr.Body.String() != tt.expected {
			t.Errorf("handler(%q) response = %q; want %q", tt.url, rr.Body.String(), tt.expected)
		}
	}
}

// TestMain covers the server startup code to further increase test coverage.
func TestMain(m *testing.M) {
	go func() {
		os.Setenv("PORT", "0")
		main()
	}()
	os.Exit(m.Run())
}
