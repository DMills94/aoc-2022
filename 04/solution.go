package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const charValues = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	// data, _ := os.ReadFile("./example.txt")
	data, _ := os.ReadFile("./input.txt")
	pairs := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	// Part 1
	answerp1 := part1(pairs)
	fmt.Println("Completely overlapping elf work =", answerp1)

	// Part 2
	answerp2 := part2(pairs)
	fmt.Println("Partial overlapping elf work =", answerp2)
}

func part1(pairs []string) int {
	var overlappingPairs int

	for _, pair := range pairs {
		elfone := strings.Split(pair, ",")[0]
		elftwo := strings.Split(pair, ",")[1]

		elfoneStart, _ := strconv.Atoi(strings.Split(elfone, "-")[0])
		elfoneEnd, _ := strconv.Atoi(strings.Split(elfone, "-")[1])
		elftwoStart, _ := strconv.Atoi(strings.Split(elftwo, "-")[0])
		elftwoEnd, _ := strconv.Atoi(strings.Split(elftwo, "-")[1])

		if elfoneStart <= elftwoStart && elfoneEnd >= elftwoEnd || elfoneStart >= elftwoStart && elfoneEnd <= elftwoEnd {
			overlappingPairs++
		}
	}

	return overlappingPairs
}

func part2(pairs []string) int {
	var partOverlappingPairs int

	for _, pair := range pairs {
		elfone := strings.Split(pair, ",")[0]
		elftwo := strings.Split(pair, ",")[1]

		elfoneStart, _ := strconv.Atoi(strings.Split(elfone, "-")[0])
		elfoneEnd, _ := strconv.Atoi(strings.Split(elfone, "-")[1])
		elftwoStart, _ := strconv.Atoi(strings.Split(elftwo, "-")[0])
		elftwoEnd, _ := strconv.Atoi(strings.Split(elftwo, "-")[1])

		elfoneShifts := makeSlice(elfoneStart, elfoneEnd)
		elftwoShifts := makeSlice(elftwoStart, elftwoEnd)

		for _, num := range elfoneShifts {
			if slices.Contains(elftwoShifts, num) {
				partOverlappingPairs++
				break
			}
		}
	}

	return partOverlappingPairs
}

func makeSlice(min int, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}
