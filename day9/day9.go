package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(numbers []int, preamble int) (outlier int) {
	for i := preamble; i < len(numbers); i++ {
		valid := false
		target := numbers[i]

	Search:
		for j := i - preamble; j < i; j++ {
			for k := i - preamble; k < i; k++ {
				if numbers[j]+numbers[k] == target {
					valid = true
					break Search
				}
			}
		}

		if !valid {
			outlier = target
			break
		}
	}
	return
}

func part2(numbers []int, preamble int) (weakness int) {
	outlier := part1(numbers, preamble)
	var sum, start, end int

	for i := 0; i < len(numbers); i++ {
		sum = 0
		start = i
		for j := start; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > outlier {
				break
			}
			if sum == outlier {
				end = j + 1
				break
			}
		}

		if end > start {
			break
		}
	}

	chain := numbers[start:end]
	min, max := chain[0], chain[0]
	for _, num := range chain {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return max + min
}

var part func([]int, int) int
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

	var err error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	numbers := make([]int, len(lines))
	for i, str := range lines {
		numbers[i], err = strconv.Atoi(str)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part(numbers, 25))
}
