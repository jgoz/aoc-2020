package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `939
7,13,x,x,59,x,31,19
`

func TestPart1(t *testing.T) {
	mult, err := part1(strings.Split(input, "\n"))

	assert.NoError(t, err)
	assert.Equal(t, 295, mult)
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"7,13,x,x,59,x,31,19", 1068781},
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	for _, c := range cases {
		mult, err := part2([]string{"foo", c.in})

		assert.NoError(t, err)
		assert.Equal(t, c.want, mult)
	}
}
