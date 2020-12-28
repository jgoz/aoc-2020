package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func contains(slice []string, val string) bool {
	for _, a := range slice {
		if a == val {
			return true
		}
	}
	return false
}

var requiredKeys = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func part1(batch string) (sum int) {
	groups := strings.Split(batch, "\n\n")

	for _, group := range groups {
		var counts = make(map[rune]int)
		answers := strings.Split(group, "\n")
		for _, answer := range answers {
			for _, char := range answer {
				counts[char] = 1
			}
		}
		for _, tally := range counts {
			sum += tally
		}
	}
	return
}

func part2(batch string) (sum int) {
	groups := strings.Split(batch, "\n\n")

	for _, group := range groups {
		var counts = make(map[rune]int)
		answers := strings.Split(group, "\n")

		for _, answer := range answers {
			for _, char := range answer {
				counts[char]++
			}
		}
		for _, tally := range counts {
			if tally == len(answers) {
				sum++
			}
		}
	}
	return
}

var part func(string) int
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

	sum := part(strings.Join(lines, "\n"))
	fmt.Println(sum)
}
