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
	distances := calculateDistances(coords)

	fmt.Println(part1(distances, 1000), part2(distances, coords))
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

func part1(distances []distance, top int) int {
	graph := MakeGraph()

	for i := range top {
		graph.ConnectPoints(distances[i])
	}

	components := make([][]int, 0) // this would take up less memory if we only cache the size of the component, not all the members
	for _, v := range graph.Components {
		components = append(components, v)
	}
	slices.SortFunc(components, func(a, b []int) int { return len(b) - len(a) })
	return len(components[0]) * len(components[1]) * len(components[2])
}

func part2(distances []distance, coords []pos) int {
	graph := MakeGraph()

	i := 0
	for i = 0; len(graph.Components[1]) < len(coords); i++ { // loop until we have one component with every coord in it
		graph.ConnectPoints(distances[i])
	}

	lastDistance := distances[i-1] // last connection we made
	return int(coords[lastDistance.a].x * coords[lastDistance.b].x)
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

	slices.SortFunc(distances, func(a, b distance) int { return int(a.dist - b.dist) }) // very important: sort in increasing order of distance
	return distances
}

type graph struct {
	Indices    map[int]int   // index to connected component idx
	Components map[int][]int // connected component to list of indices
	newComp    int           // idx we can use for a brand new component - always incrementing
}

func MakeGraph() *graph {
	g := graph{
		make(map[int]int, 0),
		make(map[int][]int, 0),
		1,
	}

	return &g
}

func (g *graph) ConnectPoints(dist distance) {
	setA := g.Indices[dist.a]
	setB := g.Indices[dist.b]

	if setA > 0 && setB > 0 && setA != setB { // we need to connect two components
		if setB < setA { // setA should be lower - always merge into the lower idx component
			tmp := setA
			setA = setB
			setB = tmp
		}

		g.Components[setA] = append(g.Components[setA], g.Components[setB]...)
		for _, b := range g.Components[setB] {
			g.Indices[b] = setA
		}
		delete(g.Components, setB)

	} else if setA > 0 && setB == 0 { // add one to existing component
		g.Indices[dist.b] = setA
		g.Components[setA] = append(g.Components[setA], dist.b)
	} else if setB > 0 && setA == 0 {
		g.Indices[dist.a] = setB
		g.Components[setB] = append(g.Components[setB], dist.a)
	} else if setA == 0 && setB == 0 { // create new connected component
		g.Components[g.newComp] = []int{dist.a, dist.b}
		g.Indices[dist.a] = g.newComp
		g.Indices[dist.b] = g.newComp
		g.newComp++
	}
}
