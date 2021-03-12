package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
	"sync"
)

var polymer string
var wg sync.WaitGroup

func init() {
	dat, err := ioutil.ReadFile("input/day05.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	polymer = strings.Split(string(dat), "\n")[0]
}

func shrinkPolymer(p string) (string, int) {
	var shrinked strings.Builder
	destroyedUnits := 0

	// Iterate from 0 to the second last index as we will access p[i+1]
	for i := 0; i < len(p); i++ {
		if i == len(p)-1 {
			shrinked.WriteByte(p[i])
			break
		}

		diff := p[i] - p[i+1]
		if diff != 32 && diff != 224 {
			shrinked.WriteByte(p[i])
		} else {
			destroyedUnits++
			i++
		}
	}

	return shrinked.String(), destroyedUnits
}

func runReduction(polymer string, filterChar byte, outChan chan<- int) int {
	defer wg.Done()
	filteredPolymer := polymer
	if filterChar > 0 {
		filteredPolymer = strings.Replace(filteredPolymer, string(filterChar), "", -1)
		filteredPolymer = strings.Replace(filteredPolymer, string(filterChar-32), "", -1)
	}

	c := len(polymer)
	for c != 0 {
		filteredPolymer, c = shrinkPolymer(filteredPolymer)
	}

	outChan <- len(filteredPolymer)

	return len(filteredPolymer)
}

func Day05() (int, int) {
	fmt.Println("Day 05")
	nonFilteredChan := make(chan int, 100)
	filteredChan := make(chan int, 100)

	wg.Add(1)
	go runReduction(polymer, 0, nonFilteredChan)

	for i := 0; i < 26; i++ {
		wg.Add(1)
		go runReduction(polymer, byte('a'+i), filteredChan)
	}

	wg.Wait()
	close(filteredChan)

	smallestPolymer := math.MaxInt64
	for pc := range filteredChan {
		if pc < smallestPolymer {
			smallestPolymer = pc
		}
	}

	nonFilteredReductionLen := <-nonFilteredChan
	fmt.Println("Part 1:", nonFilteredReductionLen)
	fmt.Println("Part 2:", smallestPolymer)

	return nonFilteredReductionLen, smallestPolymer
}
