package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Part1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(510, Day01Part1())
}

func TestDay01Part2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(69074, Day01Part2())
}
