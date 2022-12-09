package main

import (
	"fmt"
	"regexp"

	"github.com/scottyw/adventofcode2022/pkg/aoc"
)

// [B]                     [N]     [H]
// [V]         [P] [T]     [V]     [P]
// [W]     [C] [T] [S]     [H]     [N]
// [T]     [J] [Z] [M] [N] [F]     [L]
// [Q]     [W] [N] [J] [T] [Q] [R] [B]
// [N] [B] [Q] [R] [V] [F] [D] [F] [M]
// [H] [W] [S] [J] [P] [W] [L] [P] [S]
// [D] [D] [T] [F] [G] [B] [B] [H] [Z]
//  1   2   3   4   5   6   7   8   9

// move 2 from 8 to 1
var r = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func main() {
	stacks := [][]byte{
		{'D', 'H', 'N', 'Q', 'T', 'W', 'V', 'B'},
		{'D', 'W', 'B'},
		{'T', 'S', 'Q', 'W', 'J', 'C'},
		{'F', 'J', 'R', 'N', 'Z', 'T', 'P'},
		{'G', 'P', 'V', 'J', 'M', 'S', 'T'},
		{'B', 'W', 'F', 'T', 'N'},
		{'B', 'L', 'D', 'Q', 'F', 'H', 'V', 'N'},
		{'H', 'P', 'F', 'R'},
		{'Z', 'S', 'M', 'B', 'L', 'N', 'P', 'H'},
	}
	directions := aoc.FileLinesToStringSlice("input/05.txt")
	fmt.Println(process(stacks, directions))
}

func process(stacks [][]byte, directions []string) string {

	// printStacks(stacks)

	for _, direction := range directions {

		matches := r.FindStringSubmatch(direction)
		count := aoc.Atoi(matches[1])
		from := aoc.Atoi(matches[2]) - 1
		to := aoc.Atoi(matches[3]) - 1

		fmt.Printf("move %d from %d to %d\n", count, from, to)

		for count > 0 {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
			count--
		}

		// printStacks(stacks)

	}

	var message string
	for _, stack := range stacks {
		message += string(stack[(len(stack) - 1)])
	}

	return message

}

func printStacks(stacks [][]byte) {
	for _, stack := range stacks {
		for _, item := range stack {
			fmt.Printf("%s ", string(item))
		}
		fmt.Println()
	}
	fmt.Println()
}
