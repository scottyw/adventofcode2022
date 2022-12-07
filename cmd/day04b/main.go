package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

// 2-4,6-8
var r = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func main() {
	fmt.Println(countOverlaps(aoc.FileLinesToStringSlice("input/04.txt")))
}

func countOverlaps(lines []string) int {
	var total int
	for _, line := range lines {
		if overlaps(line) {
			total++
		}
	}
	return total
}

func overlaps(assignment string) bool {
	matches := r.FindStringSubmatch(assignment)
	a1 := aoc.Atoi(matches[1])
	a2 := aoc.Atoi(matches[2])
	b1 := aoc.Atoi(matches[3])
	b2 := aoc.Atoi(matches[4])

	if b2 >= a1 && b2 <= a2 {
		return true
	}

	if a2 >= b1 && a2 <= b2 {
		return true
	}

	return false

}
