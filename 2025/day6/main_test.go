package main

import (
	"strings"
	"testing"
)

func TestCleanupLine(t *testing.T) {
	tests := []struct {
		line     string
		expected string
	}{
		{" 0 1    2    123    14       \n", "0 1 2 123 14"},
		{"     *     +    ", "* +"},
		{"1 2 3 4 5", "1 2 3 4 5"},
		{"", ""},
		{"\n", ""},
	}

	for _, tt := range tests {
		out := cleanupLine(tt.line)
		if out != tt.expected {
			t.Errorf("cleanupLine(%s) = %s, expected %s", tt.line, out, tt.expected)
		}
	}
}

func TestParseProblems(t *testing.T) {
	testInput := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
	expectedProblems := 4
	expectedNPerProblem := 3

	inputTrimmed := strings.TrimSpace(testInput)
	lines := strings.Split(inputTrimmed, "\n")
	numbers, operators := parseProblems(lines)
	if len(numbers) != expectedNPerProblem {
		t.Errorf("Expected %d lines of numbers, got %d", expectedNPerProblem, len(numbers))
	}
	if len(operators) != expectedProblems {
		t.Errorf("Expected %d problems, got %d", expectedProblems, len(operators))
	}
	for i, n := range numbers {
		if len(n) != expectedProblems {
			t.Errorf("Expected %d problems on line %d, got %d", expectedProblems, i, len(n))
		}
	}
}

func TestPart1(t *testing.T) {
	testInput := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
	expected := 4277556

	inputTrimmed := strings.TrimSpace(testInput)
	lines := strings.Split(inputTrimmed, "\n")
	numbers, operators := parseProblems(lines)
	result := part1(numbers, operators)

	if result != expected {
		t.Errorf("part1() returned %d, expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	testInput := "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
	expected := 3263827

	lines := strings.Split(testInput, "\n")
	_, operators := parseProblems(lines)
	result := part2(lines, operators)

	if result != expected {
		t.Errorf("part2() returned %d, expected %d", result, expected)
	}
}
