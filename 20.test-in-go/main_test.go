package main

import (
	"testing"
)

var tests = []struct {
	name     string
	dividend float32
	devisor  float32
	expected float32
	isErr    bool
}{
	{"Valid-data", 100.0, 10.0, 10.0, false},
	{"Invalid-data", 100.0, 0.0, 0.0, true},
}

func TestDevide(t *testing.T) {

	for _, tt := range tests {
		got, err := Devide(tt.dividend, tt.devisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected Error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expected Error but got one")
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}
