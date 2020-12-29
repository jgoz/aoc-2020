package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Layout constants
const (
	Empty = byte('L')
	Full  = byte('#')
	Floor = byte('.')
)

// A Layout is a 2-dimensional floorplan for an airplane
type Layout [][]byte

// Occupied describes an algorithm that counts the number of
// occpupied seats near a given seat.
type Occupied func(layout Layout, row, col int) (occupied int)

func (layout Layout) String() (str string) {
	var b strings.Builder
	for _, row := range layout {
		b.Write(row)
		b.WriteString("\n")
	}
	return b.String()
}

// IsValid indicates whether a seat exists within the layout.
func (layout Layout) IsValid(row, col int) (valid bool) {
	return row >= 0 && col >= 0 && row < len(layout) && col < len(layout[0])
}

// OccupiedAdjacent returns the number of occupied seats that are directly
// adjacent to the given seat.
func (layout Layout) OccupiedAdjacent(row, col int) (occupied int) {
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r == row && c == col {
				continue
			}
			if layout.IsValid(r, c) && layout[r][c] == Full {
				occupied++
			}
		}
	}
	return
}

var offsets = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

// OccupiedVisible returns the number of occupied seats in the direct
// line of sight in each direction of a given seat, ignoring floors.
func (layout Layout) OccupiedVisible(row, col int) (occupied int) {
	for _, offset := range offsets {
		ro, co := offset[0], offset[1]

	Look:
		for r, c := row+ro, col+co; layout.IsValid(r, c); r, c = r+ro, c+co {
			switch layout[r][c] {
			case Empty:
				break Look
			case Full:
				occupied++
				break Look
			}
		}
	}
	return
}

// NextValue returns the next occupied/empty state for a seat given its
// empty threshold and occupied algorithm.
func (layout Layout) NextValue(row, col int, emptyThreshold int, occupied Occupied) (value byte) {
	value = layout[row][col]
	switch value {
	case Empty:
		if occupied(layout, row, col) == 0 {
			value = Full
		}
	case Full:
		if occupied(layout, row, col) >= emptyThreshold {
			value = Empty
		}
	}
	return
}

// NextLayout returns the next seating layout given an empty threshold
// and occupied algorithm.
func (layout Layout) NextLayout(emptyThreshold int, occupied Occupied) (next Layout) {
	rowMax := len(layout)
	if rowMax == 0 {
		return
	}
	colMax := len(layout[0])

	next = make(Layout, rowMax)
	for row := 0; row < rowMax; row++ {
		next[row] = make([]byte, colMax)
		for col := 0; col < colMax; col++ {
			next[row][col] = layout.NextValue(row, col, emptyThreshold, occupied)
		}
	}
	return next
}

// Equals returns true if two layouts are identical.
func (layout Layout) Equals(other Layout) bool {
	for r := 0; r < len(layout); r++ {
		if !bytes.Equal(layout[r], other[r]) {
			return false
		}
	}
	return true
}

func makeLayout(lines []string) (layout Layout) {
	for _, line := range lines {
		layout = append(layout, []byte(line))
	}
	return
}

func part1(lines []string) (occupied int) {
	layout := make(Layout, len(lines))
	next := makeLayout(lines)

	for !layout.Equals(next) {
		layout = next
		next = layout.NextLayout(4, func(layout Layout, row, col int) (occupied int) {
			return layout.OccupiedAdjacent(row, col)
		})
	}

	rowMax := len(layout)
	colMax := len(layout[0])

	for row := 0; row < rowMax; row++ {
		for col := 0; col < colMax; col++ {
			if layout[row][col] == Full {
				occupied++
			}
		}
	}
	return
}

func part2(lines []string) (occupied int) {
	layout := make(Layout, len(lines))
	next := makeLayout(lines)

	for !layout.Equals(next) {
		layout = next
		next = layout.NextLayout(5, func(layout Layout, row, col int) (occupied int) {
			return layout.OccupiedVisible(row, col)
		})
	}

	rowMax := len(layout)
	colMax := len(layout[0])

	for row := 0; row < rowMax; row++ {
		for col := 0; col < colMax; col++ {
			if layout[row][col] == Full {
				occupied++
			}
		}
	}
	return
}

var part func([]string) int
var usePart2 bool

func init() {
	flag.BoolVar(&usePart2, "2", false, "Run part 2")
}

func main() {
	flag.Parse()
	if usePart2 {
		part = part2
	} else {
		part = part1
	}

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(part(lines))
}
