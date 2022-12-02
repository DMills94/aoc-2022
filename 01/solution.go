package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func main() {
	// data, _ := os.ReadFile("./example.txt")
	data, _ := os.ReadFile("./input.txt")
	stringInput := string(data)

	elves := strings.Split(strings.ReplaceAll(stringInput, "\r\n", "\n"), "\n\n")
	
	elfCals := lo.Map(elves, func(carriedFood string, index int) int {
		calsArr := strings.Split(carriedFood, "\n")
		totalCals := lo.Reduce(calsArr, func(agg int, cal string, index int) int {
			intCal, _ := strconv.Atoi(cal)
			return agg + intCal
		}, 0)
		return totalCals
	})

	// Part 1
	fmt.Println("Max cals = ", lo.Max(elfCals))

	// Part 2
	sort.Ints(elfCals)
	fmt.Println("Max 3 values are: ", elfCals[len(elfCals) - 3], elfCals[len(elfCals) - 2], elfCals[len(elfCals) - 1])
	fmt.Println("The sum of these values are: ", elfCals[len(elfCals) - 3] + elfCals[len(elfCals) - 2] + elfCals[len(elfCals) - 1])

}
