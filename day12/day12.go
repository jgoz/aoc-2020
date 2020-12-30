package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

// Pos contains an x, y position or offset.
type Pos struct {
	x int
	y int
}

var directions = map[rune]Pos{
	'E': {x: 1, y: 0},
	'S': {x: 0, y: -1},
	'W': {x: -1, y: 0},
	'N': {x: 0, y: 1},
}

var turns = map[rune]map[rune]rune{
	'E': {'L': 'N', 'R': 'S'},
	'S': {'L': 'E', 'R': 'W'},
	'W': {'L': 'S', 'R': 'N'},
	'N': {'L': 'W', 'R': 'E'},
}

func changeBearing(bearing rune, dir rune, deg int) (nextBearing rune) {
	numTurns := deg / 90
	nextBearing = bearing
	for t := 0; t < numTurns; t++ {
		nextBearing = turns[nextBearing][dir]
	}
	return
}

func part1(program []string) (manhattan int, err error) {
	bearing := 'E'
	offset := directions[bearing]
	var pos Pos

Nav:
	for _, line := range program {
		switch line[0] {
		case 'F':
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			pos.x += offset.x * steps
			pos.y += offset.y * steps
		case 'N', 'E', 'S', 'W':
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			dir := directions[rune(line[0])]
			pos.x += dir.x * steps
			pos.y += dir.y * steps
		case 'R', 'L':
			dir := line[0]
			deg, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			bearing = changeBearing(bearing, rune(dir), deg)
			offset = directions[bearing]
		}
	}

	manhattan = abs(pos.x) + abs(pos.y)
	return
}

func rotate(waypoint Pos, dir rune, deg int) (next Pos) {
	s := int(math.Round(math.Sin(float64(deg) * math.Pi / 180)))
	c := int(math.Round(math.Cos(float64(deg) * math.Pi / 180)))

	switch dir {
	case 'R':
		next = Pos{
			x: waypoint.x*c + waypoint.y*s,
			y: -waypoint.x*s + waypoint.y*c,
		}
	case 'L':
		next = Pos{
			x: waypoint.x*c - waypoint.y*s,
			y: waypoint.x*s + waypoint.y*c,
		}
	}
	return
}

func part2(program []string) (manhattan int, err error) {
	waypoint := Pos{x: 10, y: 1}
	var ship Pos

Nav:
	for _, line := range program {
		switch line[0] {
		case 'F':
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			ship.x += waypoint.x * steps
			ship.y += waypoint.y * steps
		case 'N', 'E', 'S', 'W':
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			dir := directions[rune(line[0])]
			waypoint.x += dir.x * steps
			waypoint.y += dir.y * steps
		case 'R', 'L':
			dir := line[0]
			deg, err := strconv.Atoi(line[1:])
			if err != nil {
				break Nav
			}
			waypoint = rotate(waypoint, rune(dir), deg)
		}
	}

	manhattan = abs(ship.x) + abs(ship.y)
	return

}

var part func([]string) (int, error)
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

	manhattan, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manhattan)
}
