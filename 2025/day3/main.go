package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputTrimmed := strings.TrimSpace(input)
	lines := strings.Split(inputTrimmed, "\n")

	fmt.Println("Part 1:", sumJolts(lines, 2), "Part 2:", sumJolts(lines, 12))
}

func sumJolts(input []string, n int) int {
	var sum int
	for _, line := range input {
		sum += maxJolt(strings.TrimRight(line, "\r\n"), n)
	}
	return sum
}

func maxJolt(batteries string, n int) int {
	joltString := make([]rune, 0, n)

	var start = 0
	for i := 0; i < n; i++ {
		end := len(batteries) - (n - i - 1) // eg: if n is 12, we need to leave room for 11 characters

		maxRune, subIdx := firstMaxRune(batteries[start:end])
		start += subIdx + 1

		joltString = append(joltString, maxRune)
	}

	out, err := strconv.Atoi(string(joltString))
	if err != nil {
		panic(err)
	}

	return out
}

func firstMaxRune(s string) (rune, int) {
	var max rune
	var idx int
	for i, char := range s {
		if char > max {
			max = char
			idx = i
		}
	}
	return max, idx
}
