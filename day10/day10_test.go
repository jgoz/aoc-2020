package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{
	28,
	33,
	18,
	42,
	31,
	14,
	46,
	20,
	48,
	47,
	24,
	23,
	49,
	45,
	19,
	38,
	39,
	11,
	1,
	32,
	25,
	35,
	8,
	17,
	7,
	9,
	4,
	2,
	34,
	10,
	3,
}

func TestPart1(t *testing.T) {
	product := part1(input)

	assert.Equal(t, 220, product)
}

var short = []int{
	16,
	10,
	15,
	5,
	1,
	11,
	7,
	19,
	6,
	12,
	4,
}

func TestPart2(t *testing.T) {
	product := part2(input)

	assert.Equal(t, 19208, product)
}
