package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	// raw, _ := os.ReadFile("./example1.txt")
	// raw, _ := os.ReadFile("./example2.txt")
	// raw, _ := os.ReadFile("./example3.txt")
	// raw, _ := os.ReadFile("./example4.txt")
	// raw, _ := os.ReadFile("./example5.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := string(raw)

	// Part 1
	answerp1 := solve(data, 4)
	fmt.Println("Characters parsed before message", answerp1)

	// Part 1
	answerp2 := solve(data, 14)
	fmt.Println("Characters parsed before message", answerp2)
}

func solve(buffer string, length int) int {
	for i := 0; i < len(buffer)-length-1; i++ {
		chars := []rune(buffer[i : i+length])
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		prev := 1
		for curr := 1; curr < len(chars); curr++ {
			if chars[curr-1] != chars[curr] {
				chars[prev] = chars[curr]
				prev++
			}
		}
		uniqueVals := len(chars[:prev])
		if uniqueVals == length {
			return i + length
		}
	}
	return 0
}
