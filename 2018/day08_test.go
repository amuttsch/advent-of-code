package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day08()
	assert.Equal(36307, part1)
	assert.Equal(25154, part2)
}
