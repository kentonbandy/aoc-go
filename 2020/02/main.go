package main

import (
	"fmt"
	"github.com/kentonbandy/aoc-go/helpers"
	"strings"
	"regexp"
)

func main() {
	lines := helpers.GetLines()

	validCount1, validCount2 := 0, 0
	var num1, num2 int

	for _, line := range lines {
		num1, num2 = isValidPassword(line)
		validCount1 += num1
		validCount2 += num2
	}

	fmt.Printf("%v\n%v", validCount1, validCount2)
}

func isValidPasswordReg(line string) (part1 int, part2 int) {
	regex := regexp.MustCompile(`(\d{1})-(\d{1}) (.{1}): (?:[^g]*g){4,8}[^g]*$`)
	matched := regex.MatchString(line)
	if matched {
		part1 = 1
	}
	return
}

func isValidPassword(line string) (part1 int, part2 int) {
	// such parsing
	pair := strings.Split(line, ": ")
	policy := strings.Split(pair[0], " ")
	char := policy[1]
	minmax := strings.Split(policy[0], "-")
	min, max := helpers.StringToInt(minmax[0]), helpers.StringToInt(minmax[1])

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
		part2 = 1
	}

	return
}