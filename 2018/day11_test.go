package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerlevelCalcCases(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		x  int
		y  int
		s  int
		pl int
	}{
		{x: 3, y: 5, s: 8, pl: 4},
		{x: 122, y: 79, s: 57, pl: -5},
		{x: 217, y: 196, s: 39, pl: 0},
		{x: 101, y: 153, s: 71, pl: 4},
	}

	for _, tc := range testCases {
		assert.Equal(tc.pl, calculatePowerLevel(tc.x, tc.y, tc.s))
	}
}

func TestDay11(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day11()
	assert.Equal("20,50", part1)
	assert.Equal("238,278,9", part2)
}
