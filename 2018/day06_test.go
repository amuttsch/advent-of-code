package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay06(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day06()
	assert.Equal(4887, part1)
	assert.Equal(34096, part2)
}
