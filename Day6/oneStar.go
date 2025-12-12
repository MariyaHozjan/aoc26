package main

import (
	"AOC26/utils"
	"fmt"
	"strconv"
)

func oneStar() {
	grid, err := utils.ReadGridSplitBySpaces("Day6/puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}

	allAnswers := 0

	/*
		// Print first row
		for row := 0; row < len(grid); row++ {
			fmt.Println(grid[row][0])
		}
	*/

	for i := 0; i < len(grid[0]); i++ { //len(grid[0])
		first, err := strconv.Atoi(grid[0][i])
		fmt.Printf("First answer: %d\n", first)
		if err != nil {
			fmt.Println(err)
		}
		second, err := strconv.Atoi(grid[1][i])
		if err != nil {
			fmt.Println(err)
		}
		third, err := strconv.Atoi(grid[2][i])
		if err != nil {
			fmt.Println(err)
		}
		fourth, err := strconv.Atoi(grid[3][i])
		if err != nil {
			fmt.Println(err)
		}
		if grid[4][i] == "+" {
			allAnswers += first + second + third + fourth
		} else {
			allAnswers += first * second * third * fourth
		}
	}
	fmt.Println(allAnswers)
}
