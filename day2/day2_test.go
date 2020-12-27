package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

func TestPart1(t *testing.T) {
	valid, err := part1(strings.Split(input, "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 2, valid)
}

func TestPart2(t *testing.T) {
	valid, err := part2(strings.Split(input, "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 1, valid)
}
