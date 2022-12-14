package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// raw, _ := os.ReadFile("./example.txt")
	raw, _ := os.ReadFile("./input.txt")
	data := strings.ReplaceAll(string(raw), "\r\n", "\n")
	monkies := strings.Split(data, "\n\n")

	// Part 1
	answerp1 := solve(monkies, 20, true)
	fmt.Println("Level of monkey business:", answerp1)

	// Part 2
	answerp2 := solve(monkies, 10000, false)
	fmt.Println("Level of monkey business:", answerp2)
}

func solve(monkies []string, rounds int, worried bool) int {
	items, ops, tests := ProcessInput(monkies)

	monkeycount := len(monkies)
	divisor := 1
	for _, v := range tests {
		divisor *= v["val"]
	}
	inspected := [8]int{}

	for i := 0; i < rounds; i++ {
		for m := 0; m < monkeycount; m++ {
			monkeyItems := items[m]
			inspected[m] += len(monkeyItems)
			monkeyOp := ops[m]
			monkeyTests := tests[m]

			for _, v := range monkeyItems {
				newV := HandleOp(v, monkeyOp["operator"], monkeyOp["amount"], worried) % divisor
				var throwToMonkey int
				if newV%monkeyTests["val"] == 0 {
					throwToMonkey = monkeyTests["y"]
				} else {
					throwToMonkey = monkeyTests["n"]
				}
				items[throwToMonkey] = append(items[throwToMonkey], newV)
			}
			items[m] = []int{}
		}
	}

	sort.Ints(inspected[:])
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}

// func part2(arr []string) {
// }

func ProcessInput(monkies []string) ([][]int, []map[string]string, []map[string]int) {
	items := [][]int{}
	ops := []map[string]string{}
	tests := []map[string]int{}
	for _, monke := range monkies {
		lines := strings.Split(monke, "\n")
		t := map[string]int{}
		for _, line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "Starting") {
				re := regexp.MustCompile(`\d+`)
				vals := re.FindAllString(line, -1)
				i := lo.Map(vals, func(str string, i int) int {
					intVal, _ := strconv.Atoi(str)
					return intVal
				})
				items = append(items, i)
			}
			if strings.HasPrefix(strings.TrimSpace(line), "Operation") {
				o := map[string]string{}
				o["operator"] = strings.Split(strings.TrimSpace(line), " ")[4]
				o["amount"] = strings.Split(strings.TrimSpace(line), " ")[5]
				ops = append(ops, o)
			}
			if strings.HasPrefix(strings.TrimSpace(line), "Test") {
				val, _ := strconv.Atoi(strings.Split(strings.TrimSpace(line), " ")[3])
				t["val"] = val
			}
			if strings.HasPrefix(strings.TrimSpace(line), "If true") {
				monkeyInt, _ := strconv.Atoi(strings.Split(strings.TrimSpace(line), " ")[5])
				t["y"] = monkeyInt
			}
			if strings.HasPrefix(strings.TrimSpace(line), "If false") {
				monkeyInt, _ := strconv.Atoi(strings.Split(strings.TrimSpace(line), " ")[5])
				t["n"] = monkeyInt
			}
		}
		tests = append(tests, t)
	}
	return items, ops, tests
}

func HandleOp(old int, operator string, amount string, worried bool) int {
	var num int
	if amount == "old" {
		num = old
	} else {
		num, _ = strconv.Atoi(amount)
	}
	var new int
	if operator == "*" {
		new = old * num
	} else if operator == "+" {
		new = old + num
	}
	if worried {
		return new / 3
	} else {
		return new
	}
}
