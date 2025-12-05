package main

import (
	"AOC26/utils"
	"fmt"
	"strconv"
)

func main() {
	var puzzle, err = utils.ReadPuzzle("Day1/puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}

	var currentNumber = 50
	var password = 0

	for x := 0; x < len(puzzle); x++ {
		var currentTurn = puzzle[x]
		letter := string(currentTurn[0])
		numStr := string(currentTurn[1:])
		number, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		for y := 0; y < number; y++ {
			if letter == "R" {
				if currentNumber == 99 {
					currentNumber = 0
					password++
				} else {
					currentNumber++
				}
			}
			if letter == "L" {
				if currentNumber == 0 {
					currentNumber = 99
				} else {
					currentNumber--
				}

				if currentNumber == 0 {
					password++
				}
			}
		}

		/*
			if currentNumber == 0 {
				password++
				//fmt.Printf(letter)
				//fmt.Println(number)
			}
		*/
	}
	fmt.Println(password)
}
