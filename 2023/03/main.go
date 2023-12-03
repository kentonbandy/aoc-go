package main

import (
	"github.com/kentonbandy/aoc-go/helpers"
	"fmt"
)

func main() {
	lines := helpers.GetLines()
	// process the input into two maps that store data about the numbers
	// coordtonumid maps an individual coordinate to the beginning
	// coordinate of the number it's a part of (we'll call this the numid)
	// numidtonum maps a numid to the int it represents
	// this lets us easily look up a number given a coordinate of any part of it
	coordtonumid, numidtonum := buildCoordToNum(lines)

	partnumberids := make(map[string]struct{})
	gearsum := 0

	for i, line := range lines {
		previous, next := getPreviousAndNext(lines, i)

		for j, c := range line {
			if !isSymbol(byte(c)) {
				continue
			}
			// if it's a symbol, get part number ids
			getPartNumberIds(previous, line, next, i, j, coordtonumid, partnumberids)

			if !isGear(byte(c)) {
				continue
			}
			gearsum += getGearValue(previous, line, next, i, j, coordtonumid, numidtonum)
		}
	}

	fmt.Println(getPartSum(partnumberids, numidtonum))
	fmt.Println(gearsum)
}

func getPartSum(partnumberids map[string]struct{}, numidtonum map[string]int) (sum int) {
	for numid := range partnumberids {
		sum += numidtonum[numid]
	}
	return
}

func getPartNumberIds(previous string, line string, next string, i int, j int, coordtonumid map[string]string, partnumberids map[string]struct{}) {
	lookaround(previous, line, next, i, j, coordtonumid, partnumberids)
}

// given the coords of a *, looks around for numbers
// if it has exactly two number neighbors, it's a gear, returns the value
// otherwise returns 0
func getGearValue(previous string, line string, next string, i int, j int,
coordtonumid map[string]string, numidtonum map[string]int) (gearvalue int) {
	// use a map because go doesn't have sets lol
	numids := make(map[string]struct{})

	// populate numids with the ids of the numbers around the gear
	lookaround(previous, line, next, i, j, coordtonumid, numids)

	if len(numids) != 2 {
		return
	}
	// if there are exactly two numbers around the gear, it's a gear
	keys := []string{}
	for k := range numids {
		keys = append(keys, k)
	}
	return numidtonum[keys[0]] * numidtonum[keys[1]]
}

func lookaround(previous string, line string, next string, i int, j int,
coordtonumid map[string]string, numids map[string]struct{}) {
	// line above
	if previous != "" {
		lookaroundline(previous, i - 1, j, coordtonumid, numids)
	}
	// current line
	lookaroundline(line, i, j, coordtonumid, numids)
	// line below
	if next != "" {
		lookaroundline(next, i + 1, j, coordtonumid, numids)
	}
}

func lookaroundline(line string, i int, j int, coordtonumid map[string]string, numids map[string]struct{}) {
	// left
	if j > 0 && helpers.ByteIsInt(line[j - 1]) {
		coord := fmt.Sprintf("%d,%d", j - 1, i)
		numids[coordtonumid[coord]] = struct{}{}
	}
	// current
	if helpers.ByteIsInt(line[j]) {
		coord := fmt.Sprintf("%d,%d", j, i)
		numids[coordtonumid[coord]] = struct{}{}
	}
	// right
	if j < len(line) - 1 && helpers.ByteIsInt(line[j + 1]) {
		coord := fmt.Sprintf("%d,%d", j + 1, i)
		numids[coordtonumid[coord]] = struct{}{}
	}
}

func buildCoordToNum(lines []string) (coordtonumid map[string]string, numidtonum map[string]int) {
	coordtonumid = make(map[string]string)
	numidtonum = make(map[string]int)

	var coords []string
	thisnum := ""
	for i, line := range lines {
		for j, c := range line {
			char := string(c)
			if (helpers.IsInt(char)) {
				thisnum += char
				coord := fmt.Sprintf("%d,%d", j, i)
				coords = append(coords, coord)
				if j == len(line) - 1 {
					num := helpers.StringToInt(thisnum)
					for _, coord := range coords {
						coordtonumid[coord] = coords[0]
					}
					numidtonum[coords[0]] = num
					thisnum = ""
					coords = nil
				}
			} else {
				if (thisnum != "") {
					num := helpers.StringToInt(thisnum)
					for _, coord := range coords {
						coordtonumid[coord] = coords[0]
					}
					numidtonum[coords[0]] = num
					thisnum = ""
					coords = nil
				}
			}
		}
	}
	return
}

func getPreviousAndNext(lines []string, i int) (previous string, next string) {
	if (i > 0) {
		previous = lines[i - 1]
	}
	if (i < len(lines) - 1) {
		next = lines[i + 1]
	}
	return
}

func isSymbol(b byte) bool {
	notsymbols := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9','.'}
	
	return !helpers.ContainsByte(notsymbols, b)
}

func isGear(b byte) bool {
	return b == '*'
}