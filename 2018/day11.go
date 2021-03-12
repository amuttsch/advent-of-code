package main

import (
	"fmt"
	"sync"
)

func calculatePowerLevel(x, y, s int) int {
	//Find the fuel cell's rack ID, which is its X coordinate plus 10.
	rackID := x + 10

	//Begin with a power level of the rack ID times the Y coordinate.
	powerLevel := rackID * y

	//Increase the power level by the value of the grid serial number (your puzzle input).
	powerLevel += s

	//Set the power level to itself multiplied by the rack ID.
	powerLevel *= rackID

	//Keep only the hundreds digit of the power level (so 12345 becomes 3; numbers with no hundreds digit become 0).
	powerLevel %= 1000
	powerLevel /= 100

	//Subtract 5 from the power level.
	return powerLevel - 5
}

func calculatePowerlevelCell(xs, ys, gridSize, s int) (powerLevel int) {
	for x := xs; x < xs+gridSize; x++ {
		for y := ys; y < ys+gridSize; y++ {
			powerLevel += calculatePowerLevel(x, y, s)
		}
	}
	return
}

func Day11() (string, string) {
	fmt.Println("Day 11")

	gridSerialNumber := 2866

	type gridResult struct {
		x  int
		y  int
		s  int
		pl int
	}
	var part1GridResult gridResult
	var part2GridResult gridResult

	for x := 0; x < 300-3; x++ {
		for y := 0; y < 300-3; y++ {
			pl := calculatePowerlevelCell(x, y, 3, gridSerialNumber)
			if pl > part1GridResult.pl {
				part1GridResult.pl = pl
				part1GridResult.x = x
				part1GridResult.y = y
			}
		}
	}

	var wg sync.WaitGroup
	part2ResultChan := make(chan gridResult, 300)

	for s := 1; s <= 300; s++ {
		wg.Add(1)
		s := s
		go func() {
			for x := 0; x <= 300-s; x++ {
				for y := 0; y <= 300-s; y++ {
					pl := calculatePowerlevelCell(x, y, s, gridSerialNumber)
					r := gridResult{
						x: x, y: y, s: s, pl: pl,
					}

					part2ResultChan <- r
				}
			}
			wg.Done()
		}()
	}

	go func() {
		for r := range part2ResultChan {
			if r.pl > part2GridResult.pl {
				part2GridResult = r
			}
		}
	}()
	wg.Wait()
	close(part2ResultChan)

	part1 := fmt.Sprintf("%d,%d", part1GridResult.x, part1GridResult.y)
	part2 := fmt.Sprintf("%d,%d,%d", part2GridResult.x, part2GridResult.y, part2GridResult.s)

	fmt.Printf("Part 1: %s - Level %d\n", part1, part1GridResult.pl)
	fmt.Printf("Part 2: %s - Level %d\n", part2, part2GridResult.pl)

	return part1, part2
}
