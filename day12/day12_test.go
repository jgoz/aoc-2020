package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
F10
N3
F7
R90
F11
`

func TestPart1(t *testing.T) {
	acc, err := part1(strings.Split(strings.TrimSpace(input), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 25, acc)
}

func TestPart2(t *testing.T) {
	acc, err := part2(strings.Split(strings.TrimSpace(input), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 286, acc)
}
