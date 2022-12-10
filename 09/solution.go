package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// R 4
// U 4
// L 3
// D 1
// R 4
// D 1
// L 5
// R 2

func main() {
	// raw, _ := os.ReadFile("./examplesmol.txt")
	// raw, _ := os.ReadFile("./examplebig.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	moves := strings.Split(data, "\n")

	// Part 1
	answerp1 := solve(moves, 2)
	fmt.Println("Unique tail visits", answerp1)

	// Part 2
	answerp2 := solve(moves, 10)
	fmt.Println("Unique tail visits", answerp2)
}

func solve(moves []string, ropeLength int) int {
	rope := [][2]int{}
	start := [2]int{0, 0}
	for i := 0; i < ropeLength; i++ {
		rope = append(rope, start)
	}

	tailVisits := map[[2]int]bool{}

	for _, move := range moves {
		dir := strings.Split(move, " ")[0]
		dis, _ := strconv.Atoi(strings.Split(move, " ")[1])
		switch dir {
		case "R":
			for i := 0; i < dis; i++ {
				rope = MoveRight(rope)
				tailVisits[rope[len(rope)-1]] = true
			}
		case "L":
			for i := 0; i < dis; i++ {
				rope = MoveLeft(rope)
				tailVisits[rope[len(rope)-1]] = true
			}
		case "U":
			for i := 0; i < dis; i++ {
				rope = MoveUp(rope)
				tailVisits[rope[len(rope)-1]] = true
			}
		case "D":
			for i := 0; i < dis; i++ {
				rope = MoveDown(rope)
				tailVisits[rope[len(rope)-1]] = true
			}
		}
	}
	return len(tailVisits)
}

func MoveRight(rope [][2]int) [][2]int {
	rope[0][0] = rope[0][0] + 1
	for i, knot := range rope[1:] {
		prevKnot := rope[i]
		rope = MoveTail(rope, prevKnot, knot, i)
	}
	return rope
}

func MoveUp(rope [][2]int) [][2]int {
	rope[0][1] = rope[0][1] + 1
	for i, knot := range rope[1:] {
		prevKnot := rope[i]
		rope = MoveTail(rope, prevKnot, knot, i)
	}
	return rope
}

func MoveLeft(rope [][2]int) [][2]int {
	rope[0][0] = rope[0][0] - 1
	for i, knot := range rope[1:] {
		prevKnot := rope[i]
		rope = MoveTail(rope, prevKnot, knot, i)
	}
	return rope
}

func MoveDown(rope [][2]int) [][2]int {
	rope[0][1] = rope[0][1] - 1
	for i, knot := range rope[1:] {
		prevKnot := rope[i]
		rope = MoveTail(rope, prevKnot, knot, i)
	}
	return rope
}

func MoveTail(rope [][2]int, prevKnot [2]int, currentKnot [2]int, ropeIndex int) [][2]int {
	touching, sameC, sameR := IsTouching(prevKnot, currentKnot)
	if !touching {
		isLeft := prevKnot[0]-currentKnot[0] > 0
		isDown := prevKnot[1]-currentKnot[1] > 0
		if !sameR && !sameC {
			if isLeft {
				rope[ropeIndex+1][0] = currentKnot[0] + 1
			} else {
				rope[ropeIndex+1][0] = currentKnot[0] - 1
			}
			if isDown {
				rope[ropeIndex+1][1] = currentKnot[1] + 1
			} else {
				rope[ropeIndex+1][1] = currentKnot[1] - 1
			}
		} else if !sameC {
			if isLeft {
				rope[ropeIndex+1] = [2]int{currentKnot[0] + 1, currentKnot[1]}
			} else {
				rope[ropeIndex+1] = [2]int{currentKnot[0] - 1, currentKnot[1]}
			}
		} else {
			if isDown {
				rope[ropeIndex+1] = [2]int{currentKnot[0], currentKnot[1] + 1}
			} else {
				rope[ropeIndex+1] = [2]int{currentKnot[0], currentKnot[1] - 1}
			}
		}
	}
	return rope
}

func IsTouching(head [2]int, tail [2]int) (bool, bool, bool) {
	touching := false
	if IntAbs(head[0]-tail[0]) <= 1 && IntAbs(head[1]-tail[1]) <= 1 {
		touching = true
	}
	sameCol := IntAbs(head[0]-tail[0]) == 0
	sameRow := IntAbs(head[1]-tail[1]) == 0
	return touching, sameCol, sameRow
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Can print the grid if you want
func MakeGrid(rope [][2]int) [][]string {
	sq := 20
	grid := [][]string{}
	for x := -sq; x <= sq; x++ {
		rowArr := []string{}
		for y := -sq; y <= sq; y++ {
			rowArr = append(rowArr, ".")
		}
		grid = append(grid, rowArr)
	}

	for i, v := range rope {
		if grid[-v[1]+sq][v[0]+sq] == "." {
			switch i {
			case 0:
				grid[-v[1]+sq][v[0]+sq] = "H"
			case len(rope) - 1:
				grid[-v[1]+sq][v[0]+sq] = "T"
			default:
				grid[-v[1]+sq][v[0]+sq] = fmt.Sprint(i)
			}
		}
	}
	if grid[0+sq][0+sq] == "." {
		grid[0+sq][0+sq] = "s"
	}

	return grid
}

func PrintGrid(rope [][2]int) {
	grid := MakeGrid(rope)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			fmt.Print(grid[row][col])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
