package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var boxList []string

func init() {
	fmt.Println("init Day02")
	dat, err := ioutil.ReadFile("input/day02.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	boxList = strings.Split(string(dat), "\n")
}

func Day02Part1() int {
	fmt.Println("Day 02 - Part 1")

	numberBoxesWithTwoLetters := 0
	numberBoxesWithThreeLetters := 0

	for _, boxID := range boxList {
		charMap := make(map[rune]int)
		for _, r := range boxID {
			charCount := charMap[r]
			charMap[r] = charCount + 1
		}
		hasTwoLetter := false
		hasThreeLetter := false
		for _, v := range charMap {
			if v == 2 {
				hasTwoLetter = true
			}
			if v == 3 {
				hasThreeLetter = true
			}
		}
		if hasTwoLetter {
			numberBoxesWithTwoLetters++
		}
		if hasThreeLetter {
			numberBoxesWithThreeLetters++
		}
	}

	result := numberBoxesWithTwoLetters * numberBoxesWithThreeLetters
	fmt.Println("numberBoxesWithTwoLetters:", numberBoxesWithTwoLetters, " * numberBoxesWithThreeLetters:", numberBoxesWithThreeLetters, " = ", result)

	return result
}

type boxCombination struct {
	b1 string
	b2 string
}

func getCombinations(boxIds []string, outChan chan<- boxCombination) {
	if len(boxIds) == 0 || boxIds[0] == "" {
		close(outChan)
		return
	}

	firstID := boxIds[0]

	for _, boxID := range boxIds[1:] {
		if boxID != "" {
			outChan <- boxCombination{b1: firstID, b2: boxID}
		}
	}

	getCombinations(boxIds[1:], outChan)
}

func Day02Part2() string {
	fmt.Println("Day 02 - Part 2")

	combOutputChannel := make(chan boxCombination)

	go getCombinations(boxList, combOutputChannel)

	for bc := range combOutputChannel {
		if len(bc.b1) != len(bc.b2) {
			fmt.Println("Length of inputs differ. This should not occur!")
			continue
		}

		var correctBoxID strings.Builder
		for i, r := range bc.b1 {
			if r == rune(bc.b2[i]) {
				correctBoxID.WriteRune(r)
			}
		}

		diff := len(bc.b1) - correctBoxID.Len()
		if diff == 1 {
			fmt.Println("Found boxid", correctBoxID.String())
			return correctBoxID.String()
		}
	}

	return ""
}
