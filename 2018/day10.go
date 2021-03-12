package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

var points []*point

func init() {
	dat, err := ioutil.ReadFile("input/day10.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	pointsStr := strings.Split(string(dat), "\n")

	for _, pStr := range pointsStr {
		if pStr != "" {
			var p point
			fmt.Sscanf(pStr, "position=<%d, %d> velocity=<%d, %d>", &p.x, &p.y, &p.vx, &p.vy)
			points = append(points, &p)
		}
	}
}

func Day10() (string, int) {
	fmt.Println("Day 10")

	minYDist := math.MaxInt64
	secondsElapsed := 0

	for true {
		minY := math.MaxInt64
		maxY := 0
		lastDist := minYDist
		for _, p := range points {
			p.x += p.vx
			p.y += p.vy

			if p.y > maxY {
				maxY = p.y
			}
			if p.y < minY {
				minY = p.y
			}
		}
		minYDist = maxY - minY
		if minYDist > lastDist {
			break
		}
		secondsElapsed++
	}

	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	col := color.RGBA{255, 255, 255, 255}
	for _, p := range points {
		//fmt.Println(p.x, p.y)
		p.x -= p.vx
		p.y -= p.vy
		img.Set(p.x, p.y, col)
	}
	myfile, err := os.Create("day10.png")
	if err != nil {
		panic(err.Error())
	}
	defer myfile.Close()
	png.Encode(myfile, img)

	fmt.Println("Part 1:", "XPFXXXKL")
	fmt.Println("Part 2:", secondsElapsed)

	return "XPFXXXKL", secondsElapsed
}
