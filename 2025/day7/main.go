package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputTrimmed := strings.TrimSpace(input)
	lines := strings.Split(inputTrimmed, "\n")

	fmt.Println("Part 1:", part1(lines), "Part 2:", part2(lines))
}

func part1(lines []string) int {
	var splitCount int
	startPos := strings.Index(lines[0], "S")
	set := map[int]bool{startPos: true}

	for _, line := range lines[1:] {
		for key := range set {
			if line[key] == '^' {
				delete(set, key)
				set[key+1] = true
				set[key-1] = true
				splitCount++
			}
		}
	}

	return splitCount
}

var cachedPaths map[string]int

func part2(lines []string) int {
	cachedPaths = make(map[string]int)
	startPos := strings.Index(lines[0], "S")
	return countTimelines(lines, startPos, 0)
}

func countTimelines(lines []string, x, y int) int {
	atY := y
	for notFound := true; notFound; notFound = (lines[atY][x] != '^') {
		atY++
		if atY >= len(lines) {
			return 1
		}
	}

	loc := fmt.Sprintf("%d,%d", x, atY)
	if match := cachedPaths[loc]; match > 0 {
		return match
	} else {
		timelines := countTimelines(lines, x-1, atY) + countTimelines(lines, x+1, atY)
		cachedPaths[loc] = timelines
		return timelines
	}
}
