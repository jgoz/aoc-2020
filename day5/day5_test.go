package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = map[string]seat{
	"BFFFBBFRRR": {row: 70, column: 7, seatID: 567},
	"FFFBBBFRRR": {row: 14, column: 7, seatID: 119},
	"BBFFBBFRLL": {row: 102, column: 4, seatID: 820},
}

func TestPart1(t *testing.T) {
	var lines = []string{}
	for seat := range input {
		lines = append(lines, seat)
	}

	highestID, err := part1(lines)

	assert.NoError(t, err)
	assert.Equal(t, 820, highestID)
}
