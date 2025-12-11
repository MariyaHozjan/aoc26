package main

import (
	"AOC26/utils"
	"fmt"
	"log"
)

type Spot string

const (
	Empty          Spot = "."
	Roll           Spot = "@"
	AccessibleRoll Spot = "X"
)

func spotFromChar(c byte) Spot {
	if c == '.' {
		return Empty
	}
	return Roll
}

func ReadPuzzleSpots(path string) ([][]Spot, error) {
	return utils.ReadGrid(path, spotFromChar)
}

func CheckSurroundingRolls(grid [][]Spot, r, c int) int {
	rows := len(grid)
	cols := len(grid[r])

	count := 0

	if r > 0 {
		if c > 0 && grid[r-1][c-1] == Roll {
			count++
		}
		if grid[r-1][c] == Roll {
			count++
		}
		if c < cols-1 && grid[r-1][c+1] == Roll {
			count++
		}
	}

	if c > 0 && grid[r][c-1] == Roll {
		count++
	}
	if c < cols-1 && grid[r][c+1] == Roll {
		count++
	}

	if r < rows-1 {
		if c > 0 && grid[r+1][c-1] == Roll {
			count++
		}
		if grid[r+1][c] == Roll {
			count++
		}
		if c < cols-1 && grid[r+1][c+1] == Roll {
			count++
		}
	}

	return count
}

func cloneGrid(grid [][]Spot) [][]Spot {
	cloned := make([][]Spot, len(grid))
	for i := range grid {
		cloned[i] = make([]Spot, len(grid[i]))
		copy(cloned[i], grid[i])
	}
	return cloned
}

func main() {
	spots, err := ReadPuzzleSpots("Day4/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	removedRolls := 0

	for {
		removedRollsThisRound := 0

		next := cloneGrid(spots)

		for r, row := range spots {
			for c, column := range row {
				if column == Roll {
					foundRolls := CheckSurroundingRolls(spots, r, c)
					if foundRolls < 4 {
						next[r][c] = Empty
						removedRollsThisRound++
					}
				}
			}
		}
		if removedRollsThisRound == 0 {
			break
		}
		removedRolls += removedRollsThisRound
		spots = next
	}

	utils.PrintGrid(spots)
	fmt.Println(removedRolls)
}
