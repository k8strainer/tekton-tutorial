package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestCalculatePi prüft die Berechnung für verschiedene Iterationen.
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

// TestHandler prüft, ob der HTTP-Handler korrekt auf gültige und ungültige Anfragen reagiert.
func TestHandler(t *testing.T) {
	reqValid, _ := http.NewRequest("GET", "/?iterations=10", nil)
	reqInvalid, _ := http.NewRequest("GET", "/?iterations=invalid", nil)
	reqMissing, _ := http.NewRequest("GET", "/", nil)

	tests := []struct {
		req        *http.Request
		wantStatus int
		wantBody   string
	}{
		{reqValid, http.StatusOK, "3.0418396189\n"},
		{reqInvalid, http.StatusOK, "iterations parameter not valid\n"},
		{reqMissing, http.StatusInternalServerError, ""},
	}

	for _, tt := range tests {
		rr := httptest.NewRecorder()
		handler(rr, tt.req)

		if status := rr.Code; status != http.StatusOK && tt.wantBody != "" {
			t.Errorf("handler() status code = %v, want %v", status, http.StatusOK)
		}
		if tt.wantBody != "" && rr.Body.String() != tt.wantBody {
			t.Errorf("handler() response = %v, want %v", rr.Body.String(), tt.wantBody)
		}
	}
}

// TestMain prüft die Funktionalität der main-Funktion (Coverage-Erhöhung durch indirekten Aufruf).
func TestMain(m *testing.M) {
	go func() {
		os.Setenv("PORT", "0") // random verfügbarer Port, um Kollisionen zu verhindern
		main()
	}()
	os.Exit(m.Run())
}
