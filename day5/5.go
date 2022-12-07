package main

import (
	"advent-of-code-22/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseText(lines []string) ([][]string, [][]int) {
	var separatorIndex int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			separatorIndex = i
			break
		}
	}

	stackNumbers := lines[separatorIndex-1]
	stackNumbers = strings.TrimSpace(stackNumbers)
	numStacks, _ := strconv.Atoi(stackNumbers[len(stackNumbers)-1:])
	numMoves := len(lines) - separatorIndex

	stacks := make([][]string, numStacks)
	moves := make([][]int, numMoves)

	// parsing the stacks
	for i := 0; i < separatorIndex-1; i++ {
		for j := 0; 4*j+1 < len(lines[i]); j++ {
			if lines[i][4*j+1:4*j+2] != " " {
				stacks[j] = append(stacks[j], lines[i][4*j+1:4*j+2])
			}
		}
	}

	// parsing the moves
	re, _ := regexp.Compile(`move ([0-9]{1,2}) from ([0-9]{1,2}) to ([0-9]{1,2})`)
	for i := separatorIndex + 1; i < len(lines); i++ {
		matches := re.FindAllStringSubmatch(lines[i], -1)
		for _, match := range matches[0][1:] {
			num, _ := strconv.Atoi(match)
			moves[i-separatorIndex-1] = append(moves[i-separatorIndex-1], num)
		}
	}

	return stacks, moves
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	stacks, moves := parseText(lines)

	for _, move := range moves {
		if len(move) <= 0 {
			continue
		}
		numMoves := move[0]
		from := move[1] - 1
		to := move[2] - 1
		for i := 0; i < numMoves; i++ {
			x := stacks[from][0]
			stacks[from] = stacks[from][1:]
			stacks[to] = append([]string{x}, stacks[to]...)
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%s", stack[0])
	}
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	stacks, moves := parseText(lines)

	for _, move := range moves {
		if len(move) <= 0 {
			continue
		}
		numMoves := move[0]
		from := move[1] - 1
		to := move[2] - 1

		x := make([]string, numMoves)
		copy(x, stacks[from])
		stacks[from] = stacks[from][numMoves:]
		stacks[to] = append(x, stacks[to]...)
	}

	for _, stack := range stacks {
		fmt.Printf("%s", stack[0])
	}
}

func main() {
	B()
}
