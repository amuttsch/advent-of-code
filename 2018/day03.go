package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type claim struct {
	ID     int
	left   int
	top    int
	width  int
	height int
}

var claimList []claim

type squareClaim struct {
	claimIDs []int
}

func listOfStringToListOfInt(ss []string) []int {
	is := make([]int, len(ss))
	for i, s := range ss {
		is[i], _ = strconv.Atoi(s)
	}
	return is
}

func init() {
	dat, err := ioutil.ReadFile("input/day03.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	claimRe := regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")

	claimListStr := strings.Split(string(dat), "\n")

	for _, claimStr := range claimListStr {
		r := claimRe.FindStringSubmatch(claimStr)
		if r != nil && len(r) == 6 {
			claimParams := listOfStringToListOfInt(r[1:])
			claimList = append(claimList, claim{
				ID:     claimParams[0],
				left:   claimParams[1],
				top:    claimParams[2],
				width:  claimParams[3],
				height: claimParams[4],
			})
		}
	}
}

func Day03() (int, int) {
	fmt.Println("Day 03")

	// Create matrix
	fabicMatrix := [1000][1000]squareClaim{}

	for _, claim := range claimList {
		for x := claim.left; x < claim.left+claim.width; x++ {
			for y := claim.top; y < claim.top+claim.height; y++ {
				fabicMatrix[x][y].claimIDs = append(fabicMatrix[x][y].claimIDs, claim.ID)
			}
		}
	}

	overlappingSquares := 0
	nonOverlappingClaims := make(map[int]bool, 0)
	for x, a := range fabicMatrix {
		for y := range a {
			if len(fabicMatrix[x][y].claimIDs) > 1 {
				overlappingSquares++
				for _, c := range fabicMatrix[x][y].claimIDs {
					nonOverlappingClaims[c] = false
				}
			} else if len(fabicMatrix[x][y].claimIDs) == 1 {
				if _, ok := nonOverlappingClaims[fabicMatrix[x][y].claimIDs[0]]; !ok {
					nonOverlappingClaims[fabicMatrix[x][y].claimIDs[0]] = true
				}
			}
		}
	}

	nonOverlappingClaim := 0
	for k, v := range nonOverlappingClaims {
		if v {
			nonOverlappingClaim = k
			break
		}
	}

	fmt.Println("Part 1: Found", overlappingSquares, "overlapping squares")
	fmt.Println("Part 2: Non overlapping claim has ID", nonOverlappingClaim)

	return overlappingSquares, nonOverlappingClaim
}
