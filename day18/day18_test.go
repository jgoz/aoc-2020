package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input1 = `
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`

func TestPart1(t *testing.T) {
	sum, err := part1(strings.Split(strings.TrimSpace(input1), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 26+437+12240+13632, sum)
}

const input2 = `
1 + (2 * 3) + (4 * (5 + 6))
2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`

func TestPart2(t *testing.T) {
	acc, err := part2(strings.Split(strings.TrimSpace(input2), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 51+46+1445+669060+23340, acc)
}
