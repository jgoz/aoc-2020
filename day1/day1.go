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

func main() {
	var usePart2 bool
	flag.BoolVar(&usePart2, "2", false, "Run part 2")
	flag.Parse()

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
	} else {
		var mult int
		var err error
		if usePart2 {
			mult, err = part2(numbers)
		} else {
			mult, err = part1(numbers)
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mult)
	}
}
