package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() [][]int {
	dat, err := os.ReadFile("input.txt")
	check(err)

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	parsed := make([][]int, len(lines))
	for i, l := range lines {
		s := strings.Split(l, " ")
		nums := []int{}

		for _, n := range s {
			if num, err := strconv.Atoi(n); err == nil {
				nums = append(nums, num)
			}
		}
		parsed[i] = nums
	}
	return parsed
}

func partOne(m [][]int) int {
	total := 0
	for _, r := range m {
		safe := true
		pos := (r[0] - r[1]) >= 0
		for i, n := range r[:len(r)-1] {
			diff := n - r[i+1]
			distance := math.Abs(float64(diff))
			if (diff >= 0) != pos || distance == 0 || distance > 3 {
				safe = false
				break
			}
		}
		if safe {
			total += 1
		}
	}

	return total
}

func removeIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func partTwo(m [][]int) int {
	total := 0
	for _, r := range m {
		safe := true
		pos := (r[0] - r[1]) >= 0
		for i, n := range r[:len(r)-1] {
			diff := n - r[i+1]
			distance := math.Abs(float64(diff))
			if (diff >= 0) != pos || distance == 0 || distance > 3 {
				safe = false
				break
			}
		}
		// check if initial report is safe
		// and make all possible copies of a report with one level removed
		stripped := make([][]int, len(r))
		for i, n := range r {
			//check
			if i < len(r)-1 {
				diff := n - r[i+1]
				distance := math.Abs(float64(diff))
				if (diff >= 0) != pos || distance == 0 || distance > 3 {
					safe = false
				}
			}
			// copy
			stripped[i] = make([]int, len(r))
			copy(stripped[i], r)
			stripped[i] = removeIndex(stripped[i], i)
		}
		if !safe {
			// searching for one good level means we are good
			for _, s := range stripped {
				safe = true
				pos := (s[0] - s[len(s)-1]) >= 0
				for i, n := range s[:len(s)-1] {
					diff := n - s[i+1]
					distance := math.Abs(float64(diff))
					if (diff >= 0) != pos || distance == 0 || distance > 3 {
						safe = false
						break
					}
				}
				if safe {
					break
				}
			}
		}

		if safe {
			total += 1
		}
	}

	return total
}

func main() {
	fmt.Println("AoC day 2")
	m := parseInput()
	fmt.Printf("part 1: %v\n", partOne(m))
	fmt.Printf("part 2: %v\n", partTwo(m))
}
