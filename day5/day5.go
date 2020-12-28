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

type seat struct {
	row    int
	column int
	seatID int
}

func toSeat(code string) (s *seat, err error) {
	rowPos, colPos := code[0:7], code[7:10]
	rowBin := strings.ReplaceAll(strings.ReplaceAll(rowPos, "F", "0"), "B", "1")
	colBin := strings.ReplaceAll(strings.ReplaceAll(colPos, "L", "0"), "R", "1")

	row, err := strconv.ParseInt(rowBin, 2, 0)
	if err != nil {
		return
	}

	col, err := strconv.ParseInt(colBin, 2, 0)
	if err != nil {
		return
	}

	s = &seat{row: int(row), column: int(col), seatID: int(row*8 + col)}
	return
}

func part1(codes []string) (highestID int, err error) {
	var seats = []seat{}
	for _, code := range codes {
		s, err := toSeat(code)
		if err != nil {
			break
		}
		seats = append(seats, *s)
	}
	for _, s := range seats {
		if s.seatID > highestID {
			highestID = s.seatID
		}
	}
	return
}

func part2(codes []string) (id int, err error) {
	var seats = []seat{}
	for _, code := range codes {
		s, err := toSeat(code)
		if err != nil {
			break
		}
		seats = append(seats, *s)
	}
	if err != nil {
		return
	}

	sort.SliceStable(seats, func(i, j int) bool {
		return seats[i].seatID < seats[j].seatID
	})

	for i := 1; i < len(seats)-1; i++ {
		prev, cur := seats[i-1], seats[i]
		if cur.seatID == prev.seatID+2 {
			id = cur.seatID - 1
			break
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

	id, err := part(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)
}
