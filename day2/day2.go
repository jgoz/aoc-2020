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

func part1(lines []string) (valid int, err error) {
	re := regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+) (?P<char>\w): (?P<pass>.+)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil && len(matches) == 5 {
			min, err := strconv.Atoi(matches[1])
			if err != nil {
				break
			}

			max, err := strconv.Atoi(matches[2])
			if err != nil {
				break
			}

			char := matches[3]
			pass := matches[4]
			count := strings.Count(pass, char)

			if count >= min && count <= max {
				valid++
			}
		}
	}
	return
}

func part2(lines []string) (valid int, err error) {
	re := regexp.MustCompile(`^(?P<pos1>\d+)-(?P<pos2>\d+) (?P<char>\w): (?P<pass>.+)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil && len(matches) == 5 {
			pos1, err := strconv.Atoi(matches[1])
			if err != nil {
				break
			}

			pos2, err := strconv.Atoi(matches[2])
			if err != nil {
				break
			}

			char := matches[3]
			pass := matches[4]
			atPos1 := pass[pos1-1] == char[0]
			atPos2 := pass[pos2-1] == char[0]

			// XOR
			if !(atPos1 && atPos2) && (atPos1 || atPos2) {
				valid++
			}
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

	mult, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mult)
}
