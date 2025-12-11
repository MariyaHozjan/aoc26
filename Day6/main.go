package main

import (
	"AOC26/utils"
	"fmt"
	"strconv"
)

func main() {
	grid, err := utils.ReadGridSplitBySpaces("Day6/puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}

	allAnswers := 0

	for i := 0; i < len(grid[0]); i++ {
		//fmt.Println(grid[i])

		first, err := strconv.Atoi(grid[0][i])
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(first)
		second, err := strconv.Atoi(grid[1][i])
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(second)
		third, err := strconv.Atoi(grid[2][i])
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(third)
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
