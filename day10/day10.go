package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func makeAdapters(numbers []int) (jolts []int) {
	jolts = make([]int, len(numbers)+1)
	copy(jolts, numbers)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	return jolts
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

// A Chain represents a string of adapters that are exactly one jolt apart.
type Chain struct {
	length int
	hash   string
	values []int
}

func (chain Chain) String() string {
	return chain.hash
}

// Without removes the value at the given index and returns a new Chain.
func (chain *Chain) Without(index int) Chain {
	values := make([]int, chain.length-1)
	copy(values[:index], chain.values[:index])
	copy(values[index:], chain.values[index+1:])
	return makeChain(values)
}

// IsValid returns true if the chain is a contiguous set of increasing values.
func (chain *Chain) IsValid() bool {
	for i := 1; i < chain.length; i++ {
		if chain.values[i]-chain.values[i-1] > 3 {
			return false
		}
	}
	return true
}

func makeChain(values []int) (chain Chain) {
	var strs []string
	for _, num := range values {
		strs = append(strs, strconv.Itoa(num))
	}
	hash := strings.Join(strs, ",")
	length := len(values)
	return Chain{length, hash, values}
}

// exhaust recursively removes inner values of a chain and returns
// the number of valid alternative chains.
func exhaust(chain *Chain, seen map[string]bool) (possible int) {
	possible = 1
	for i := 1; i < chain.length-1; i++ {
		candidate := chain.Without(i)

		if candidate.IsValid() && !seen[candidate.hash] {
			seen[candidate.hash] = true
			possible += exhaust(&candidate, seen)
		}
	}
	return possible
}

func part2(numbers []int) (possible int) {
	var chains []Chain
	var values []int
	jolts := makeAdapters(numbers)

	for i, j := 0, 1; j < len(jolts); j++ {
		prev, cur := jolts[j-1], jolts[j]

		values = append(values, prev)
		if cur-prev == 3 {
			i++
			chains = append(chains, makeChain(values))
			values = make([]int, 0)
		}
	}

	possible = 1
	cache := make(map[int]int, 0)
	for _, chain := range chains {
		p, ok := cache[chain.length]
		if !ok {
			p = exhaust(&chain, make(map[string]bool))
			cache[chain.length] = p
		}
		possible *= p
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
