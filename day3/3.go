package main

import (
	"fmt"
	"strings"

	"advent-of-code-22/utils"
)

func isUpper(char string) bool {
	return strings.ToUpper(char) == char
}

func calculatePriority(char rune) int {
	if isUpper(string(char)) {
		return int(char - 38)
	}

	return int(char - 96)
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	var totalSum int
	for i := 0; i < len(lines); i++ {
		n := len(lines[i]) / 2
		compartment_1 := lines[i][:n]
		compartment_2 := lines[i][n:]

		var common_items []rune
		for _, c1 := range compartment_1 {
			for _, c2 := range compartment_2 {
				if c1 == c2 {
					common_items = append(common_items, c1)
				}
			}
		}

		totalSum += calculatePriority(common_items[0])

	}
	fmt.Printf("The solution to part 1 is %d\n", totalSum)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	var totalSum int
	for i := 0; i < len(lines)/3; i++ {
		char_counts := make(map[rune]int)

		elf_1_chars := make(map[rune]bool)
		for _, c := range lines[3*i] {
			if !elf_1_chars[c] {
				char_counts[c]++
				// mark as seen so we don't double count it
				elf_1_chars[c] = true
			}
		}

		elf_2_chars := make(map[rune]bool)
		for _, c := range string(lines[3*i+1]) {
			if !elf_2_chars[c] {
				char_counts[c]++
				// mark as seen so we don't double count it
				elf_2_chars[c] = true
			}
		}

		elf_3_chars := make(map[rune]bool)
		for _, c := range string(lines[3*i+2]) {
			if !elf_3_chars[c] {
				char_counts[c]++
				// mark as seen so we don't double count it
				elf_3_chars[c] = true
			}
		}

		for char, val := range char_counts {
			if val == 3 {
				totalSum += calculatePriority(char)
			}
		}
	}
	fmt.Printf("the solution to part 2 is %d", totalSum)
}

func main() {
	A()
	B()
}
