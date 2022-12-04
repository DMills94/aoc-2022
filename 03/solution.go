package main

import (
	"fmt"
	"os"
	"strings"
)

const charValues = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	// data, _ := os.ReadFile("./example.txt")
	data, _ := os.ReadFile("./input.txt")
	rucksacks := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	// Part 1
	answerp1 := part1(rucksacks)
	fmt.Println("Sum of common priority items =", answerp1)

	// Part 2
	answerp2 := part2(rucksacks)
	fmt.Println("Sum of common priority items =", answerp2)
}

func part1(rucksacks []string) int {
	var count int
	for _, bag := range rucksacks {
		halfBag := len(bag) / 2
		firstComp := bag[:halfBag]
		secondComp := bag[halfBag:]

		var commonChar string
		for _, char := range firstComp {
			stringChar := string(char)
			matchedChar := strings.Contains(secondComp, stringChar)
			if matchedChar {
				commonChar = stringChar
				break
			}
		}
		commonCharVal := strings.Index(charValues, commonChar) + 1
		count += commonCharVal
	}
	return count
}

func part2(rucksacks []string) int {
	var groups [][3]string
	var currentGroup [3]string

	for index, bag := range rucksacks {
		currentGroup[index%3] = bag
		if index%3 == 2 {
			groups = append(groups, currentGroup)
		}
	}

	var count int
	for _, group := range groups {
		var commonChar string
		for _, char := range group[0] {
			stringChar := string(char)
			matchedCharElf2 := strings.Contains(group[1], stringChar) && strings.Contains(group[2], stringChar)
			if matchedCharElf2 {
				commonChar = stringChar
			}
		}
		commonCharVal := strings.Index(charValues, commonChar) + 1
		count += commonCharVal
	}
	return count
}
