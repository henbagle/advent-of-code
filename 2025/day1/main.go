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

	fmt.Println("Part 1:", part1(lines), "Part 2:", part2(lines))
}

func part1(cmds []string) int {
	dial := 50
	var zeroes int

	for _, cmd := range cmds {
		dial, _ = rot(dial, cmd)
		if dial == 0 {
			zeroes++
		}
	}
	return zeroes
}

func part2(cmds []string) int {
	dial := 50
	var allZeroes int

	for _, cmd := range cmds {
		var zeroes int
		dial, zeroes = rot(dial, cmd)
		allZeroes += zeroes
	}
	return allZeroes
}

func rot(startPos int, cmd string) (outPos int, zeroCount int) {
	amount, err := strconv.Atoi(cmd[1:])
	if err != nil {
		panic(err)
	}

	var out int
	dir := cmd[:1]
	switch dir {
	case "R":
		out = startPos + amount%100
	case "L":
		out = startPos - amount%100
	default:
		panic(fmt.Sprintf("unknown direction '%s'", dir))
	}

	zeroes := amount / 100 // number of times we passed zero spinning
	if out < 0 {
		out += 100
		if startPos > 0 {
			zeroes++
		}
	}
	if out > 99 {
		out -= 100
		zeroes++
	}

	// edge case when moving left back to zero
	if dir == "L" && amount%100 > 0 && out == 0 {
		zeroes++
	}

	return out, zeroes
}
