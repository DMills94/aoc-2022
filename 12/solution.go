package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	rows := strings.Split(data, "\n")
	rowC := len(rows)
	colC := len(rows[0])
	grid := make([][]int, rowC)
	for i := range grid {
		grid[i] = make([]int, colC)
	}
	start := [2]int{}
	end := [2]int{}

	chars := "abcdefghijklmnopqrstuvwxyz"
	for x, row := range rows {
		for y, val := range row {
			char := string(val)
			if char == "S" {
				start = [2]int{x, y}
				char = "a"
			} else if char == "E" {
				end = [2]int{x, y}
				char = "z"
			}
			elevation := strings.Index(chars, char)
			grid[x][y] = elevation
		}
	}

	// Part 1
	answerp1 := part1(grid, start, end)
	fmt.Println("Minimum steps", answerp1)

	// Part 2
	answerp2 := part2(grid, end)
	fmt.Println("Minimum steps", answerp2)
}

func part1(grid [][]int, start [2]int, end [2]int) int {
	prev := BFS(grid, start, end)
	path := RecontructPath(start, end, prev, 0)

	return len(path) - 1
}

func part2(grid [][]int, end [2]int) int {
	steps := []int{}
	starts := [][2]int{}
	for x, row := range grid {
		for y, val := range row {
			if val == 0 {
				starts = append(starts, [2]int{x, y})
			}
		}
	}

	for _, start := range starts {
		prev := BFS(grid, start, end)
		path := RecontructPath(start, end, prev, lo.Min(steps))

		if len(path) > 0 {
			steps = append(steps, len(path)-1)
		}
	}
	return lo.Min(steps)
}

func BFS(grid [][]int, start [2]int, end [2]int) map[[2]int][2]int {
	Q := [][2]int{}
	Q = append(Q, start)

	visited := map[[2]int]bool{}
	visited[start] = true

	prev := map[[2]int][2]int{}
	for len(Q) > 0 {
		node := [2]int{}
		node, Q = Q[0], Q[1:]
		neighbours := GetNeighbours(grid, node)

		for _, neighbour := range neighbours {
			if !visited[neighbour] {
				Q = append(Q, neighbour)
				visited[neighbour] = true
				prev[neighbour] = node
			}
		}
	}

	return prev
}

func RecontructPath(start [2]int, end [2]int, prev map[[2]int][2]int, shortestPath int) [][2]int {
	path := [][2]int{}
	at := end
	stillPath := true
	for stillPath {
		path = append(path, at)
		if shortestPath > 0 && len(path) > shortestPath {
			return [][2]int{}
		}
		if at == start {
			stillPath = false
		}
		at = prev[at]
	}

	revPath := lo.Reverse(path)
	if revPath[0] == start {
		return revPath
	}
	return [][2]int{}
}

func GetNeighbours(grid [][]int, node [2]int) [][2]int {
	dirs := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	neighbours := [][2]int{}

	for _, dir := range dirs {
		newCoord := [2]int{node[0] + dir[0], node[1] + dir[1]}
		if 0 <= newCoord[0] && newCoord[0] <= len(grid)-1 && 0 <= newCoord[1] && newCoord[1] <= len(grid[0])-1 {
			if grid[newCoord[0]][newCoord[1]]-grid[node[0]][node[1]] <= 1 {
				neighbours = append(neighbours, newCoord)
			}
		}
	}

	return neighbours
}
