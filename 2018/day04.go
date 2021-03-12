package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guardRecord struct {
	ts      time.Time
	guardID int
	action  string
}

var guardRecords []guardRecord

func init() {
	dat, err := ioutil.ReadFile("input/day04.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	guardRe := regexp.MustCompile("\\[(\\d+-\\d+-\\d+ \\d+:\\d+)\\] (\\w+) #?(\\d+)?")

	guardListStr := strings.Split(string(dat), "\n")
	sort.Sort(sort.StringSlice(guardListStr))

	guardID := 0
	for _, guardStr := range guardListStr {
		r := guardRe.FindStringSubmatch(guardStr)
		if len(r) != 4 {
			continue
		}

		ts, _ := time.Parse("2006-01-02 15:04", r[1])
		if newGuardID, err := strconv.Atoi(r[3]); err == nil {
			guardID = newGuardID
		}

		guardRecords = append(guardRecords, guardRecord{
			ts:      ts,
			guardID: guardID,
			action:  r[2],
		})
	}
}

func Day04() (int, int) {
	fmt.Println("Day 04")

	guardMap := make(map[int][60]int, 0)

	fallsAsleepTime := time.Time{}
	for _, gr := range guardRecords {
		switch gr.action {
		case "falls":
			fallsAsleepTime = gr.ts
		case "wakes":
			minutes := guardMap[gr.guardID]
			for m := fallsAsleepTime.Minute(); m < gr.ts.Minute(); m++ {
				minutes[m]++
			}
			guardMap[gr.guardID] = minutes
		default:
		}
	}

	guardMostlyAsleep := 0
	guardMostlyAsleepAtMinute := 0
	totalMinutesAsleep := 0
	minuteMostlyAsleep := 0
	minuteMostlyAsleep2 := 0
	timesAsleepAtMinute := 0
	for id, mins := range guardMap {
		tms := 0
		mma := 0
		taas := 0
		for i, m := range mins {
			tms += m
			if m > taas {
				mma = i
				taas = m
			}
		}
		if tms > totalMinutesAsleep {
			guardMostlyAsleep = id
			totalMinutesAsleep = tms
			minuteMostlyAsleep = mma
		}
		if taas > timesAsleepAtMinute {
			minuteMostlyAsleep2 = mma
			guardMostlyAsleepAtMinute = id
		}
	}

	resultPart1 := guardMostlyAsleep * minuteMostlyAsleep
	resultPart2 := guardMostlyAsleepAtMinute * minuteMostlyAsleep2
	fmt.Println("Part 1:", resultPart1)
	fmt.Println("Part 2:", resultPart2)

	return resultPart1, resultPart2
}
