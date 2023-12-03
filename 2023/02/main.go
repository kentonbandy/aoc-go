package main

import (
	"github.com/kentonbandy/aoc-go/helpers"
	"fmt"
	"regexp"
)

func main() {
	lines := helpers.GetLines()

	possible := 0
	sumpower := 0
	for i, line := range lines {
		// both parts of the puzzle use the largest values for each color
		largest := getLargest(line)
		if isPossible(largest) {
			possible += (i + 1)
		}
		sumpower += largest["red"] * largest["green"] * largest["blue"]
	}

	fmt.Println(possible)
	fmt.Println(sumpower)
}

func getLargest(line string) map[string]int {
	largest := map[string]int { "red": 0, "green": 0, "blue": 0, }

	// grab the numbers and corresponding colors
	regex := regexp.MustCompile(`(\d+) ([a-z]+)`)
	subgroups := regex.FindAllStringSubmatch(line, -1)
	for _, group := range subgroups {
		num := helpers.StringToInt(group[1])
		color := group[2]
		if num > largest[color] {
			largest[color] = num
		}
	}

	return largest
}

func isPossible(largest map[string]int) bool {
	// canned values from the puzzle
	max := map[string]int { "red": 12, "green": 13, "blue": 14, }

	for color, value := range max {
		if largest[color] > value {
			return false
		}
	}
	return true
}