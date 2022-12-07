package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code-22/utils"
)

func findSectionsCleaned(sections string) map[int]bool {
	limits := strings.Split(sections, "-")
	start, _ := strconv.Atoi(limits[0])
	end, _ := strconv.Atoi(limits[1])

	sections_cleaned := make(map[int]bool)
	for i := start; i <= end; i++ {
		sections_cleaned[i] = true
	}

	return sections_cleaned
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	var countOverlap int

	for i := 0; i < len(lines); i++ {
		// create a counter of the sections for both elves
		sections := strings.Split(string(lines[i]), ",")

		elf1 := findSectionsCleaned(sections[0])
		elf2 := findSectionsCleaned(sections[1])

		overlapping := true
		if len(elf1) <= len(elf2) {
			for section := range elf1 {
				if !elf2[section] {
					overlapping = false
					break
				}
			}
		} else {
			for section := range elf2 {
				if !elf1[section] {
					overlapping = false
					break
				}
			}
		}
		if overlapping {
			countOverlap++
		}
	}
	fmt.Println(countOverlap)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	var countOverlap int

	for i := 0; i < len(lines); i++ {
		// create a counter of the sections for both elves
		sections := strings.Split(string(lines[i]), ",")

		elf1 := findSectionsCleaned(sections[0])
		elf2 := findSectionsCleaned(sections[1])

		overlapping := false
		if len(elf1) <= len(elf2) {
			for section := range elf1 {
				if elf2[section] {
					overlapping = true
					break
				}
			}
		} else {
			for section := range elf2 {
				if elf1[section] {
					overlapping = true
					break
				}
			}
		}
		if overlapping {
			countOverlap++
		}
	}
	fmt.Println(countOverlap)
}

func main() {
	A()
	B()
}
