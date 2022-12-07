package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(calculateTotal(aoc.FileLinesToStringSlice("input/03.txt")))
}

func calculateTotal(lines []string) int {
	var total int
	for _, line := range lines {
		total += findDuplicate(line)
	}
	return total
}

func findDuplicate(rucksack string) int {
	l := len(rucksack)
	for i := 0; i < l/2; i++ {
		for j := l / 2; j < l; j++ {
			if rucksack[i] == rucksack[j] {
				return findPriority(rucksack[i])
			}
		}
	}
	panic("no duplicate")
}

func findPriority(duplicate byte) int {
	if duplicate >= 97 {
		// Lower case
		return int(duplicate) - 96 // 'a' maps to 1
	} else if duplicate >= 65 {
		// Upper case
		return int(duplicate) - 64 + 26 // 'A' maps to 27
	} else {
		panic("unexpected value")
	}
}
