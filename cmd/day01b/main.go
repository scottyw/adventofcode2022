package main

import (
	"fmt"
	"sort"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(findMaximum(aoc.FileBlocksToStringSliceSlice("input/01.txt")))
}

func findMaximum(foodGroups [][]string) int {
	var totals []int
	for _, foodGroup := range foodGroups {
		var total int
		for _, food := range foodGroup {
			total += aoc.Atoi(food)
		}
		totals = append(totals, total)
	}
	sort.Ints(totals)
	return totals[len(totals)-3] + totals[len(totals)-2] + totals[len(totals)-1]
}
