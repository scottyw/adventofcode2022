package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, consider the following list of section assignment pairs:

// 2-4,6-8
// 2-3,4-5
// 5-7,7-9
// 2-8,3-7
// 6-6,4-6
// 2-6,4-8

func TestCountOverlaps(t *testing.T) {
	assignments := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	assert.Equal(t, 4, countOverlaps(assignments))
}
