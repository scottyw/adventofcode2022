package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, suppose you were given the following strategy guide:

// A Y
// B X
// C Z

func TestFindScore(t *testing.T) {
	rounds := []string{
		"A Y",
		"B X",
		"C Z",
	}
	assert.Equal(t, 12, findScore(rounds))
}
