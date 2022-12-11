package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	arr := strings.Split(data, "\n")

	// Part 1
	answerp1 := part1(arr)
	fmt.Println("Visible trees:", answerp1)

	// Part 2
	// answerp2 := part2(treegrid)
	// fmt.Println("Optimal scenic score:", answerp2)
}

func part1(arr []string) {
}

// func part2(arr []string) {
// }
