package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	pairs := strings.Split(data, "\n\n")

	// Part 1
	answerp1 := part1(pairs)
	fmt.Println("Sum of correct pair indices", answerp1)

	// Part 2
	// answerp2 := part2(treegrid)
	// fmt.Println("Optimal scenic score:", answerp2)
}

func part1(pairs []string) int {
	correctPairIndexes := []int{}
	for i := 1; i <= 10; i++ {
		pair := pairs[i-1]
		left, right := ProcessPairs(pair)
		fmt.Println("PAIR", i)
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println("---")

		index := 0
		run := true
		for run {
			if len(left) == 0 && len(right) != 0 {
				run = false
				correctPairIndexes = append(correctPairIndexes, i)
				fmt.Println("CORRECT > LEFT ARRAY EMPTY")
				break
			} else if len(left) != 0 && len(right) == 0 {
				run = false
				fmt.Println("WRONG > RIGHT ARRAY EMPTY")
				break
			}
			if index == len(left) {
				correctPairIndexes = append(correctPairIndexes, i)
				run = false
				fmt.Println("CORRECT > LEFT OUT OF ENTRIES")
				break
			}
			l := left[index]
			if index == len(right) {
				run = false
				fmt.Println("WRONG > RIGHT OUT OF ENTRIES")
				break
			}
			r := right[index]
			var res string
			res = Compare(l, r)

			if res == "LEFT" {
				correctPairIndexes = append(correctPairIndexes, i)
				run = false
				fmt.Println("CORRECT")
				break
			} else if res == "RIGHT" {
				run = false
				fmt.Println("WRONG")
				break
			}
			index++
		}
		fmt.Println()
	}
	fmt.Println(correctPairIndexes)
	return lo.Sum(correctPairIndexes)
}

// func part2(arr []string) {
// }

func ProcessPairs(pair string) ([]interface{}, []interface{}) {
	left := strings.Split(pair, "\n")[0]
	right := strings.Split(pair, "\n")[1]
	var leftJson []interface{}
	var rightJson []interface{}
	json.Unmarshal([]byte(left), &leftJson)
	json.Unmarshal([]byte(right), &rightJson)
	return leftJson, rightJson
}

func Compare(a interface{}, b interface{}) string {
	var res string
	fmt.Println(a, "VS", b)
	if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
		res = RightOrderInts(a.(float64), b.(float64))
	} else if reflect.TypeOf(b).Kind() == reflect.Float64 {
		newL := a.([]interface{})
		res = Compare(newL, []interface{}{b.(float64)})
	} else if reflect.TypeOf(a).Kind() == reflect.Float64 {
		newR := b.([]interface{})
		res = Compare([]interface{}{a.(float64)}, newR)
	} else {
		if len(a.([]interface{})) == 0 && len(b.([]interface{})) != 0 {
			return "LEFT"
		} else if len(a.([]interface{})) != 0 && len(b.([]interface{})) == 0 {
			return "RIGHT"
		} else {
			for i := 0; i < len(a.([]interface{})) && i < len(b.([]interface{})); i++ {
				newL := a.([]interface{})
				newR := b.([]interface{})
				res = Compare(newL[i], newR[i])

				if res == "LEFT" || res == "RIGHT" {
					break
				}
			}
		}
	}
	return res
}

func RightOrderInts(a float64, b float64) string {
	if a < b {
		return "LEFT"
	} else if a > b {
		return "RIGHT"
	} else {
		return "DRAW"
	}
}
