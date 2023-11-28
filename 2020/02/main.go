package main

import (
	"fmt"
	"github.com/kentonbandy/aoc-go/helpers"
	"strings"
)

func main() {
	lines := helpers.GetLines()
	
	validCount1 := 0
	validCount2 := 0
	var num1 int
	var num2 int

	for _, line := range lines {
		num1, num2 = isValidPassword(line)
		validCount1 += num1
		validCount2 += num2
	}

	fmt.Println(validCount1)
	fmt.Println(validCount2)
}

func isValidPassword(line string) (int, int) {
	part1 := 0
	// such parsing
	pair := strings.Split(line, ": ")
	policy := strings.Split(pair[0], " ")
	char := policy[1]
	minmax := strings.Split(policy[0], "-")
	min := helpers.StringToInt(minmax[0])
	max := helpers.StringToInt(minmax[1])

	// such counting
	count := strings.Count(pair[1], char)

	// such logic
	if count >= min && count <= max {
		part1 = 1
	}
	
	// such casting
	minstring := string(pair[1][min-1])
	maxstring := string(pair[1][max-1])
	// such xor logic
	if ((minstring == char || maxstring == char) && minstring != maxstring) {
		return part1, 1
	}
	return part1, 0
}