package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func countTrees(lines []string, slopeX int, slopeY int) (trees int) {
	width := len(lines[0])

	for x, y := 0, 0; y < len(lines); {
		if lines[y][x] == '#' {
			trees++
		}
		y += slopeY
		x = (x + slopeX) % width
	}
	return
}

func part1(lines []string) (trees int) {
	return countTrees(lines, 3, 1)
}

func part2(lines []string) (trees int) {
	return countTrees(lines, 1, 1) *
		countTrees(lines, 3, 1) *
		countTrees(lines, 5, 1) *
		countTrees(lines, 7, 1) *
		countTrees(lines, 1, 2)
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
