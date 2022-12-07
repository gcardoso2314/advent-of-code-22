package main

import (
	"fmt"

	"advent-of-code-22/utils"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func calculateRoundScore(opponent string, response string) int {
	var score int
	oppArray := [3]string{"A", "B", "C"}
	respArray := [3]string{"X", "Y", "Z"}

	oppIndex := -1
	for i := 0; i < len(oppArray); i++ {
		if opponent == oppArray[i] {
			oppIndex = i
		}
	}

	respIndex := -1
	for i := 0; i < len(respArray); i++ {
		if response == respArray[i] {
			respIndex = i
		}
	}

	// determine if win
	if oppIndex == respIndex {
		score += 3
	} else {
		if respIndex == mod(oppIndex+1, 3) {
			score += 6
		}
	}

	// score for the value you played
	switch response {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	return score
}

func calculateResponse(opponent string, result string) string {
	oppArray := [3]string{"A", "B", "C"}
	respArray := [3]string{"X", "Y", "Z"}

	oppIndex := -1
	for i := 0; i < len(oppArray); i++ {
		if opponent == oppArray[i] {
			oppIndex = i
		}
	}

	if result == "Y" {
		return respArray[oppIndex]
	} else if result == "X" {
		return respArray[mod(oppIndex-1, 3)]
	}

	return respArray[mod(oppIndex+1, 3)]
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	fmt.Printf("Trial %d", -1%3)
	var totalScore int
	for i := 0; i < len(lines); i++ {
		opponent := string(lines[i][0])
		response := string(lines[i][2])
		totalScore += calculateRoundScore(opponent, response)
	}

	fmt.Printf("Solution for part 1 is %d.", totalScore)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	var totalScore int
	for i := 0; i < len(lines); i++ {
		opponent := string(lines[i][0])
		result := string(lines[i][2])
		response := calculateResponse(opponent, result)
		totalScore += calculateRoundScore(opponent, response)
	}

	fmt.Printf("Solution for part 2 is %d.", totalScore)
}

func main() {
	A()
	B()
}
