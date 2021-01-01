package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input1 = `
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`

func TestPart1(t *testing.T) {
	rate, err := part1(strings.Split(strings.TrimSpace(input1), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 71, rate)
}

const input2 = `
departure class: 0-1 or 4-19
departure row: 0-5 or 8-19
departure seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`

func TestPart2(t *testing.T) {
	mult, err := part2(strings.Split(strings.TrimSpace(input2), "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 1716, mult)
}
