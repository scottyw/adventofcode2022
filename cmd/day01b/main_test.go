package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, suppose the Elves finish writing their items' Calories and end up with the following list:

// 1000
// 2000
// 3000

// 4000

// 5000
// 6000

// 7000
// 8000
// 9000

// 10000

func TestFindMaximum(t *testing.T) {
	foods := [][]string{
		{
			"1000",
			"2000",
			"3000",
		},
		{
			"4000",
		},
		{
			"5000",
			"6000",
		},
		{
			"7000",
			"8000",
			"9000",
		},
		{
			"10000",
		}}
	assert.Equal(t, 45000, findMaximum(foods))
}
