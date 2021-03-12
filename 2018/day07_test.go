package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay07(t *testing.T) {
	assert := assert.New(t)

	part1, part2 := Day07()
	assert.Equal("GJKLDFNPTMQXIYHUVREOZSAWCB", part1)
	assert.Equal(967, part2)
}
