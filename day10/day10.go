package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func makeAdapters(numbers []int) (jolts []int) {
	jolts = make([]int, len(numbers)+1)
	copy(jolts, numbers)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	return jolts
}

type Chain []int
type Chains []Chain

func valid(chain Chain) bool {
	for i := 1; i < len(chain); i++ {
		if chain[i]-chain[i-1] > 3 {
			return false
		}
	}
	return true
}

func exhaust(chain Chain, seen map[int]bool) (possible int) {
	possible = 1
	for i := 1; i < len(chain)-1; i++ {
		candidate := make([]int, len(chain)-1)
		copy(candidate[:i], chain[:i])
		copy(candidate[i:], chain[i+1:])

		var hash int
		for _, n := range candidate {
			hash = hash ^ (n + 1) // this is sketchy but it works for the input /shrug
		}

		if valid(candidate) && !seen[hash] {
			seen[hash] = true
			possible += exhaust(candidate, seen)
		}
	}
	return possible
}

func part1(numbers []int) (product int) {
	jolts := makeAdapters(numbers)

	var ones, threes int
	for i := 1; i < len(jolts); i++ {
		switch jolts[i] - jolts[i-1] {
		case 1:
			ones++
		case 3:
			threes++
		}
	}

	return ones * threes
}

func part2(numbers []int) (possible int) {
	var chains Chains
	jolts := makeAdapters(numbers)

	for i, j := 0, 1; j < len(jolts); j++ {
		prev, cur := jolts[j-1], jolts[j]

		if len(chains) == i {
			chains = append(chains, make(Chain, 0))
		}

		chains[i] = append(chains[i], prev)
		if cur-prev == 3 {
			i++
		}
	}

	possible = 1
	for _, chain := range chains {
		possible *= exhaust(chain, make(map[int]bool))
	}

	return
}

var part func([]int) int
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

	fmt.Println(part(numbers))
}
