package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")

	// Part 1
	answerp1 := part1(data)
	fmt.Println("Sum of correct pair indices", answerp1)

	// Part 2
	answerp2 := part2(data)
	fmt.Println("Decoder:", answerp2)
}

func part1(data string) int {
	pairs := strings.Split(data, "\n\n")
	correctPairIndexes := []int{}
	for i := 1; i <= len(pairs); i++ {
		pair := pairs[i-1]
		left, right := ProcessPairs(pair)
		res := Compare(left, right)

		if res == "LEFT" {
			correctPairIndexes = append(correctPairIndexes, i)
		}
	}
	return lo.Sum(correctPairIndexes)
}

func part2(data string) int {
	packets := strings.Split(strings.ReplaceAll(data, "\n\n", "\n"), "\n")
	divider1 := "[[2]]"
	divider2 := "[[6]]"
	packets = append(packets, divider1, divider2)

	sort.Slice(packets, func(i, j int) bool {
		pairsStr := packets[i] + "\n" + packets[j]
		left, right := ProcessPairs(pairsStr)
		res := Compare(left, right)
		if res == "LEFT" {
			return true
		} else {
			return false
		}
	})

	indexDiv1 := lo.IndexOf(packets, divider1) + 1
	indexDiv2 := lo.IndexOf(packets, divider2) + 1
	return indexDiv1 * indexDiv2
}

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
	if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
		res = RightOrderInts(a.(float64), b.(float64))
	} else if reflect.TypeOf(b).Kind() == reflect.Float64 {
		newL := a.([]interface{})
		res = Compare(newL, []interface{}{b.(float64)})
	} else if reflect.TypeOf(a).Kind() == reflect.Float64 {
		newR := b.([]interface{})
		res = Compare([]interface{}{a.(float64)}, newR)
	} else {
		if len(a.([]interface{})) == 0 && len(b.([]interface{})) == 0 {
			res = "DRAW"
		} else if len(a.([]interface{})) == 0 && len(b.([]interface{})) != 0 {
			res = "LEFT"
		} else if len(a.([]interface{})) != 0 && len(b.([]interface{})) == 0 {
			res = "RIGHT"
		} else {
			for i := 0; i < lo.Max([]int{len(a.([]interface{})), len(b.([]interface{}))}); i++ {
				if i == len(a.([]interface{})) {
					res = "LEFT"
					break
				} else if i == len(b.([]interface{})) {
					res = "RIGHT"
					break
				} else {
					res = Compare(a.([]interface{})[i], b.([]interface{})[i])
					if res != "DRAW" {
						break
					}
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
