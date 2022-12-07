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
	for i := 0; i < len(lines); i += 3 {
		total += findCommon(lines[i], lines[i+1], lines[i+2])
	}
	return total
}

func findCommon(r1, r2, r3 string) int {
	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {
			if r1[i] == r2[j] {
				for k := 0; k < len(r3); k++ {
					if r1[i] == r3[k] {
						return findPriority(r1[i])
					}
				}
			}
		}
	}
	panic("no common")
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
