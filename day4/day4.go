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
	"pid"}

type passport map[string]string

func (p *passport) ValidateKeys() bool {
	for _, key := range requiredKeys {
		if _, ok := (*p)[key]; !ok {
			return false
		}
	}
	return true
}

var hgtRe = regexp.MustCompile(`(\d+)(cm|in)`)
var hclRe = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var eclValid = strings.Split("amb blu brn gry grn hzl oth", " ")
var pidRe = regexp.MustCompile(`^\d{9}$`)

func (p *passport) ValidateValues() (valid bool, err error) {
	if !p.ValidateKeys() {
		return
	}

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi((*p)["byr"])
	if err != nil || byr < 1920 || byr > 2002 {
		return
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi((*p)["iyr"])
	if err != nil || iyr < 2010 || iyr > 2020 {
		return
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi((*p)["eyr"])
	if err != nil || eyr < 2020 || eyr > 2030 {
		return
	}

	// hgt (Height) - a number followed by either cm or in:
	//  - If cm, the number must be at least 150 and at most 193.
	//  - If in, the number must be at least 59 and at most 76.
	hgtMatches := hgtRe.FindStringSubmatch((*p)["hgt"])
	if len(hgtMatches) != 3 {
		return
	}
	hgt, err := strconv.Atoi(hgtMatches[1])
	unit := hgtMatches[2]
	if err != nil || unit != "cm" && unit != "in" {
		return
	}
	if unit == "cm" && (hgt < 150 || hgt > 193) {
		return
	}
	if unit == "in" && (hgt < 59 || hgt > 76) {
		return
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !hclRe.MatchString((*p)["hcl"]) {
		return
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := (*p)["ecl"]
	if !contains(eclValid, ecl) {
		return
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if !pidRe.MatchString((*p)["pid"]) {
		return
	}

	valid = true
	return
}

func split(r rune) bool {
	return r == '\n' || r == ' '
}

func part1(batch string) (valid int, err error) {
	passports := strings.Split(batch, "\n\n")

	for _, passportStr := range passports {
		pass := make(passport)
		tokens := strings.FieldsFunc(passportStr, split)

		for _, token := range tokens {
			pair := strings.Split(token, ":")
			if len(pair) == 2 {
				pass[pair[0]] = pair[1]
			}
		}

		if pass.ValidateKeys() {
			valid++
		}
	}
	return
}

func part2(batch string) (valid int, err error) {
	passports := strings.Split(batch, "\n\n")

Root:
	for _, passportStr := range passports {
		pass := make(passport)
		tokens := strings.FieldsFunc(passportStr, split)

		for _, token := range tokens {
			pair := strings.Split(token, ":")
			if len(pair) == 2 {
				pass[pair[0]] = pair[1]
			}
		}

		isValid, err := pass.ValidateValues()
		if err != nil {
			break Root
		}
		if isValid {
			valid++
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

	valid, err := part(strings.Join(lines, "\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(valid)
}
