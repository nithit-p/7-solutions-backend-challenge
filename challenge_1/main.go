package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type inputFile [][]int

func main() {
	var hardInput inputFile
	file, err := os.ReadFile("../files/hard.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &hardInput); err != nil {
		panic(err)
	}
	basicInput := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	fmt.Println("Basic :", maxPathSumLoop(basicInput))
	fmt.Println("Hard :", maxPathSumLoop(hardInput))
}

func maxPathSumRecursive(trees [][]int, level int, nodeIndex int) int {
	node := trees[level][nodeIndex]
	if level == len(trees)-1 {
		return node
	}
	left := maxPathSumRecursive(trees, level+1, nodeIndex)
	right := maxPathSumRecursive(trees, level+1, nodeIndex+1)
	if left > right {
		return left + node
	}
	return right + node
}

func maxPathSumLoop(trees [][]int) int {
	maxSumTrees := make([][]int, len(trees))
	for i := len(maxSumTrees) - 1; i >= 0; i-- {
		maxSumTrees[i] = make([]int, len(trees[i]))
		copy(maxSumTrees[i], trees[i])
		if i == len(maxSumTrees)-1 {
			continue
		}
		for j := 0; j < len(maxSumTrees[i]); j++ {
			maxSumTrees[i][j] += max(maxSumTrees[i+1][j], maxSumTrees[i+1][j+1])
		}
	}
	// fmt.Println(maxSumTrees)
	return maxSumTrees[0][0]
}
