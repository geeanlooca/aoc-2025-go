package main

import (
	"aoc2025/internals/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Part1 solves the part 1 of the puzzle
func Part1(lines []string) {
	// fmt.Println(lines)

	totalResult := 0
	for _, line := range lines {
		totalResult += SolvePart1Line(line)
	}
	fmt.Println("Result: ", totalResult)
}

func SolvePart1Line(line string) int {

	done := false

	sumOfOps := 0

	for done != true {

		// first find the mul( occurrence
		match := "mul("
		idx := strings.Index(line, match)

		if idx < 0 {
			break
		}

		if len(line) < len(match) {
			break
		}

		// remove mul(
		line = line[idx+len(match):]

		// now find a comma
		commaIdx := strings.Index(line, ",")
		if commaIdx < 0 {
		}

		// from 0 to commaIdx we should get an integer to parse
		integerString := line[0:commaIdx]

		firstVal, err := strconv.Atoi(integerString)

		if err != nil {
			continue
		}

		// remove the integer from the string
		if len(line) < len(integerString) {
			break
		}

		line = line[len(integerString):]

		braceIdx := strings.Index(line, ")")
		if braceIdx < 0 {
			break
		}

		integerString = line[1:braceIdx]

		secondVal, err := strconv.Atoi(integerString)
		if err != nil {
			continue
		}

		result := firstVal * secondVal
		sumOfOps += result

		line = line[braceIdx+1:]
	}

	return sumOfOps
}

func Part1File(filename string) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		log.Fatal("Could not read file: ", err)
	}

	Part1(lines)
}

func main() {

	Part1File("days/day03/part1_example.txt")
	Part1File("days/day03/part1.txt")

}
