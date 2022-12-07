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
	assert.Equal(t, 157, calculateTotal(rucksacks))
}

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 16, findDuplicate("vJrwpWtwJgWrhcsFMMfFFhFp"))
	assert.Equal(t, 38, findDuplicate("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"))
	assert.Equal(t, 42, findDuplicate("PmmdzqPrVvPwwTWBwg"))
	assert.Equal(t, 22, findDuplicate("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"))
	assert.Equal(t, 20, findDuplicate("ttgJtRGJQctTZtZT"))
	assert.Equal(t, 19, findDuplicate("CrZsJsPPZsGzwwsLwLmpwMDw"))
}

func TestCalulcatePriority(t *testing.T) {
	assert.Equal(t, 16, findPriority('p'))
	assert.Equal(t, 38, findPriority('L'))
	assert.Equal(t, 42, findPriority('P'))
	assert.Equal(t, 22, findPriority('v'))
	assert.Equal(t, 20, findPriority('t'))
	assert.Equal(t, 19, findPriority('s'))
}
