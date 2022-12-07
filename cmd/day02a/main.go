package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

func main() {
	fmt.Println(findScore(aoc.FileLinesToStringSlice("input/02.txt")))
}

func findScore(rounds []string) int {
	var total int
	for _, round := range rounds {
		total += scores[round]
	}
	return total
}

var scores = map[string]int{
	"A X": 1 + 3,
	"A Y": 2 + 6,
	"A Z": 3 + 0,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 1 + 6,
	"C Y": 2 + 0,
	"C Z": 3 + 3,
}
