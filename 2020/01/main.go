package main

import (
	"errors"
	"fmt"
	"github.com/kentonbandy/aoc-go/helpers"
	"strconv"
)

func main() {
	lines := helpers.GetLines()

	for i, line := range lines {
		lineNum, _ := strconv.Atoi(line)
		toFind := 2020 - lineNum
		// part 1
		if helpers.ContainsString(lines[i+1:], fmt.Sprintf("%d", toFind)) {
			fmt.Println(lineNum * toFind)
		}
		// part 2
		num2, err := findNum2(lines[i+1:], toFind)
		if err != nil {
			continue
		}
		fmt.Println(lineNum * num2 * (2020 - lineNum - num2))
	}
}

func findNum2(lines []string, num int) (int, error) {
	for i, line := range lines {
		lineNum, _ := strconv.Atoi(line)
		if lineNum >= num {
			continue
		}
		toFind := num - lineNum
		if helpers.ContainsString(lines[i+1:], fmt.Sprintf("%d", toFind)) {
			return lineNum, nil
		}
	}
	return 0, errors.New("not found")
}
