package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAocCases(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		players      int
		marbels      int
		winningScore int
	}{
		{players: 9, marbels: 25, winningScore: 32},
		{players: 10, marbels: 1618, winningScore: 8317},
		{players: 13, marbels: 7999, winningScore: 146373},
		{players: 17, marbels: 1104, winningScore: 2764},
		{players: 21, marbels: 6111, winningScore: 54718},
		{players: 30, marbels: 5807, winningScore: 37305},
	}

	for _, tc := range testCases {
		assert.Equal(tc.winningScore, gameOfMarbelsRing(tc.players, tc.marbels))
	}
}

func BenchmarkGameOfMarbles(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gameOfMarbels(459, 72103)
	}
}

func BenchmarkGameOfMarblesRing(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gameOfMarbelsRing(459, 72103)
	}
}

func TestDay09(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day09()
	assert.Equal(388131, part1)
	assert.Equal(3239376988, part2)
}
