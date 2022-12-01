package aoc

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Check panics if the error is not nil
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Insert an int into a slice of ints
func Insert(s []int, i, value int) []int {
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = value
	return s
}

// Delete an int from a slice of ints
func Delete(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// Atoi converts a string to an int and panics if it goes wrong
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// FileToBytes reads a file into an array of bytes
func FileToBytes(filePath string) []byte {
	bs, err := ioutil.ReadFile(filePath)
	Check(err)
	return bs
}

// FileToString reads a file into a string
func FileToString(filePath string) string {
	return string(FileToBytes(filePath))
}

// FileLinesToStringSlice reads the lines of a file into a slice of strings
func FileLinesToStringSlice(filePath string) []string {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var strs []string
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	Check(scanner.Err())
	return strs
}

// FileLinesToIntSlice reads the lines of a file into a slice of ints
func FileLinesToIntSlice(filePath string) []int {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var ints []int
	for scanner.Scan() {
		ints = append(ints, Atoi(scanner.Text()))
	}
	Check(scanner.Err())
	return ints
}

// FileToIntSlice reads a CSV file into a slice of ints
func FileToIntSlice(filePath string) []int {
	return ToIntSlice(strings.TrimSpace(FileToString(filePath)), ",")
}

// ToIntSlice reads CSV input into a slice of ints
func ToIntSlice(input, sep string) []int {
	var ints []int
	for _, s := range strings.Split(input, sep) {
		if s == "" {
			continue
		}
		ints = append(ints, Atoi(s))
	}
	return ints
}

// To3DIntSlice reads a series of empty-line separated 2D series into a slice
func To3DIntSlice(lines []string, sep string) [][][]int {
	is3d := [][][]int{}
	is2d := [][]int{}
	for _, line := range lines {
		if len(line) == 0 {
			if len(is2d) > 0 {
				is3d = append(is3d, is2d)
				is2d = [][]int{}
			}
			continue
		}
		is := ToIntSlice(line, sep)
		if len(is) > 0 {
			is2d = append(is2d, is)
		}
	}
	if len(is2d) > 0 {
		is3d = append(is3d, is2d)
	}
	return is3d
}

// FileLinesToIntSliceSlice reads a file where each line is CSV into a slice of slices of ints
func FileLinesToIntSliceSlice(filePath string) [][]int {
	f, err := os.Open(filePath)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var ints [][]int
	for scanner.Scan() {
		var xs []int
		for _, s := range strings.Split(scanner.Text(), ",") {
			xs = append(xs, Atoi(s))
		}
		ints = append(ints, xs)
	}
	Check(scanner.Err())
	return ints
}
