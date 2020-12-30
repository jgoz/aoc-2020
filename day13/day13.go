package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) (mult int, err error) {
	target, err := strconv.Atoi(lines[0])
	busStrings := strings.Split(lines[1], ",")
	times := make(map[int]int)

	for _, bus := range busStrings {
		if bus == "x" {
			continue
		}
		busTime, err := strconv.Atoi(bus)
		if err != nil {
			break
		}

		time := int(math.Ceil(float64(target)/float64(busTime))) * busTime
		times[busTime] = time
	}

	minTime, minBus := math.MaxInt32, math.MaxInt32
	for bus, time := range times {
		if time < minTime {
			minTime, minBus = time, bus
		}
	}

	mult = (minTime - target) * minBus
	return
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func part2(lines []string) (time int, err error) {
	busStrings := strings.Split(lines[1], ",")
	var timestamps []int

	for _, bus := range busStrings {
		if bus == "x" {
			timestamps = append(timestamps, -1)
			continue
		}
		timestamp, err := strconv.Atoi(bus)
		if err != nil {
			break
		}
		timestamps = append(timestamps, timestamp)
	}

	var found bool
	inc := timestamps[0]

	for time = 0; !found; {
		found = true // optimism

		for i, ts := range timestamps {
			if ts < 0 {
				continue // skip "x"
			}

			u := time + i
			correct := u%ts == 0

			if !correct {
				found = false // sigh
			} else if correct && i < len(timestamps)-1 {
				// This one is in the correct position, so adjust the
				// incrementor to the new lcm because they'll
				// always line up at this interval
				inc = lcm(inc, ts)
			}
		}

		if !found {
			time += inc
		}
	}
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

	res, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
