package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	numbers, operators := parseProblems(lines)

	fmt.Println("Part 1:", part1(numbers, operators), "Part 2:", part2(lines, operators))
}

func part1(numbers [][]string, operators []string) int {
	var totalSum int

	for i, op := range operators {
		var result int
		isPlus := op == "+"
		if !isPlus {
			result = 1
		}

		for _, arr := range numbers {
			n, err := strconv.Atoi(arr[i])
			if err != nil {
				panic(err)
			}

			if isPlus {
				result += n
			} else {
				result *= n
			}
		}
		totalSum += result
	}
	return totalSum
}

func part2(lines []string, operators []string) int {
	numberLines := lines[:len(lines)-1]
	numbers := cephalopodNumbers(numberLines)

	var totalSum int

	for i, op := range operators {
		var result int
		isPlus := op == "+"
		if !isPlus {
			result = 1
		}

		for _, str := range numbers[i] {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}

			if isPlus {
				result += n
			} else {
				result *= n
			}
		}
		totalSum += result
	}
	return totalSum
}

func cephalopodNumbers(lines []string) [][]string {
	numbers := make([][]string, 0, len(lines[0])/5*4)
	problem := make([]string, 0, 1)
	for i := range lines[0] {
		str := make([]byte, 0)
		for j := range lines {
			r := lines[j][i]
			if r != ' ' {
				str = append(str, r)
			}
		}

		if len(str) == 0 {
			numbers = append(numbers, problem)
			problem = make([]string, 0, 1)
		} else {
			problem = append(problem, string(str))
		}
	}
	numbers = append(numbers, problem)
	return numbers
}

func parseProblems(lines []string) ([][]string, []string) {
	numbers := make([][]string, len(lines)-1)
	var operators []string
	for i, line := range lines {
		if i < len(lines)-1 {
			numbers[i] = strings.Split(cleanupLine(line), " ")
		} else {
			operators = strings.Split(cleanupLine(line), " ")
		}
	}
	return numbers, operators
}

func cleanupLine(line string) string {
	stripped := strings.TrimSpace(line)
	r, _ := regexp.Compile("[ ]+")
	return r.ReplaceAllString(stripped, " ")
}
