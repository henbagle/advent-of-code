package main

import (
	"strings"
	"testing"
)

var testInput = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func TestPart1(t *testing.T) {
	expected := 21

	inputTrimmed := strings.TrimSpace(testInput)
	lines := strings.Split(inputTrimmed, "\n")
	result := part1(lines)

	if result != expected {
		t.Errorf("part1() returned %d, expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 40

	inputTrimmed := strings.TrimSpace(testInput)
	lines := strings.Split(inputTrimmed, "\n")
	result := part2(lines)

	if result != expected {
		t.Errorf("part2() returned %d, expected %d", result, expected)
	}
}
