package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() ([]int, []int) {
	dat, err := os.ReadFile("input.txt")
	check(err)
	
	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, l := range lines {
		pairs := strings.Split(l, " ")
		if num, err := strconv.Atoi(pairs[0]); err == nil && num != 0 {
			left[i] = num
		}
		if num, err := strconv.Atoi(pairs[len(pairs) - 1]); err == nil && num != 0 {
			right[i] = num
		}
	}
	return left, right
}

func partOne(left []int, right []int) float64 {
	slices.Sort(left)
	slices.Sort(right)
	var distance float64
	for i := range left {
		sep := float64(left[i] - right[i])
		distance += math.Abs(sep)
	}
	return distance
}

func partTwo(left []int, right []int) int {
	counts := map[int]int{}
	for _, v := range right {
		counts[v] += 1
	}
	total := 0
	for _, v := range left {
		total += (v * counts[v])
	}
	return total
}

func main() {
	fmt.Println("AoC day 1")
	left, right := parseInput()
	fmt.Printf("part 1: %.0f\n", partOne(left, right))
	fmt.Printf("part 2: %d\n", partTwo(left, right))
}
