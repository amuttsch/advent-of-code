package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day10()
	assert.Equal("XPFXXXKL", part1)
	assert.Equal(10521, part2)
}
