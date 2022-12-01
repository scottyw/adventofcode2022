package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(findMaximum(aoc.FileLinesToStringSlice("input/01.txt")))
}

func findMaximum(foods []string) int {
	var total, max int
	for _, food := range foods {
		if food == "" {
			if total > max {
				max = total
			}
			total = 0
		} else {
			total += aoc.Atoi(food)
		}
	}
	return max
}
