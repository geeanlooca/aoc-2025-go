package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func part_one_find_sum_distance(x, y []int) int {

	sort.Ints(x)
	sort.Ints(y)

	sum_of_distances := 0

	for i := range x {
		sum_of_distances += AbsDiff(x[i], y[i])
	}

	return sum_of_distances
}

func part_one_test() {
	first_list := []int{3, 4, 2, 1, 3, 3}
	second_list := []int{4, 3, 5, 3, 9, 3}

	test_result := part_one_find_sum_distance(first_list, second_list)
	if test_result != 11 {
		panic("Part 1 test failed")
	}
}

func ParseLine(line string) (int, int) {
	elements := strings.Split(line, "   ")
	a, err := strconv.Atoi(elements[0])
	if err != nil {
		panic("Error converting string to integer")
	}
	b, err := strconv.Atoi(elements[1])
	if err != nil {
		panic("Error converting string to integer")
	}

	return a, b
}

func part_two_find_similarity(x, y []int) int {
	m := make(map[int]int)

	for _, num := range y {
		m[num]++
	}

	similarity := 0
	for _, num := range x {
		similarity += num * m[num]
	}

	return similarity
}

func part_two_test() {

	first_list := []int{3, 4, 2, 1, 3, 3}
	second_list := []int{4, 3, 5, 3, 9, 3}
	result := part_two_find_similarity(first_list, second_list)

	if result != 31 {
		panic("Failed test for part two")
	}
}

func main() {
	fmt.Println("hello")

	part_one_test()

	x := make([]int, 0)
	y := make([]int, 0)
	lines := ReadInput("day01/part1.txt")
	for _, line := range lines {
		a, b := ParseLine(line)
		x = append(x, a)
		y = append(y, b)
	}

	fmt.Println(part_one_find_sum_distance(x, y))

	part_two_test()

	result := part_two_find_similarity(x, y)
	fmt.Println(result)

}
