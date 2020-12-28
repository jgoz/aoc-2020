package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `
abc

a
b
c

ab
ac

a
a
a
a

b`

func TestPart1(t *testing.T) {
	valid := part1(strings.TrimSpace(input))

	assert.Equal(t, 11, valid)
}

func TestPart2(t *testing.T) {
	valid := part2(strings.TrimSpace(input))

	assert.Equal(t, 6, valid)
}
