package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const STEP_SIZE_LIMIT = 3

type Report []int

func (r Report) allIncrease() bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] >= r[i+1] {
			return false
		}
	}
	return true
}

func (r Report) allDecrease() bool {
	for i := 0; i < len(r)-1; i++ {
		if r[i] <= r[i+1] {
			return false
		}
	}
	return true
}

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (r Report) smallSteps() bool {
	for i := 0; i < len(r)-1; i++ {
		if absInt(r[i], r[i+1]) > STEP_SIZE_LIMIT {
			return false
		}
	}
	return true
}

func (r Report) isSafe() bool {
	if (r.allDecrease() || r.allIncrease()) && r.smallSteps() {
		return true
	}
	return false
}

func (r Report) isSafeDampened() bool {
	if r.isSafe() {
		return true
	} else {
		// iterate through r, removing a single element at each index in turn
		for i := 0; i < len(r); i++ {
			tmp := make(Report, len(r))
			copy(tmp, r)
			tmp = slices.Delete(tmp, i, i+1)
			if tmp.isSafe() {
				return true
			}
		}
	}
	return false
}

func parseInput(path string) ([]Report, error) {
	reports := []Report{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		report := make([]int, len(line))
		for i, s := range line {
			report[i], _ = strconv.Atoi(s)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func part1(reports []Report) int {
	var total int
	for _, r := range reports {
		if r.isSafe() {
			total++
		}
	}
	return total
}

func part2(reports []Report) int {
	var total int
	for _, r := range reports {
		if r.isSafeDampened() {
			total++
		}
	}
	return total
}

func main() {
	input, err := parseInput(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))

}
