package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	datasplit := strings.Split(data, "\n\n")
	drawing := strings.Split(datasplit[0], "\n")
	instructions := strings.Split(datasplit[1], "\n")

	colCount := (len(drawing[0]) + 1) / 4

	columns := make(map[int][]string)

	for index, boxes := range drawing {
		if index == len(drawing)-1 {
			break
		}
		col := 1
		charIndex := 1
		for col <= colCount {
			boxInCol := string(boxes[charIndex])
			if boxInCol != string(' ') {
				columns[col] = append(columns[col], boxInCol)
			}
			col++
			charIndex += 4
		}
	}

	// Part 1
	answerp1 := part1(columns, instructions)
	fmt.Println("Crates on the top of the stack Cratemover 9000 =", answerp1)

	// Part 2
	answerp2 := part2(columns, instructions)
	fmt.Println("Crates on the top of the stack Cratemover 9001 =", answerp2)
}

func part1(columns map[int][]string, instructions []string) string {
	p1cols := make(map[int][]string)
	for k, v := range columns {
		p1cols[k] = v
	}
	for _, instruction := range instructions {
		re := regexp.MustCompile(`\d+`)
		vals := re.FindAllString(instruction, -1)
		amount, _ := strconv.Atoi(vals[0])
		fromCol, _ := strconv.Atoi(vals[1])
		toCol, _ := strconv.Atoi(vals[2])

		for i := 1; i <= amount; i++ {
			movingBlock := make([]string, 1)
			copy(movingBlock, p1cols[fromCol][:1])
			p1cols[fromCol] = p1cols[fromCol][1:]
			p1cols[toCol] = append(movingBlock, p1cols[toCol]...)
		}
	}
	return getTopBoxes(p1cols)
}

func part2(columns map[int][]string, instructions []string) string {
	p2cols := make(map[int][]string)
	for k, v := range columns {
		p2cols[k] = v
	}
	for _, instruction := range instructions {
		re := regexp.MustCompile(`\d+`)
		vals := re.FindAllString(instruction, -1)
		amount, _ := strconv.Atoi(vals[0])
		fromCol, _ := strconv.Atoi(vals[1])
		toCol, _ := strconv.Atoi(vals[2])

		movingBlock := make([]string, amount)
		copy(movingBlock, p2cols[fromCol][:amount])
		p2cols[fromCol] = p2cols[fromCol][amount:]
		p2cols[toCol] = append(movingBlock, p2cols[toCol]...)
	}
	return getTopBoxes(p2cols)
}

func getTopBoxes(cols map[int][]string) string {
	var topBoxes string
	for i := 1; i <= len(cols); i++ {
		topBoxes += cols[i][0]
	}
	return topBoxes
}
