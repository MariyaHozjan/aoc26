package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCharGrid(path string) ([][]rune, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines [][]rune
	maxLen := 0

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		r := []rune(sc.Text())
		if len(r) > maxLen {
			maxLen = len(r)
		}
		lines = append(lines, r)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	grid := make([][]rune, len(lines))
	for i := range lines {
		row := make([]rune, maxLen)
		for j := 0; j < maxLen; j++ {
			if j < len(lines[i]) {
				row[j] = lines[i][j]
			} else {
				row[j] = ' '
			}
		}
		grid[i] = row
	}
	return grid, nil
}

func isBlankColumn(grid [][]rune, col int) bool {
	for r := 0; r < len(grid); r++ {
		if grid[r][col] != ' ' {
			return false
		}
	}
	return true
}

func Sum(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func Multiply(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	total := 1
	for i := 0; i < len(nums); i++ {
		total *= nums[i]
	}
	return total
}

func main() {
	grid, err := readCharGrid("Day6/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}

	finalAnswer := 0
	c := len(grid[0]) - 1

	for c >= 0 {
		if isBlankColumn(grid, c) {
			c--
			continue
		}

		var numbers []int
		operation := '?'

		for c >= 0 && !isBlankColumn(grid, c) {
			current := ""

			for r := 0; r < 4; r++ {
				ch := grid[r][c]
				if ch != ' ' {
					current += string(ch)
				}
			}

			if current != "" {
				val, err := strconv.Atoi(current)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, val)
			}

			if grid[4][c] == '+' || grid[4][c] == '*' {
				operation = grid[4][c]
			}

			c--
		}

		if operation == '+' {
			finalAnswer += Sum(numbers)
		} else {
			finalAnswer += Multiply(numbers)
		}
	}

	fmt.Println(finalAnswer)
}
