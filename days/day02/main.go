package main

import (
	"aoc2025/internals/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// ParseLineIntoReport takes a line and converts it to the array of integer values
func ParseLineIntoReport(line string) ([]int, error) {
	parts := strings.Split(line, " ")
	values := make([]int, 0)
	for _, part := range parts {
		intValue, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("error converting %s into integer: %w", part, err)
		}

		values = append(values, intValue)
	}

	return values, nil

}

// IsReportSafe checks whether a report is safe (sorted in either order)
func IsReportSafe(report []int) bool {

	// Reports are safe if
	// 1. Numbers are all increasing or all decreasing
	// 2. Adjacent numbers differ by at least 1 and at most 3

	firstDifference := 0.0
	for i := 1; i < len(report); i++ {
		difference := report[i] - report[i-1]

		if absDiff := math.Abs(float64(difference)); absDiff < 1 || absDiff > 3 {
			return false
		}

		// check increasing/decreasing
		if i > 1 && firstDifference*float64(difference) <= 0 {
			return false
		}

		firstDifference = float64(difference)
	}

	return true
}

// remove will remove the i-th element in a slice
func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:s], newSlice[s+1:]...)
}

// IsReportSafeRemovingLevel checks whether a report is safe when removing a level
func IsReportSafeRemovingLevel(report []int) bool {

	if IsReportSafe(report) {
		return true
	}

	for i := range len(report) {
		changed := remove(report, i)
		if IsReportSafe(changed) {
			return true
		}
	}

	return false
}

// CountSafeReports prints the number of safe reports in a given file
func CountSafeReports(filename string) {
	lines, err := utils.ReadFileLines(filename)
	if err != nil {
		log.Fatal("Could not read file: ", err)
	}

	counts := 0
	countsRemovingLevel := 0

	for _, line := range lines {
		report, err := ParseLineIntoReport(line)
		if err != nil {
			log.Fatal("Error converting string to int", err)
		}

		if IsReportSafe(report) {
			counts++
		}

		if IsReportSafeRemovingLevel(report) {
			countsRemovingLevel++
		}
	}

	fmt.Printf("%s => Number of safe reports: %d\tRemoving level: %d\n", filename, counts, countsRemovingLevel)
}

func main() {
	CountSafeReports("days/day02/part1_example.txt")
	CountSafeReports("days/day02/part1.txt")
}
