package main

import (
	"advent-of-code-22/utils"
	"fmt"
)

func contains(a []string, c string) bool {
	for _, s := range a {
		if s == c {
			return true
		}
	}
	return false
}

func findStart(line string, numDistinct int) int {
	for i := numDistinct; i < len(line); i++ {
		charCounter := make(map[rune]int)
		start := true
		for _, c := range line[i-numDistinct : i] {
			charCounter[c]++
		}
		for _, val := range charCounter {
			if val > 1 {
				start = false
			}
		}

		if start {
			return i
		}
	}
	return -1
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	for i := 0; i < len(lines); i++ {
		fmt.Println(findStart(lines[i], 4))
	}
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	for i := 0; i < len(lines); i++ {
		fmt.Println(findStart(lines[i], 14))
	}
}

func main() {
	A()
	B()
}
