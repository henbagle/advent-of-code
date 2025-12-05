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

func sumInvalid(ranges []string, invalidFunc func(int) bool) int {
	var sum int
	for _, r := range ranges {
		bounds := strings.Split(r, "-")
		start, err := strconv.Atoi(bounds[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(bounds[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			if invalidFunc(i) {
				sum += i
			}
		}
	}
	return sum
}

// part 1
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
		if len(str)%i > 0 { // skip if not evenly divisible by string length
			continue
		}
		sequence := str[:len(str)/i]
		if str == strings.Repeat(sequence, i) {
			return true
		}
	}
	return false
}
