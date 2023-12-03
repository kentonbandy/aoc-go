package main

import (
	"github.com/kentonbandy/aoc-go/helpers"
	"fmt"
	"strings"
)

func main() {
	lines := helpers.GetLines()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for _, line := range lines {
		var num1 string
		var num2 string
		first := true
		for _, char := range line {
			if helpers.IsInt(string(char)) {
				num := string(char)
				if first {
					num1 = num
					first = false
				}
				num2 = num
			}
		}
		total += (helpers.StringToInt(num1 + num2))
	}

	fmt.Println(total)
}

func part2(lines []string) {
	digitStrings := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitConversions := map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}

	total := 0

	// HasPrefix HasSuffix
	for _, line := range lines {
		var first string
		var second string
		for first == "" {
			if helpers.IsInt(string(line[0])) {
				first = string(line[0])
			} else {
				for _, digit := range digitStrings {
					if strings.HasPrefix(line, digit) {
						first = digitConversions[digit]
					}
				}
			}
			if (first == "") {
				line = line[1:]
			}
		}
		for second == "" {
			lastInd := len(line) - 1
			if helpers.IsInt(string(line[lastInd])) {
				second = string(line[lastInd])
			} else {
				for _, digit := range digitStrings {
					if strings.HasSuffix(line, digit) {
						second = digitConversions[digit]
					}
				}
			}
			if (second == "") {
				line = line[:lastInd]
			}
		}
		total += (helpers.StringToInt(first + second))
	}

	fmt.Println(total)
}