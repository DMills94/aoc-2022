package main

import (
	"fmt"
	"os"
	"strings"
)

// A / X = rock
// B / Y = paper
// C / Z = scissors

// Rock = 1
// Paper = 2
// Scissors = 3

// Loss = 0
// Draw = 3
// Win = 6

var moveType = map[string]string{"A":"rock", "X":"rock", "B":"paper", "Y":"paper", "C":"scissors", "Z":"scissors"}
var outcomePoints = map[string]int{"w":6, "d":3, "l":0}

func main() {
	// data, _ := os.ReadFile("./example.txt")
	data, _ := os.ReadFile("./input.txt")
	moves := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")
	
	// Part 1
	var totalPointsP1 int
	for _, turn := range moves {
		opponentAction := strings.Split(turn, " ")[0]
		myAction := strings.Split(turn, " ")[1]
		actionPoints := PointsForChoice(moveType[myAction])
		outcomePoints := PointsForOutcomeP1(myAction, opponentAction)

		totalPointsP1 += actionPoints + outcomePoints
	}

	fmt.Println("My total score from P1 strategy is:", totalPointsP1)

	// Part 2
	var totalPointsP2 int
	outcomeMap := map[string]string{"X":"l", "Y":"d", "Z":"w"}
	for _, turn := range moves {
		opponentAction := strings.Split(turn, " ")[0]
		outcome := outcomeMap[strings.Split(turn, " ")[1]]
		myAction := determineMove(opponentAction, outcome)
		actionPoints := PointsForChoice(myAction)

		totalPointsP2 += actionPoints + outcomePoints[outcome]
	}
	fmt.Println("My total score from P2 strategy is:", totalPointsP2)
}

func PointsForChoice(move string) int {
	movePoints := map[string]int{"rock":1, "paper":2, "scissors":3}
	return movePoints[move]
}

func PointsForOutcomeP1(myAction string, opponentAction string) int {
	if moveType[myAction] == moveType[opponentAction] {
		return outcomePoints["d"]
	}

	if myAction == "X" { // Rock
		if opponentAction == "B" { // Paper
			return outcomePoints["l"]
		} else { // Scissors
			return outcomePoints["w"]
		}
	} else if myAction == "Y" { // Paper
		if opponentAction == "A" { // Rock
			return outcomePoints["w"]
		} else { // Scissors
			return outcomePoints["l"]
		}
	} else { // myAction == "Z" - Scissors
		if opponentAction == "A" { // Rock
			return outcomePoints["l"]
		} else { // Paper
			return outcomePoints["w"]
		}
	}
}

func determineMove(opponentAction string, outcome string) string {
	if outcome == "d" {
		return moveType[opponentAction]
	}
	if opponentAction == "A" { // Rock
		if outcome == "w" {
			return "paper"
		} else {
			return "scissors"
		}
	} else if opponentAction == "B" { // Paper
		if outcome == "w" {
			return "scissors"
		} else {
			return "rock"
		}
	} else { // Scissors
		if outcome == "w" {
			return "rock"
		} else {
			return "paper"
		}
	}
}