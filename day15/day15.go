package main

import (
	"flag"
	"fmt"
)

func run(starting []int, turns int) (nth int) {
	tracker := make(map[int]int)
	var next int

	for turn := 1; turn < turns; turn++ {
		var num int
		if turn-1 < len(starting) {
			num = starting[turn-1]
		} else {
			num = next
		}

		prev, ok := tracker[num]
		if !ok {
			next = 0
		} else {
			next = turn - prev
		}
		tracker[num] = turn
	}
	nth = next
	return
}

func part1(starting []int) (nth int) {
	return run(starting, 2020)
}

func part2(starting []int) (nth int) {
	return run(starting, 30000000)
}

var part func([]int) (int)
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

	res := part([]int{0, 1, 4, 13, 15, 12, 16})
	fmt.Println(res)
}
