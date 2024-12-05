/*
--- Day 4: Ceres Search ---

"Looks like the Chief's not here. Next!" One of The Historians pulls out a device and pushes the only button on it.
After a brief flash, you recognize the interior of the Ceres monitoring station!

As the search for the Chief continues, a small Elf who lives on the station tugs on your shirt; she'd like to know if
you could help her with her word search (your puzzle input). She only has to find one word: XMAS.

This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping other words.
It's a little unusual, though, as you don't merely need to find one instance of XMAS - you need to find all of them.
Here are a few ways XMAS might appear, where irrelevant characters have been replaced with .:

..X...
.SAMX.
.A..A.
XMAS.S
.X....

The actual word search will be full of letters instead. For example:

MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX

In this word search, XMAS occurs a total of 18 times; here's the same word search again, but where letters not involved
in any XMAS have been replaced with .:

....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX

Take a look at the little Elf's word search. How many times does XMAS appear?
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	TERM = "XMAS"
)

type Grid [][]string

type Coord struct {
	x, y int
}

func parseInput(path string) (Grid, error) {
	var grid Grid
	fhandle, err := os.Open(path)
	if err != nil {
		return grid, err
	}
	defer fhandle.Close()
	scanner := bufio.NewScanner(fhandle)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		grid = append(grid, row)
	}
	return grid, nil
}

func part1(grid Grid) int {
	// TODO: make this
	// going to be a depth-first search
	//
	// start at every index
	// look for first elem of term
	// then search all adjacent tiles
	// if not next elem in term, prune
	// keep going until end of term
	return 0
}

func (c Coord) isWithinLimits(xlim, ylim int) bool {
	if c.x < 0 || c.x > xlim || c.y < 0 || c.y > ylim {
		return false
	}
	return true
}

func (g Grid) termStartsAtCoord(c Coord) bool {
	for _, char := range TERM {
		char := string(char)
		if g[c.x][c.y] == char {
			fmt.Println(char)
		}
	}
	return false
}

func (c Coord) getAdjacent(xlim, ylim int) []Coord {
	var coords []Coord
	deltas := []int{-1, 0, 1}
	for _, dx := range deltas {
		for _, dy := range deltas {
			if dx == 0 && dy == 0 {
				continue
			}
			newCoord := Coord{c.x + dx, c.y + dy}
			if newCoord.isWithinLimits(xlim, ylim) {
				coords = append(coords, newCoord)
			}
		}
	}
	return coords
}

func main() {
	grid, err := parseInput(os.Args[1])
	if err != nil {
		panic(err)
	}
	grid.termStartsAtCoord(Coord{1, 1})
}
