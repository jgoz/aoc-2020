package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input1 = `
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
`

func TestPart1(t *testing.T) {
	acc, err := part1(strings.Split(strings.TrimSpace(input1), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 165, acc)
}

const input2 = `
mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
`

func TestPart2(t *testing.T) {
	acc, err := part2(strings.Split(strings.TrimSpace(input2), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 208, acc)
}
