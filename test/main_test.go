package main

import "testing"

// tabla de pruebas
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"validate-data", 10.0, 1.0, 10.0, false},
	{"invalidate-data", 10.0, 0, 0, true},
	{"expect-5", 50.0, 10, 5.0, false},
}

func TestMain(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isError {
			if err == nil {
				t.Error("Expected error, got nil")
			}
		} else {
			if err != nil {
				t.Error("Expected nil, got error")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f, got %f", tt.expected, got)
		}
	}
}

// unit test
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Expected error, got nil")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
