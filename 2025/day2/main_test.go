package main

import "testing"

func TestSameSequenceTwice(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{11, true},
		{21, false},
		{22, true},
		{101, false},
		{1010, true},
		{1698522, false},
		{446446, true},
		{38593859, true},
		{1012, false},
	}

	for _, tt := range tests {
		valid := sameSequenceTwice(tt.input)

		if valid != tt.expected {
			t.Errorf("isIdInvalid(%d) = %t, expected %t", tt.input, valid, tt.expected)
		}
	}
}

func TestAnySequenceRepeated(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{11, true},
		{21, false},
		{22, true},
		{101, false},
		{1010, true},
		{565656, true},
		{5656561, false},
		{824824824, true},
		{38593859, true},
		{1012, false},
		{2121212121, true},
		{222, true},
		{223, false},
	}

	for _, tt := range tests {
		valid := anySequenceRepeated(tt.input)

		if valid != tt.expected {
			t.Errorf("isIdInvalid(%d) = %t, expected %t", tt.input, valid, tt.expected)
		}
	}
}
