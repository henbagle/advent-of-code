package main

import "testing"

func TestRot(t *testing.T) {
	tests := []struct {
		instruction string
		starting    int
		expected    int
	}{
		{"R10", 0, 10},
		{"L5", 10, 5},
		{"L1", 0, 99},
		{"R1", 99, 0},
		{"R50", 99, 49},
		{"L200", 30, 30},
	}

	for _, tt := range tests {
		out, _ := rot(tt.starting, tt.instruction)

		if out != tt.expected {
			t.Errorf("rot(%d, %s) = %d, expected %d", tt.starting, tt.instruction, out, tt.expected)
		}
	}
}

func TestRotCountZeroes(t *testing.T) {
	tests := []struct {
		instruction    string
		starting       int
		expectedZeroes int
	}{
		{"L2", 1, 1},
		{"R2", 99, 1},
		{"R100", 50, 1},
		{"R200", 50, 2},
		{"L200", 50, 2},
		{"R1000", 50, 10},
		{"L200", 0, 2},
		{"L1", 1, 1},
		{"R1", 99, 1},
		{"L99", 99, 1},
		{"L199", 99, 2},
		{"R101", 99, 2},
	}

	for _, tt := range tests {
		_, zeroes := rot(tt.starting, tt.instruction)

		if zeroes != tt.expectedZeroes {
			t.Errorf("rot(%d, %s) = %d, expected %d", tt.starting, tt.instruction, zeroes, tt.expectedZeroes)
		}
	}
}

func TestPart1(t *testing.T) {
	example := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	answer := 3

	got := part1(example)
	if got != answer {
		t.Errorf("part1()=%d incorrect", got)
	}
}

func TestPart2(t *testing.T) {
	example := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}
	answer := 6

	got := part2(example)
	if got != answer {
		t.Errorf("part2()=%d incorrect", got)
	}
}

func TestPart2_LandingAndPassingZeroes(t *testing.T) {
	example := []string{
		"L50",
		"R300",
		"L4",
		"R4",
	}
	answer := 5

	got := part2(example)
	if got != answer {
		t.Errorf("part2()=%d incorrect", got)
	}
}
