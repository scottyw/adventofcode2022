package main

import (
	"fmt"
	"sort"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(findMaximum(aoc.FileLinesToStringSlice("input/01.txt")))
}

func findMaximum(foods []string) int {
	var total int
	var max []int
	for _, food := range foods {
		if food == "" {
			max = append(max, total)
			total = 0
		} else {
			total += aoc.Atoi(food)
		}
	}
	if total > 0 {
		max = append(max, total)
	}
	sort.Ints(max)
	var top3 int
	for i := len(max) - 3; i < len(max); i++ {
		top3 += max[i]
	}
	return top3
}
