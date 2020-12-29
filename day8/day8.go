package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func run(program []string) (acc int, terminated bool, err error) {
	visited := make(map[int]bool)
	var line int

	for !visited[line] {
		if line >= len(program) {
			terminated = true
			break
		}
		visited[line] = true

		tokens := strings.Split(program[line], " ")
		inst, offsetStr := tokens[0], tokens[1]

		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			break
		}

		switch inst {
		case "nop":
			line++
		case "jmp":
			line += offset
		case "acc":
			acc += offset
			line++
		}
	}
	return
}

func part1(program string) (acc int, err error) {
	acc, _, err = run(strings.Split(program, "\n"))
	return
}

func part2(program string) (acc int, err error) {
	var terminated bool
	lines := strings.Split(program, "\n")

	// Check unmodified program, just for safety
	acc, terminated, err = run(lines)

	if terminated {
		return
	}

	modified := make([]string, len(lines))

	for i, line := range lines {
		copy(modified, lines)

		switch {
		case strings.HasPrefix(line, "nop"):
			modified[i] = strings.ReplaceAll(line, "nop", "jmp")
		case strings.HasPrefix(line, "jmp"):
			modified[i] = strings.ReplaceAll(line, "jmp", "nop")
		}

		acc, terminated, err = run(modified)
		if terminated {
			break
		}
	}
	return
}

var part func(string) (int, error)
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

	sum, err := part(strings.Join(lines, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
