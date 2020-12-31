package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 3, 6}, 436},
		{[]int{1, 3, 2}, 1},
		{[]int{2, 1, 3}, 10},
		{[]int{1, 2, 3}, 27},
		{[]int{2, 3, 1}, 78},
		{[]int{3, 2, 1}, 438},
		{[]int{3, 1, 2}, 1836},
	}
	for _, c := range cases {
		nth := part1(c.in)
		assert.Equal(t, c.want, nth)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0, 3, 6}, 175594},
		{[]int{1, 3, 2}, 2578},
		{[]int{2, 1, 3}, 3544142},
		{[]int{1, 2, 3}, 261214},
		{[]int{2, 3, 1}, 6895259},
		{[]int{3, 2, 1}, 18},
		{[]int{3, 1, 2}, 362},
	}
	for _, c := range cases {
		nth := part2(c.in)
		assert.Equal(t, c.want, nth)
	}
}
