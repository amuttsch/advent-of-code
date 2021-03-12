package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
)

var players int
var marbles int

type marble struct {
	value    int
	next     *marble
	previous *marble
}

func init() {
	dat, err := ioutil.ReadFile("input/day09.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	fmt.Sscanf(string(dat), "%d players; last marble is worth %d points", &players, &marbles)
}

func gameOfMarbelsRing(players, marbles int) int {
	// Each player has a score in this slice
	scores := make([]int, players)

	currentMarble := &ring.Ring{Value: 0}

	for nextMarble := 1; nextMarble <= marbles; nextMarble++ {
		// If the marble is a multple of 23 a player gets some points
		if nextMarble%23 == 0 {
			player := nextMarble % players

			// Add the marble number to their points
			scores[player] += nextMarble

			// Move the current marble 8 times to the left
			currentMarble = currentMarble.Move(-8)

			// And remove the next element (so the 7th element from the previous current one)
			removed := currentMarble.Unlink(1)

			// Add value to score
			scores[player] += removed.Value.(int)

			currentMarble = currentMarble.Move(1)
		} else {
			currentMarble = currentMarble.Move(1)
			newMarbleElem := &ring.Ring{Value: nextMarble}
			currentMarble = currentMarble.Link(newMarbleElem).Prev()
		}
	}

	// calculate highscore
	highscore := 0
	for _, s := range scores {
		if s > highscore {
			highscore = s
		}
	}
	return highscore
}

func gameOfMarbels(players, marbles int) int {
	// Each player has a score in this slice
	scores := make([]int, players)

	currentMarble := &marble{
		value: 0,
	}
	currentMarble.next = currentMarble
	currentMarble.previous = currentMarble

	for nextMarble := 1; nextMarble <= marbles; nextMarble++ {
		// If the marble is a multple of 23 a player gets some points
		if nextMarble%23 == 0 {
			player := nextMarble % players

			// Add the marble number to their points
			scores[player] += nextMarble

			// Move the current marble 7 times to the left
			for i := 0; i < 7; i++ {
				currentMarble = currentMarble.previous
			}

			// Add value to score
			scores[player] += currentMarble.value

			// Remove marble from chain
			currentMarble.previous.next = currentMarble.next
			currentMarble.next.previous = currentMarble.previous

			currentMarble = currentMarble.next
		} else {
			marbleBeforeNew := currentMarble.next
			marbleAfterNew := currentMarble.next.next

			newMarble := &marble{
				value:    nextMarble,
				next:     marbleAfterNew,
				previous: marbleBeforeNew,
			}
			marbleBeforeNew.next = newMarble
			marbleAfterNew.previous = newMarble
			currentMarble = newMarble
		}
	}

	// calculate highscore
	highscore := 0
	for _, s := range scores {
		if s > highscore {
			highscore = s
		}
	}
	return highscore
}

func Day09() (int, int) {
	fmt.Println("Day 09")

	highscorePart1 := gameOfMarbels(players, marbles)
	highscorePart2 := gameOfMarbels(players, marbles*100)

	fmt.Println("Part 1:", highscorePart1)
	fmt.Println("Part 2:", highscorePart2)

	return highscorePart1, highscorePart2
}
