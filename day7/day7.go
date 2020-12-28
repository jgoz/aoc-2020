package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Bag data type specifying quanity and color.
type Bag struct {
	num   int
	color string
}

// Lookup of bag color to the kinds of bags it can contain.
type Lookup map[string][]Bag

var bagRe = regexp.MustCompile(`(\d+) ([\w\s]+) bags?`)

func canContain(lookup Lookup, color string, targetColor string) bool {
	if color == targetColor {
		return true
	}
	bags := lookup[color]
	for _, bag := range bags {
		if canContain(lookup, bag.color, targetColor) {
			return true
		}
	}
	return false
}

func makeLookup(bags []string) (lookup Lookup) {
	lookup = make(map[string][]Bag)

	for _, bag := range bags {
		parts := strings.Split(bag, " bags contain ")
		if len(parts) != 2 {
			continue
		}
		color, contained := parts[0], parts[1]

		lookup[color] = make([]Bag, 0)

		if contained == "no other bags" {
			continue
		}

		containedBags := strings.Split(contained, ", ")
		for _, containedBag := range containedBags {
			matches := bagRe.FindStringSubmatch(containedBag)
			if len(matches) == 3 {
				num, err := strconv.Atoi(matches[1])
				if err != nil {
					break
				}

				lookup[color] = append(lookup[color], Bag{num: num, color: matches[2]})
			}
		}
	}
	return
}

func part1(bags []string, targetColor string) (count int, err error) {
	var lookup = makeLookup(bags)

	for color := range lookup {
		if color != targetColor && canContain(lookup, color, targetColor) {
			count++
		}
	}

	return
}

func countBags(lookup Lookup, color string) (count int) {
	for _, bag := range lookup[color] {
		count += bag.num
		count += bag.num * countBags(lookup, bag.color)
	}
	return
}

func part2(bags []string, targetColor string) (count int, err error) {
	var lookup = makeLookup(bags)
	count = countBags(lookup, targetColor)
	return
}

var part func([]string, string) (int, error)
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

	id, err := part(lines, "shiny gold")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}
