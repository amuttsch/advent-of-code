package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02Part1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6723, Day02Part1())
}

func TestDay02Part2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("prtkqyluiusocwvaezjmhmfgx", Day02Part2())
}
