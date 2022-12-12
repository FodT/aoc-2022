package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Node struct {
	Height     int
	Id         string
	Neighbours []*Node
	Visited    bool
	Distance   uint64
}

func main() {

	in := strings.Split(Input, "\n")

	gridH := len(in)
	gridW := len(in[0])
	var startX, startY, endX, endY int
	var lowPoints []*Node
	g := make([][]Node, gridH)
	for i := range g {
		g[i] = make([]Node, gridW)
	}

	for y, row := range in {
		for x, c := range row {
			h := 0
			if c == 'S' {
				h = 0
				startX = x
				startY = y
			} else if c == 'E' {
				h = 25
				endX = x
				endY = y
			} else {
				h = int(c - 97)
			}
			g[y][x].Height = h
			g[y][x].Visited = false
			g[y][x].Id = fmt.Sprintf("%d,%d", x, y)
			if g[y][x].Height == 0 {
				lowPoints = append(lowPoints, &g[y][h])
			}
		}
	}

	for y, _ := range g {
		for x, _ := range g[y] {
			g[y][x].Distance = math.MaxUint64
			if y > 0 {
				if g[y][x].Height >= g[y-1][x].Height-1 {
					g[y][x].Neighbours = append(g[y][x].Neighbours, &g[y-1][x])
				}
			}
			if y < len(g)-1 {
				if g[y][x].Height >= g[y+1][x].Height-1 {
					g[y][x].Neighbours = append(g[y][x].Neighbours, &g[y+1][x])
				}
			}
			if x > 0 {
				if g[y][x].Height >= g[y][x-1].Height-1 {
					g[y][x].Neighbours = append(g[y][x].Neighbours, &g[y][x-1])
				}
			}
			if x < len(g[y])-1 {
				if g[y][x].Height >= g[y][x+1].Height-1 {
					g[y][x].Neighbours = append(g[y][x].Neighbours, &g[y][x+1])
				}
			}

		}
	}

	fmt.Printf("Start: %d,%d\nEnd:%d,%d\n", startX, startY, endX, endY)
	start := &g[startY][startX]
	end := &g[endY][endX]

	search(start, end)
	fmt.Printf("Part 1 Distance: %d\n", end.Distance)

	var distances []uint64
	for _, lowPoint := range lowPoints {
		// reset distances
		for y, _ := range g {
			for x, _ := range g[y] {
				g[y][x].Distance = math.MaxUint64
				g[y][x].Visited = false
			}
		}
		search(lowPoint, end)
		distances = append(distances, end.Distance)
	}

	sort.Slice(distances[:], func(i, j int) bool {
		return distances[i] < distances[j]
	})
	fmt.Printf("shortest hike: %d\n", distances[0])
}

func search(start *Node, destination *Node) {
	currentNode := start
	start.Distance = 0
	unvisited := make([]*Node, 0)
	for ok := true; ok; ok = !destination.Visited {
		if !currentNode.Visited {
			for i, _ := range currentNode.Neighbours {
				stepDistance := currentNode.Distance + 1

				if currentNode.Neighbours[i].Visited {
					continue
				}
				if currentNode.Neighbours[i] != currentNode && !currentNode.Neighbours[i].Visited {
					unvisited = append(unvisited, currentNode.Neighbours[i])
				}
				if currentNode.Neighbours[i].Distance > stepDistance {
					currentNode.Neighbours[i].Distance = stepDistance
				}
			}

			currentNode.Visited = true
		}

		sort.Slice(unvisited[:], func(i, j int) bool {
			return unvisited[i].Distance < unvisited[j].Distance
		})

		if len(unvisited) > 0 {
			currentNode = unvisited[0]
			unvisited = unvisited[1:]
		} else {
			break
		}
	}

}
