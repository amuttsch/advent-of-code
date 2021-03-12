package main

import (
	"flag"
	"fmt"
)

func main() {
	day := flag.Int("day", 0, "What day to execute")
	flag.Parse()

	switch *day {
	case 1:
		Day01Part1()
		Day01Part2()
	case 2:
		Day02Part1()
		Day02Part2()
	case 3:
		Day03()
	case 4:
		Day04()
	case 5:
		Day05()
	case 6:
		Day06()
	case 7:
		Day07()
	case 8:
		Day08()
	case 9:
		Day09()
	case 10:
		Day10()
	case 11:
		Day11()
	default:
		fmt.Println("Please use -day flag to execute puzzle")
	}
}
