package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
.#.
..#
###
`

func TestPart1(t *testing.T) {
	active := part1(strings.Split(strings.TrimSpace(input), "\n"))

	assert.Equal(t, 112, active)
}

func TestPart2(t *testing.T) {
	active := part2(strings.Split(strings.TrimSpace(input), "\n"))

	assert.Equal(t, 848, active)
}
