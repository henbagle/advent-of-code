package main

import "testing"

func TestPart1(t *testing.T) {
	example := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	answer := 357

	got := sumJolts(example, 2)
	if got != answer {
		t.Errorf("Part1()=%d incorrect", got)
	}
}

func TestPart2(t *testing.T) {
	example := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}
	answer := 3121910778619

	got := sumJolts(example, 12)
	if got != answer {
		t.Errorf("Part2()=%d incorrect", got)
	}
}

func TestMaxJolt2(t *testing.T) {
	tests := []struct {
		batteries string
		expected  int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
		{"111111", 11},
		{"999999", 99},
		{"123456789", 89},
		{"98768111", 98},
	}

	for _, tt := range tests {
		out := maxJolt(tt.batteries, 2)

		if out != tt.expected {
			t.Errorf("MaxJolt(%s) = %d, expected %d", tt.batteries, out, tt.expected)
		}
	}
}

func TestMaxJolt12(t *testing.T) {
	tests := []struct {
		batteries string
		expected  int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tt := range tests {
		out := maxJolt(tt.batteries, 12)
		if out != tt.expected {
			t.Errorf("MaxJolt(%s) = %d, expected %d", tt.batteries, out, tt.expected)
		}
	}
}
