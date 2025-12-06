package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputTrimmed := strings.TrimSpace(input)
	lines := strings.Split(inputTrimmed, "\n")

	ranges, ingredients := parseIngredients(lines)
	slices.SortFunc(ranges, func(a, b freshRange) int { return a.low - b.low })

	fmt.Println("Part 1:", part1(ranges, ingredients), "Part 2:", part2(ranges))
}

type freshRange struct {
	low  int
	high int
}

func part1(ranges []freshRange, ingredients []int) int {
	var count int

	combinedRanges := combineRanges(ranges)

	for _, i := range ingredients {
		var found bool
		for _, rn := range combinedRanges {
			if i >= rn.low && i <= rn.high {
				found = true
				break
			}
		}

		if found {
			count++
		}
	}
	return count
}

func part2(ranges []freshRange) int {
	var total int

	combinedRanges := combineRanges(ranges)

	for _, rn := range combinedRanges {
		total += rn.high - rn.low + 1
	}
	return total
}

func combineRanges(ranges []freshRange) []freshRange {
	out := make([]freshRange, 0, 10)
	var rn freshRange

	for i, el := range ranges {
		if i == 0 {
			rn = el
			continue
		}

		if el.low <= rn.high {
			if el.high > rn.high {
				rn.high = el.high
			}
		} else {
			out = append(out, rn)
			rn = el
		}
	}
	return append(out, rn)
}

func parseIngredients(lines []string) ([]freshRange, []int) {
	ranges := make([]freshRange, 0, 10)
	ingredients := make([]int, 0, 10)

	onIngredients := false
	for _, ln := range lines {
		if ln == "" {
			onIngredients = true
			continue
		}

		if !onIngredients {
			splits := strings.Split(ln, "-")
			low, err := strconv.Atoi(strings.TrimSpace(splits[0]))
			if err != nil {
				panic(err)
			}
			high, err := strconv.Atoi(strings.TrimSpace(splits[1]))
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, freshRange{low, high})

		} else {
			i, err := strconv.Atoi(strings.TrimSpace(ln))
			if err != nil {
				panic(err)
			}
			ingredients = append(ingredients, i)
		}
	}
	return ranges, ingredients
}
