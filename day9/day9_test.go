package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{
	35,
	20,
	15,
	25,
	47,
	40,
	62,
	55,
	65,
	95,
	102,
	117,
	150,
	182,
	127,
	219,
	299,
	277,
	309,
	576,
}

func TestPart1(t *testing.T) {
	outlier := part1(input, 5)

	assert.Equal(t, 127, outlier)
}

func TestPart2(t *testing.T) {
	weakness := part2(input, 5)

	assert.Equal(t, 62, weakness)
}
