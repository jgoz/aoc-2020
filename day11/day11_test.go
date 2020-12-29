package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

func TestPart1(t *testing.T) {
	occupied := part1(strings.Split(strings.TrimSpace(input), "\n"))

	assert.Equal(t, 37, occupied)
}

func TestPart2(t *testing.T) {
	occupied := part2(strings.Split(strings.TrimSpace(input), "\n"))

	assert.Equal(t, 26, occupied)
}
