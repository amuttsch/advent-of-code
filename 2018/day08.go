package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var numberList []int

func init() {
	dat, err := ioutil.ReadFile("input/day08.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	strList := strings.Split(string(dat), "\n")
	strList = strings.Split(strList[0], " ")

	for _, s := range strList {
		if i, err := strconv.Atoi(s); err == nil {
			numberList = append(numberList, i)
		}
	}
}

func extractMetadata(position int) (newPosition int, metadata []int, nodeValue int) {
	newPosition = position
	// Check bounds
	if newPosition >= (len(numberList) - 2) {
		return
	}
	nodeCount := numberList[newPosition]
	metadataCount := numberList[newPosition+1]
	newPosition += 2

	childNodeValues := make([]int, nodeCount)

	for nc := 0; nc < nodeCount; nc++ {
		p, newMetadata, nv := extractMetadata(newPosition)
		childNodeValues[nc] = nv
		newPosition = p
		metadata = append(metadata, newMetadata...)
	}

	metadataSum := 0
	for mc := 0; mc < metadataCount; mc++ {
		metadata = append(metadata, numberList[newPosition])
		metadataSum += numberList[newPosition]

		if numberList[newPosition] <= nodeCount {
			nodeValue += childNodeValues[numberList[newPosition]-1]
		}

		newPosition++
	}

	if nodeCount == 0 {
		nodeValue = metadataSum
	}

	return
}

func Day08() (int, int) {
	fmt.Println("Day 08")

	_, metadata, nodeValue := extractMetadata(0)
	metadataSum := 0
	for _, m := range metadata {
		metadataSum += m
	}

	fmt.Println("Part 1:", metadataSum)
	fmt.Println("Part 2:", nodeValue)

	return metadataSum, nodeValue
}
