package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(findMaximum(aoc.FileBlocksToStringSliceSlice("input/01.txt")))
}

func findMaximum(foodGroups [][]string) int {
	var max int
	for _, foodGroup := range foodGroups {
		var total int
		for _, food := range foodGroup {
			total += aoc.Atoi(food)
		}
		if total > max {
			max = total
		}
	}
	return max
}
