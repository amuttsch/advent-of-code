package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var freqChanges []int

func init() {
	fmt.Println("init Day01")
	dat, err := ioutil.ReadFile("input/day01.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	datStr := string(dat)
	freqChangesStr := strings.Split(datStr, "\n")
	freqChanges = make([]int, 0, len(freqChangesStr))

	for _, fcs := range freqChangesStr {
		if fcInt, err := strconv.Atoi(fcs); err == nil {
			freqChanges = append(freqChanges, fcInt)
		}
	}
}

func Day01Part1() int {
	fmt.Println("Day 01 - Part 1")

	calibratedFrequency := 0

	for _, fc := range freqChanges {
		calibratedFrequency += fc
	}

	fmt.Println("  Calibrated frequency:", calibratedFrequency)

	return calibratedFrequency
}

func Day01Part2() int {
	fmt.Println("Day 01 - Part 2")

	calibratedFrequency := 0
	seenFrequencies := make(map[int]bool)
	foundFrequencyTwice := false
	seenFrequencies[0] = false

	for !foundFrequencyTwice {
		for _, fc := range freqChanges {
			calibratedFrequency += fc

			_, foundFrequencyTwice = seenFrequencies[calibratedFrequency]
			if foundFrequencyTwice {
				break
			}
			seenFrequencies[calibratedFrequency] = false
		}
	}

	fmt.Println("  First frequency found twice:", calibratedFrequency)

	return calibratedFrequency
}
