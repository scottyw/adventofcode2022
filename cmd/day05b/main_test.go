package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//     [D]
// [N] [C]
// [Z] [M] [P]
//  1   2   3

// move 1 from 2 to 1
// move 3 from 1 to 3
// move 2 from 2 to 1
// move 1 from 1 to 2

func TestProcess(t *testing.T) {
	stacks := [][]byte{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	directions := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
	assert.Equal(t, "MCD", process(stacks, directions))
}
