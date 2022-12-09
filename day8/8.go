package main

import (
	"advent-of-code-22/utils"
	"fmt"
	"strconv"
)

func createForest(lines []string) [][]int {
	forest := make([][]int, len(lines))
	for i, line := range lines {
		var treeLine []int
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			treeLine = append(treeLine, num)
		}
		forest[i] = append(forest[i], treeLine...)
	}

	return forest
}

func isVisible(array []int, max int) bool {
	for _, a := range array {
		if a >= max {
			return false
		}
	}
	return true
}

func calculateViewingDistance(array []int, max int) int {
	// assumes that the tree we're measuring is at index "-1" of the array
	// that way I can iterate left to right and break
	for i, a := range array {
		if a >= max {
			return i + 1
		}
	}

	return len(array)
}

func A() {
	lines, _ := utils.ReadLines("input.txt")
	forest := createForest(lines)

	width := len(forest[0])
	height := len(forest)
	visibleTrees := width*2 + 2*(height-2)

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			treeH := forest[i][j]
			//left
			left_side := forest[i][:j]
			//up
			var up_side []int
			for n := 0; n < i; n++ {
				up_side = append(up_side, forest[n][j])
			}
			//right
			right_side := forest[i][j+1:]
			//down
			var down_side []int
			for n := i + 1; n < height; n++ {
				down_side = append(down_side, forest[n][j])
			}
			if isVisible(left_side, treeH) || isVisible(up_side, treeH) || isVisible(right_side, treeH) || isVisible(down_side, treeH) {
				visibleTrees++
			}
		}
	}
	fmt.Println(visibleTrees)
}

func B() {
	lines, _ := utils.ReadLines("input.txt")
	forest := createForest(lines)

	width := len(forest[0])
	height := len(forest)

	var maxScenicScore int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			treeH := forest[i][j]
			//left
			var left_side []int
			for n := j - 1; n >= 0; n-- {
				left_side = append(left_side, forest[i][n])
			}
			left_score := calculateViewingDistance(left_side, treeH)
			//up
			var up_side []int
			for n := i - 1; n >= 0; n-- {
				up_side = append(up_side, forest[n][j])
			}
			up_score := calculateViewingDistance(up_side, treeH)
			//right
			right_side := forest[i][j+1:]
			right_score := calculateViewingDistance(right_side, treeH)
			//down
			var down_side []int
			for n := i + 1; n < height; n++ {
				down_side = append(down_side, forest[n][j])
			}
			down_score := calculateViewingDistance(down_side, treeH)

			if left_score*up_score*right_score*down_score > maxScenicScore {
				maxScenicScore = left_score * up_score * right_score * down_score
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func main() {
	B()
}
