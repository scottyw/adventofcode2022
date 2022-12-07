package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For example, suppose you have the following list of contents from six rucksacks:

// vJrwpWtwJgWrhcsFMMfFFhFp
// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
// PmmdzqPrVvPwwTWBwg
// wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
// ttgJtRGJQctTZtZT
// CrZsJsPPZsGzwwsLwLmpwMDw

func TestFindScore(t *testing.T) {
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	assert.Equal(t, 70, calculateTotal(rucksacks))
}

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 18, findCommon("vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"))
	assert.Equal(t, 52, findCommon("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"))
}

func TestCalulcatePriority(t *testing.T) {
	assert.Equal(t, 16, findPriority('p'))
	assert.Equal(t, 38, findPriority('L'))
	assert.Equal(t, 42, findPriority('P'))
	assert.Equal(t, 22, findPriority('v'))
	assert.Equal(t, 20, findPriority('t'))
	assert.Equal(t, 19, findPriority('s'))
}
