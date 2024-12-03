package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mul struct {
	x       int
	y       int
	enabled bool
}

func (m Mul) calc() int {
	if !m.enabled {
		return 0
	}
	return m.x * m.y
}

func parseInput(path string) ([]Mul, error) {
	var output []Mul
	bytes, err := os.ReadFile(path)
	if err != nil {
		return output, err
	}
	content := string(bytes)
	enabledArr := createEnableArr(content)
	re := regexp.MustCompile(`(?m)mul\([0-9]*,[0-9]*\)`)
	matches := re.FindAllStringIndex(content, -1)
	for _, matchIdx := range matches {
		match := content[matchIdx[0]:matchIdx[1]]
		enabled := enabledArr[matchIdx[0]]
		instr := parseInstruction(match, enabled)
		output = append(output, instr)
	}
	return output, nil
}

func createEnableArr(content string) []bool {
	enableMap := map[int]bool{0: true}
	reDo := regexp.MustCompile(`do\(\)`)
	matchIdx := reDo.FindAllStringIndex(content, -1)
	for _, m := range matchIdx {
		enableMap[m[0]] = true
	}
	reDont := regexp.MustCompile(`don't\(\)`)
	matchIdx = reDont.FindAllStringIndex(content, -1)
	for _, m := range matchIdx {
		enableMap[m[0]] = false
	}
	enabledArr := make([]bool, len(content))
	enabled := true
	for i := 0; i < len(content); i++ {
		_, present := enableMap[i]
		if present {
			enabled = enableMap[i]
		}
		enabledArr[i] = enabled
	}
	return enabledArr
}

func parseInstruction(inst string, enabled bool) Mul {
	inst = strings.Replace(inst, "mul(", "", 1)
	inst = strings.Replace(inst, ")", "", 1)
	digits := strings.Split(inst, ",")
	mul := Mul{enabled: enabled}
	mul.x, _ = strconv.Atoi(digits[0])
	mul.y, _ = strconv.Atoi(digits[1])
	return mul
}

func part1(input []Mul) int {
	var total int
	for _, m := range input {
		m.enabled = true
		total += m.calc()
	}
	return total
}

func part2(input []Mul) int {
	var total int
	for _, m := range input {
		total += m.calc()
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
