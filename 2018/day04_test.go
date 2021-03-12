package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day04()
	assert.Equal(19025, part1)
	assert.Equal(23776, part2)
}
