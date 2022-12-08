package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	instructions := strings.Split(data, "\n")

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

			for i := range directoryTree {
				dirpath := strings.Join(directoryTree[:i+1], "/")
				m[dirpath] += fileSize
			}
		}
	}

	// Part 1
	answerp1 := part1(m)
	fmt.Println("Sums of files with size under 100k on disk: ", answerp1)

	// Part 2
	answerp2 := part2(m)
	fmt.Println("Size of smallest single file removed: ", answerp2)
}

func part1(sizes map[string]int) int {
	maxSize := 100000
	t := 0
	for _, v := range sizes {
		if v <= maxSize {
			t += v
		}
	}
	return t
}

func part2(sizes map[string]int) int {
	diskSize := 70000000
	required := 30000000
	remainingspace := diskSize - sizes["root"]
	requiredspace := required - remainingspace
	biggestspace := diskSize
	var bestdirsize int
	for _, v := range sizes {
		spaceifdeleted := v - requiredspace
		if spaceifdeleted > 0 && spaceifdeleted < biggestspace {
			biggestspace = spaceifdeleted
			bestdirsize = v
		}
	}
	return bestdirsize
}

func changeDirectory(dirTree []string, path string) []string {
	if path == "/" {
		dirTree = append(dirTree, "root")
	} else if path == ".." {
		dirTree = dirTree[:len(dirTree)-1]
	} else {
		dirTree = append(dirTree, path)
	}
	return dirTree
}
