package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strings"
)

type target struct {
	ID    int
	pt    image.Point
	isInf bool
	area  int
}

var targets []*target

type gridElem struct {
	pt     image.Point
	unique bool
	dist   int
	target *target
}

func init() {
	dat, err := ioutil.ReadFile("input/day06.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	ptStr := strings.Split(string(dat), "\n")

	for i, pt := range ptStr {
		if pt != "" {
			var p image.Point
			fmt.Sscanf(pt, "%d, %d", &p.X, &p.Y)
			targets = append(targets, &target{pt: p, ID: i})
		}
	}
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func manhattenDist(p1, p2 image.Point) int {
	p := p1.Sub(p2)
	return abs(p.X) + abs(p.Y)
}

func Day06() (int, int) {
	fmt.Println("Day 06")

	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, 0, 0

	for _, t := range targets {
		if t.pt.X < minX {
			minX = t.pt.X
		}
		if t.pt.X > maxX {
			maxX = t.pt.X
		}
		if t.pt.Y < minY {
			minY = t.pt.Y
		}
		if t.pt.Y > maxY {
			maxY = t.pt.Y
		}
	}

	grid := make(map[int]map[int]*gridElem, maxX+1)
	regionWithMinDistanceToAllTargets := 0

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			absoluteDistanceToThisPoint := 0

			for _, t := range targets {
				dist := manhattenDist(t.pt, image.Point{x, y})
				absoluteDistanceToThisPoint += abs(x-t.pt.X) + abs(y-t.pt.Y)
				if _, ok := grid[x]; !ok {
					grid[x] = make(map[int]*gridElem, maxY+1)
				}

				if ge, ok := grid[x][y]; ok {
					switch {
					case ge.dist > dist:
						ge.unique = true
						ge.dist = dist
						ge.target = t
					case ge.dist == dist:
						ge.unique = false
					case ge.dist < dist:
						// Do nothing
					}
				} else {
					grid[x][y] = &gridElem{
						pt:     image.Point{x, y},
						unique: true,
						dist:   dist,
						target: t,
					}
				}
			}

			if absoluteDistanceToThisPoint < 10000 {
				regionWithMinDistanceToAllTargets++
			}
		}
	}

	// Count the area for each target
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y].unique {
				grid[x][y].target.area++
			}
		}
	}

	// Get the largest area
	largestArea := 0
	for _, t := range targets {
		// Do not take infinite areas into account
		if grid[minX][t.pt.Y].target == t {
			continue
		}
		if grid[maxX][t.pt.Y].target == t {
			continue
		}
		if grid[t.pt.X][minY].target == t {
			continue
		}
		if grid[t.pt.X][maxY].target == t {
			continue
		}

		if t.area > largestArea {
			largestArea = t.area
		}
	}

	fmt.Println("Part 1: Largest area without infinite ones", largestArea)
	fmt.Println("Part 2: Size of region within 10000 is", regionWithMinDistanceToAllTargets)

	return largestArea, regionWithMinDistanceToAllTargets
}
