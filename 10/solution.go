package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// raw, _ := os.ReadFile("./examplebig.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	instr := strings.Split(data, "\n")

	// Part 1
	answerp1 := part1(instr)
	fmt.Println("X:", answerp1)

	// Part 2
	part2(instr)
}

func part1(instr []string) int {
	reg := 1
	cycle := 1
	exe := 0
	repeatline := false
	sumSignals := 0
	signals := []int{}

	for cycle < 221 {
		if (cycle-20)%40 == 0 {
			sumSignals += cycle * reg
			signals = append(signals, cycle*reg)
		}
		line := instr[exe]
		if strings.Split(line, " ")[0] == "noop" {
			exe++
		} else if strings.Split(line, " ")[0] == "addx" {
			if repeatline {
				amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
				exe++
				reg += amount
				repeatline = false
			} else {
				repeatline = true
			}
		}
		cycle++
	}
	return sumSignals
}

func part2(instr []string) {
	reg := 1
	cycle := 1
	sprite := 0
	exe := 0
	repeatline := false

	for cycle < 241 {
		if IntAbs(reg-sprite) <= 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if cycle%40 == 0 {
			sprite = -1
			fmt.Print("\n")
		}
		line := instr[exe]
		if strings.Split(line, " ")[0] == "noop" {
			exe++
		} else if strings.Split(line, " ")[0] == "addx" {
			if repeatline {
				amount, _ := strconv.Atoi(strings.Split(line, " ")[1])
				exe++
				reg += amount
				repeatline = false
			} else {
				repeatline = true
			}
		}
		cycle++
		sprite++
	}
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
