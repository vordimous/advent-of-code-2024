package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseInput() string {
	bytes, err := os.ReadFile("input.txt")
	check(err)
	return strings.TrimSpace(string(bytes))
}

func findAllValidMuls(body string) [][]string{
	reg := `(don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\))`
	r, err := regexp.Compile(reg)
	check(err)
	numbs := r.FindAllStringSubmatch(body, -1)
	fmt.Println("numbs", numbs)
	return numbs
}


func main() {
	muls := findAllValidMuls(parseInput())
	
	total := 0
	do := true
	for _, m := range muls {
		println(m[0])
		if strings.HasPrefix(m[0], "don't") {
			println("don't")
			do = false
		} else if strings.HasPrefix(m[0], "do") {
			println("do")
			do = true
		}
		if do && strings.HasPrefix(m[0], "mul") {
			a, err := strconv.Atoi(m[2])
			check(err)
			b, err := strconv.Atoi(m[3])
			check(err)
			fmt.Printf("%d * %d = %d \n", a, b, a * b )
			total += a * b
		}
	}
	println("total", total)
}
