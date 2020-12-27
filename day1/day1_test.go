package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []int{1721,
		979,
		366,
		299,
		675,
		1456}

	mult, err := part1(input)

	assert.NoError(t, err)
	assert.Equal(t, 514579, mult)
}

func TestPart2(t *testing.T) {
	input := []int{1721,
		979,
		366,
		299,
		675,
		1456}

	mult, err := part2(input)

	assert.NoError(t, err)
	assert.Equal(t, 241861950, mult)
}
