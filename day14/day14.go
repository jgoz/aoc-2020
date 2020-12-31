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

var memRe = regexp.MustCompile(`mem\[(\d+)]`)

func applyValueMask(in int, mask string) (out int) {
	numBits := len(mask)
	out = in
	for i, m := range mask {
		bit := 1 << (numBits - i - 1)
		switch m {
		case '0':
			out = out &^ bit
		case '1':
			out = out | bit
		}
	}
	return
}

func part1(program []string) (acc int, err error) {
	vars := make(map[int]int)
	var mask string
	for _, line := range program {
		tokens := strings.Split(line, " = ")
		cmd, value := tokens[0], tokens[1]

		switch cmd {
		case "mask":
			mask = value
		default:
			match := memRe.FindStringSubmatch(cmd)
			if len(match) != 2 {
				err = errors.New("command did not match regex")
				break
			}
			address, err := strconv.Atoi(match[1])
			if err != nil {
				break
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				break
			}
			vars[address] = applyValueMask(num, mask)
		}
	}
	for _, value := range vars {
		acc += value
	}
	return
}

func applyAddressMask(address int, mask string) (addresses []int, err error) {
	maskBits, err := strconv.ParseInt(strings.ReplaceAll(mask, "X", "0"), 2, strconv.IntSize)
	if err != nil {
		return
	}

	numBits := len(mask)
	addresses = []int{address | int(maskBits)}

	for i, m := range mask {
		if m == 'X' {
			bitmask := 1 << numBits-1-i
			var next []int
			for _, address := range addresses {
				next = append(next,
					address&^bitmask, // 0
					address|bitmask,  // 1
				)
			}
			addresses = next
		}
	}
	return
}

func part2(program []string) (acc int, err error) {
	vars := make(map[int]int)
	var mask string
	for _, line := range program {
		tokens := strings.Split(line, " = ")
		cmd, value := tokens[0], tokens[1]

		switch cmd {
		case "mask":
			mask = value
		default:
			match := memRe.FindStringSubmatch(cmd)
			if len(match) != 2 {
				err = errors.New("command did not match regex")
				break
			}
			address, err := strconv.Atoi(match[1])
			if err != nil {
				break
			}
			num, err := strconv.Atoi(value)
			if err != nil {
				break
			}
			addresses, err := applyAddressMask(address, mask)
			if err != nil {
				break
			}
			for _, address := range addresses {
				vars[address] = num
			}
		}
	}
	for _, value := range vars {
		acc += value
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
