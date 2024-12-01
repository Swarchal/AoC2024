package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Input struct {
	a, b []int
}

func part1(input Input) int {
	var diff, total int
	for i := 0; i < len(input.a); i++ {
		diff = absInt(input.a[i], input.b[i])
		total += diff
	}
	return total
}

func part2(input Input) int {
	var similarity int
	for _, x := range input.a {
		var count int
		for _, y := range input.b {
			if x == y {
				count++
			}
		}
		similarity += count * x
	}
	return similarity
}

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func parseInput(path string, order bool) (Input, error) {
	var a, b int
	input := Input{}
	file, err := os.Open(path)
	if err != nil {
		return Input{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d   %d", &a, &b)
		input.a = append(input.a, a)
		input.b = append(input.b, b)
	}
	if order {
		sort.Ints(input.a)
		sort.Ints(input.b)
	}
	return input, nil
}

func main() {
	input, err := parseInput("./input.txt", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
