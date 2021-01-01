package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var fieldRe = regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)

type Field struct {
	name   string
	lower1 int
	lower2 int
	upper1 int
	upper2 int
}

func (field Field) validate(value int) (valid bool) {
	return value >= field.lower1 && value <= field.lower2 ||
		value >= field.upper1 && value <= field.upper2
}

func makeField(line string) (field Field, err error) {
	matches := fieldRe.FindStringSubmatch(line)
	if len(matches) != 6 {
		err = errors.New("field line did not match regex: " + line)
		return
	}
	name := matches[1]
	lower1, err := strconv.Atoi(matches[2])
	if err != nil {
		return
	}
	lower2, err := strconv.Atoi(matches[3])
	if err != nil {
		return
	}
	upper1, err := strconv.Atoi(matches[4])
	if err != nil {
		return
	}
	upper2, err := strconv.Atoi(matches[5])
	if err != nil {
		return
	}
	field = Field{name, lower1, lower2, upper1, upper2}
	return
}

type Ticket []int

func makeTicket(line string) (ticket Ticket, err error) {
	values := strings.Split(line, ",")
	for _, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			break
		}
		ticket = append(ticket, num)
	}
	return
}

type Section int

const (
	Fields Section = iota
	YourTicket
	NearbyTickets
)

func parse(notes []string) (fields []Field, tickets []Ticket, err error) {
	section := Fields

Parse:
	for _, line := range notes {
		switch {
		case len(line) == 0:
			continue
		case strings.HasPrefix(line, "your ticket:"):
			section = YourTicket
			continue
		case strings.HasPrefix(line, "nearby tickets:"):
			section = NearbyTickets
			continue
		}

		switch section {
		case Fields:
			field, err := makeField(line)
			if err != nil {
				break Parse
			}
			fields = append(fields, field)
		case YourTicket:
			ticket, err := makeTicket(line)
			if err != nil {
				break Parse
			}
			tickets = append([]Ticket{ticket}, tickets...)
		case NearbyTickets:
			ticket, err := makeTicket(line)
			if err != nil {
				break Parse
			}
			tickets = append(tickets, ticket)
		}
	}
	return
}

func validate(fields []Field, tickets []Ticket) (valid []Ticket, rate int, err error) {
	for _, ticket := range tickets {
		validFields := make([][]Field, len(ticket))
		isValid := true
		for i, val := range ticket {
			for _, field := range fields {
				if field.validate(val) {
					validFields[i] = append(validFields[i], field)
				}
			}
			if len(validFields[i]) == 0 {
				isValid = false
				rate += val
			}
		}
		if isValid {
			valid = append(valid, ticket)
		}
	}
	return
}

func part1(notes []string) (rate int, err error) {
	fields, tickets, err := parse(notes)
	if err != nil {
		return
	}
	_, rate, err = validate(fields, tickets[1:]) // ignore "your ticket"
	return
}

func remove(slice []int, element int) []int {
	last := len(slice) - 1
	for i, v := range slice {
		if v == element {
			slice[i] = slice[last]
			slice[last] = 0
			slice = slice[:last]
			break
		}
	}
	return slice
}

func part2(notes []string) (mult int, err error) {
	fields, tickets, err := parse(notes)
	if err != nil {
		return
	}
	valid, _, err := validate(fields, tickets[1:]) // ignore "your ticket"

	candidates := make(map[string][]int)

	for _, field := range fields {
		for i := 0; i < len(fields); i++ {
			allValid := true
			for _, ticket := range valid {
				if !field.validate(ticket[i]) {
					allValid = false
					break
				}
			}
			if allValid {
				candidates[field.name] = append(candidates[field.name], i)
			}
		}
	}

	var found []int
	done := false

	for !done {
		done = true
		for _, indices := range candidates {
			if len(indices) == 1 {
				found = append(found, indices[0])
			}
		}
		for _, index := range found {
			for name, indices := range candidates {
				if len(indices) > 1 {
					candidates[name] = remove(indices, index)
					done = false
				}
			}
		}
	}

	mult = 1
	for name, indices := range candidates {
		if strings.HasPrefix(name, "departure ") {
			mult *= tickets[0][indices[0]]
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

	sum, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
