package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
`

func TestPart1(t *testing.T) {
	acc, err := part1(strings.TrimSpace(input))

	assert.NoError(t, err)
	assert.Equal(t, 5, acc)
}

func TestPart2(t *testing.T) {
	acc, err := part2(strings.TrimSpace(input))

	assert.NoError(t, err)
	assert.Equal(t, 8, acc)
}
