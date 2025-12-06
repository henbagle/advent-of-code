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

	p1, _ := part1(lines)
	fmt.Println("Part 1:", p1, "Part 2:", part2(lines))
}

func part1(rows []string) (int, []string) {
	var count int

	outputRows := make([]string, len(rows))
	for i, row := range rows {
		rowCopy := []rune(row)
		for j, char := range row {
			if char == '@' && forkliftAccessible(rows, j, i) {
				count++
				rowCopy[j] = '.'
			}
		}
		outputRows[i] = string(rowCopy)
	}
	return count, outputRows
}

func part2(rows []string) int {
	var totalCount int

	i := 1
	newRows := rows

	for i != 0 {
		i, newRows = part1(newRows)
		totalCount += i
	}

	return totalCount
}

func forkliftAccessible(rows []string, x int, y int) bool {
	var adjacent int
	for _, pair := range getAdjacentPositions(x, y) {
		if getGridValue(rows, pair[0], pair[1]) {
			adjacent++
			if adjacent >= 4 {
				return false
			}
		}
	}
	return true
}

func getGridValue(rows []string, x int, y int) bool {
	if x < 0 || y < 0 || y >= len(rows) || x >= len(rows[0]) {
		return false
	}
	return rows[y][x] == '@'
}

func getAdjacentPositions(X int, Y int) [8][2]int {
	output := [8][2]int{
		{X - 1, Y - 1},
		{X, Y - 1},
		{X + 1, Y - 1},
		{X - 1, Y},
		{X + 1, Y},
		{X - 1, Y + 1},
		{X, Y + 1},
		{X + 1, Y + 1},
	}

	return output
}
