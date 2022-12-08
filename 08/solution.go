package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	trees := strings.Split(data, "\n")

	treegrid := [][]int{}
	for _, line := range trees {
		m := []int{}
		for _, height := range line {
			intheight, _ := strconv.Atoi(string(height))
			m = append(m, intheight)
		}
		treegrid = append(treegrid, m)
	}

	// Part 1
	answerp1 := part1(treegrid)
	fmt.Println("Visible trees:", answerp1)

	// Part 2
	answerp2 := part2(treegrid)
	fmt.Println("Optimal scenic score:", answerp2)
}

func part1(grid [][]int) int {
	rows := len(grid[0])
	edges := (rows - 1) * 4 // Rows/Cols the same

	visibleInner := 0
	for x, line := range grid {
		if x == 0 || x == len(grid)-1 {
			continue
		}
		for y, height := range line {
			if y == 0 || y == len(line)-1 {
				continue
			}

			// Top
			topNumbers := treesUp(grid, x, y)
			if lo.Max(topNumbers) < height {
				visibleInner += 1
				continue
			}

			// Bottom
			bottomNumbers := treesDown(grid, x, y)
			if lo.Max(bottomNumbers) < height {
				visibleInner += 1
				continue
			}

			// Left
			leftNumbers := treesLeft(grid, x, y)
			if lo.Max(leftNumbers) < height {
				visibleInner += 1
				continue
			}

			// Right
			rightNumbers := treesRight(grid, x, y)
			if lo.Max(rightNumbers) < height {
				visibleInner += 1
			}
		}
	}

	visibleTrees := edges + visibleInner
	return visibleTrees
}

func part2(grid [][]int) int {
	var bestSceneScore int
	for x, line := range grid {
		if x == 0 || x == len(grid)-1 {
			continue
		}
		for y, height := range line {
			if y == 0 || y == len(line)-1 {
				continue
			}
			topTrees := treesUp(grid, x, y)
			scoreTop := sceneScore(height, topTrees)
			bottomTrees := treesDown(grid, x, y)
			scoreBottom := sceneScore(height, bottomTrees)
			leftTrees := treesLeft(grid, x, y)
			scoreLeft := sceneScore(height, leftTrees)
			rightTrees := treesRight(grid, x, y)
			scoreRight := sceneScore(height, rightTrees)

			totalSceneScore := scoreTop * scoreBottom * scoreLeft * scoreRight
			if totalSceneScore > bestSceneScore {
				bestSceneScore = totalSceneScore
			}
		}
	}
	return bestSceneScore
}

func treesUp(grid [][]int, x int, y int) []int {
	res := []int{}
	for i := 1; i < x+1; i++ {
		res = append(res, grid[x-i][y])
	}
	return res
}

func treesDown(grid [][]int, x int, y int) []int {
	res := []int{}
	for i := 1; i < len(grid[x])-x; i++ {
		res = append(res, grid[x+i][y])
	}
	return res
}

func treesLeft(grid [][]int, x int, y int) []int {
	res := []int{}
	for i := 1; i < y+1; i++ {
		res = append(res, grid[x][y-i])
	}
	return res
}

func treesRight(grid [][]int, x int, y int) []int {
	res := []int{}
	for i := 1; i < len(grid)-y; i++ {
		res = append(res, grid[x][y+i])
	}
	return res
}

func sceneScore(treeHeight int, treesInLine []int) int {
	score := len(treesInLine)
	for i, h := range treesInLine {
		if h >= treeHeight {
			score = i + 1
			break
		}
	}
	return score
}
