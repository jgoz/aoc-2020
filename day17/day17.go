package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Active   = '#'
	Inactive = '.'
)

type Line []rune
type Region []Line
type Grid []Region
type Hypergrid []Grid

func (g Grid) String() string {
	var b strings.Builder
	for z := range g {
		_, err := fmt.Fprintf(&b, "z=%d\n", z)
		if err != nil {
			return "ERROR"
		}
		for y := range g[z] {
			for x := range g[z][y] {
				_, err := fmt.Fprintf(&b, "%c", g.valueAt(z, y, x))
				if err != nil {
					return "ERROR"
				}
			}
			_, err := fmt.Fprintln(&b)
			if err != nil {
				return "ERROR"
			}
		}
	}
	return b.String()
}

func (g Grid) valueAt(z, y, x int) rune {
	if z < 0 || y < 0 || x < 0 || z >= len(g) || y >= len(g[z]) || x >= len(g[z][y]) {
		return Inactive
	}
	if g[z][y][x] == 0 {
		return Inactive
	}
	return g[z][y][x]
}

func (h Hypergrid) valueAt(w, z, y, x int) rune {
	if w < 0 || w >= len(h) {
		return Inactive
	}
	return h[w].valueAt(z, y, x)
}

func (g Grid) activeNeighbours(z, y, x int) (active int) {
	for zz := z - 1; zz <= z+1; zz++ {
		for yy := y - 1; yy <= y+1; yy++ {
			for xx := x - 1; xx <= x+1; xx++ {
				if zz == z && yy == y && xx == x {
					continue // the cube
				}
				if g.valueAt(zz, yy, xx) == Active {
					active++
				}
			}
		}
	}
	return
}

func (h Hypergrid) activeNeighbours(w, z, y, x int) (active int) {
	for ww := w - 1; ww <= w+1; ww++ {
		for zz := z - 1; zz <= z+1; zz++ {
			for yy := y - 1; yy <= y+1; yy++ {
				for xx := x - 1; xx <= x+1; xx++ {
					if ww == w && zz == z && yy == y && xx == x {
						continue // the hypercube
					}
					if h.valueAt(ww, zz, yy, xx) == Active {
						active++
					}
				}
			}
		}
	}
	return
}

func (g Grid) next() (next Grid) {
	zlen := len(g) + 2
	ylen := len(g[0]) + 2
	xlen := len(g[0][0]) + 2

	next = make(Grid, zlen)
	for z := 0; z < zlen; z++ {
		next[z] = make(Region, ylen)
		for y := 0; y < ylen; y++ {
			next[z][y] = make(Line, xlen)
			for x := 0; x < xlen; x++ {
				an := g.activeNeighbours(z-1, y-1, x-1)
				cube := g.valueAt(z-1, y-1, x-1)
				next[z][y][x] = cube

				switch cube {
				case Active:
					if an != 2 && an != 3 {
						next[z][y][x] = Inactive
					}
				default:
					if an == 3 {
						next[z][y][x] = Active
					}
				}
			}
		}
	}
	return
}

func (h Hypergrid) next() (next Hypergrid) {
	wlen := len(h) + 2
	zlen := len(h[0]) + 2
	ylen := len(h[0][0]) + 2
	xlen := len(h[0][0][0]) + 2

	next = make(Hypergrid, wlen)
	for w := 0; w < wlen; w++ {
		next[w] = make(Grid, zlen)
		for z := 0; z < zlen; z++ {
			next[w][z] = make(Region, ylen)
			for y := 0; y < ylen; y++ {
				next[w][z][y] = make(Line, xlen)
				for x := 0; x < xlen; x++ {
					an := h.activeNeighbours(w-1, z-1, y-1, x-1)
					hypercube := h.valueAt(w-1, z-1, y-1, x-1)
					next[w][z][y][x] = hypercube

					switch hypercube {
					case Active:
						if an != 2 && an != 3 {
							next[w][z][y][x] = Inactive
						}
					default:
						if an == 3 {
							next[w][z][y][x] = Active
						}
					}
				}
			}
		}
	}
	return
}

func part1(lines []string) (active int) {
	grid := make(Grid, 1)
	grid[0] = make(Region, len(lines))
	for y, line := range lines {
		grid[0][y] = make(Line, len(line))
		for x, cube := range line {
			grid[0][y][x] = cube
		}
	}
	for i := 0; i < 6; i++ {
		//fmt.Printf("Cycle %d\n", i)
		//fmt.Println(grid.String())
		grid = grid.next()
	}
	for _, region := range grid {
		for _, line := range region {
			for _, cube := range line {
				if cube == Active {
					active++
				}
			}
		}
	}
	return
}

func part2(lines []string) (active int) {
	hypergrid := make(Hypergrid, 1)
	hypergrid[0] = make(Grid, 1)
	hypergrid[0][0] = make(Region, len(lines))
	for y, line := range lines {
		hypergrid[0][0][y] = make(Line, len(line))
		for x, cube := range line {
			hypergrid[0][0][y][x] = cube
		}
	}
	for i := 0; i < 6; i++ {
		hypergrid = hypergrid.next()
	}
	for _, grid := range hypergrid {
		for _, region := range grid {
			for _, line := range region {
				for _, cube := range line {
					if cube == Active {
						active++
					}
				}
			}
		}
	}
	return
}

var part func([]string) int
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

	fmt.Println(part(lines))
}
