package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay05(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day05()
	assert.Equal(11252, part1)
	assert.Equal(6118, part2)
}
