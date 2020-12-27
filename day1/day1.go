package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(numbers []int) (mult int, err error) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == 2020 {
				return (numbers[i] * numbers[j]), nil
			}
		}
	}
	return -1, errors.New("expenseReport: unable to find two numbers that add to 2020")
}

func part2(numbers []int) (mult int, err error) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					return (numbers[i] * numbers[j] * numbers[k]), nil
				}
			}
		}
	}
	return -1, errors.New("expenseReport: unable to find three numbers that add to 2020")
}

var part func([]int) (int, error)
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

	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mult, err := part(numbers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mult)
}
