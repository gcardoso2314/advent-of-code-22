package main

import (
	"fmt"
	"sort"
	"strconv"

	"advent-of-code-22/utils"
)

func A() {
	lines, _ := utils.ReadLines("input.txt")
	var current_calorie_count int
	var max_calories int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			// check if new count is greater than old max
			if current_calorie_count > max_calories {
				max_calories = current_calorie_count
			}
			// reset counter
			current_calorie_count = 0
		} else {
			count_to_add, _ := strconv.Atoi(lines[i])
			current_calorie_count += count_to_add
		}
	}

	// final check in case the last elf has most calories
	if current_calorie_count > max_calories {
		max_calories = current_calorie_count
	}

	fmt.Printf("Solution for part 1 is %d.", max_calories)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	var elf_calories []int
	var current_calorie_count int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			elf_calories = append(elf_calories, current_calorie_count)
			current_calorie_count = 0
		} else {
			count_to_add, _ := strconv.Atoi(lines[i])
			current_calorie_count += count_to_add
		}
	}
	// append last elf
	elf_calories = append(elf_calories, current_calorie_count)

	// sort and find sum of last 3
	sort.Ints(elf_calories)
	n := len(elf_calories)
	sum := elf_calories[n-1] + elf_calories[n-2] + elf_calories[n-3]

	fmt.Printf("Solution for part 2 is %d.", sum)
}

func main() {
	A()
	B()
}
