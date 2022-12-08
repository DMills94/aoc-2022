package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	raw, _ := os.ReadFile("./example.txt")
	// raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	instructions := strings.Split(data, "\n")

	// Part 1
	answerp1 := part1(instructions)
	fmt.Println("Sums of files with size under 100k on disk", answerp1)

	// Part 2
	// answerp2 := part2(columns, instructions)
	// fmt.Println("Crates on the top of the stack Cratemover 9001 =", answerp2)
}

func part1(instructions []string) int {
	m := map[string]int{}
	directoryTree := []string{}
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "$ cd") {
			path := strings.Split(instruction, " ")[2]
			directoryTree = changeDirectory(directoryTree, path)
		} else if instruction == "$ ls" || strings.HasPrefix(instruction, "dir") {
			// do nothing
		} else {
			fileSize, _ := strconv.Atoi(strings.Split(instruction, " ")[0])

			for _, dir := range directoryTree {
				m[dir] += fileSize
			}
		}
	}
	fmt.Println(m)
	maxSize := 100000
	t := 0
	for _, v := range m {
		if v <= maxSize {
			t += v
		}
	}
	return t
}

func changeDirectory(directoryTree []string, path string) []string {
	if path == "/" {
		directoryTree = append(directoryTree, "/")
	} else if path == ".." {
		directoryTree = directoryTree[:len(directoryTree)-1]
	} else {
		directoryTree = append(directoryTree, path)
	}
	return directoryTree
}
