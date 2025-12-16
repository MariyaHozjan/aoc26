package main

import (
	"AOC26/utils"
	"fmt"
	"log"
)

type Directions string

const (
	Empty Directions = "."
	Split Directions = "^"
	Start Directions = "S"
	Beam  Directions = "|"
)

func directionFromChar(c byte) Directions {
	switch c {
	case '.':
		return Empty
	case 'S':
		return Start
	case '^':
		return Split
	case '|':
		return Beam
	default:
		return Empty
	}
}

type Location struct {
	row int
	col int
}

func ReadPuzzleDirection(path string) ([][]Directions, error) {
	return utils.ReadGrid(path, directionFromChar)
}

func MoveOneDown(loc Location) Location {
	return Location{row: loc.row + 1, col: loc.col}
}

var BeamsToCheck []Location

func CountAlternateTimelines(puzzle [][]Directions, location Location, memo map[Location]int) int {
	if v, ok := memo[location]; ok {
		return v
	}

	next := MoveOneDown(location)

	if next.row >= len(puzzle) {
		memo[location] = 1
		return 1
	}

	cell := puzzle[next.row][next.col]

	switch cell {
	case Empty, Start:
		ways := CountAlternateTimelines(puzzle, next, memo)
		memo[location] = ways
		return ways

	case Split:
		left := Location{row: next.row, col: next.col - 1}
		right := Location{row: next.row, col: next.col + 1}

		ways := 0
		ways += CountAlternateTimelines(puzzle, left, memo)
		ways += CountAlternateTimelines(puzzle, right, memo)

		memo[location] = ways
		return ways

	default:
		memo[location] = 0
		return 0
	}
}

func main() {
	puzzle, err := ReadPuzzleDirection("Day7/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	startColumn := (len(puzzle[0]) - 1) / 2
	start := Location{row: 1, col: startColumn}

	memo := make(map[Location]int)

	BeamsToCheck = append(BeamsToCheck, start)

	finalSplitCount := 0
	alternateTimelineCount := CountAlternateTimelines(puzzle, start, memo)

	for i := 0; i < len(BeamsToCheck); i++ {
		beam := BeamsToCheck[i]
		next := MoveOneDown(beam)
		if next.row >= len(puzzle) {
			continue
		}

		nextCell := puzzle[next.row][next.col]

		if nextCell == Empty {
			puzzle[next.row][next.col] = Beam
			BeamsToCheck = append(BeamsToCheck, next)
			continue
		}

		if nextCell == Split {
			finalSplitCount++

			left := Location{row: next.row, col: next.col - 1}
			right := Location{row: next.row, col: next.col + 1}

			puzzle[left.row][left.col] = Beam
			puzzle[right.row][right.col] = Beam

			BeamsToCheck = append(BeamsToCheck, left, right)
			continue
		}
	}

	fmt.Println(finalSplitCount)
	fmt.Println(alternateTimelineCount)
}
