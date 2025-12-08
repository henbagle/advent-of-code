package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputTrimmed := strings.TrimSpace(input)
	lines := strings.Split(inputTrimmed, "\n")
	coords := parseCoords(lines)

	fmt.Println(part1(coords, 1000))
}

type pos struct {
	x float64
	y float64
	z float64
}

func parseCoords(lines []string) []pos {
	results := make([]pos, len(lines))
	for i, str := range lines {
		split := strings.Split(strings.TrimRight(str, "\n"), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		results[i] = pos{
			float64(x),
			float64(y),
			float64(z),
		}
	}
	return results
}

func part1(coords []pos, top int) int {
	set := make(map[int]int, 0)
	counts := make([]int, 0)
	distances := calculateDistances(coords)

	for i := range top {
		dist := distances[i]
		setA := set[dist.a]
		setB := set[dist.b]
		if setA > 0 && setB > 0 && setA != setB {
			if setB < setA { // setA should be lower
				tmp := setA
				setA = setB
				setB = tmp
			}

			counts[setA] += counts[setB]
			counts[setB] = 0
			for k, v := range set {
				if v == setB {
					set[k] = setA
				}
			}

		} else if setA > 0 && setB == 0 {
			set[dist.b] = set[dist.a]
			counts[setA]++
		} else if setB > 0 && setA == 0 {
			set[dist.a] = set[dist.b]
			counts[setB]++
		} else if setA == 0 && setB == 0 {
			counts = append(counts, 2)
			set[dist.a] = len(counts) - 1
			set[dist.b] = len(counts) - 1
		}
	}
	slices.SortFunc(counts, func(a, b int) int { return b - a })
	return counts[0] * counts[1] * counts[2]
}

type distance struct {
	dist float64
	a    int
	b    int
}

func calculateDistances(coords []pos) []distance {
	distances := make([]distance, 0, len(coords)*len(coords))

	for ai, a := range coords {
		for bi, b := range coords[ai+1:] {
			distances = append(distances, distance{
				math.Pow((b.x-a.x), 2) + math.Pow((b.y-a.y), 2) + math.Pow((b.z-a.z), 2),
				ai,
				ai + 1 + bi,
			})
		}
	}

	slices.SortFunc(distances, func(a, b distance) int { return int((a.dist - b.dist) * 100.0) })
	return distances
}
