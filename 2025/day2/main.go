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
	ranges := strings.Split(inputTrimmed, ",")

	fmt.Println("Part 1:", sumInvalid(ranges, sameSequenceTwice), "Part 2:", sumInvalid(ranges, anySequenceRepeated))
}

func sumInvalid(ranges []string, validFunc func(int) bool) int {
	var sum int
	for _, rng := range ranges {
		bounds := strings.Split(rng, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		var end int
		end, err = strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			if validFunc(i) {
				sum += i
			}
		}
	}

	return sum
}

// part 2
func sameSequenceTwice(id int) bool {
	str := strconv.Itoa(id)
	if len(str)%2 == 1 {
		return false
	}
	mid := len(str) / 2

	return str[:mid] == str[mid:]
}

// part 2
func anySequenceRepeated(id int) bool {
	str := strconv.Itoa(id)
	for i := 2; i <= len(str); i++ {
		if len(str)%i > 0 {
			continue
		}
		sequence := str[:len(str)/i]
		if str == strings.Repeat(sequence, i) {
			return true
		}
	}
	return false
}
