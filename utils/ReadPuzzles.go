package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// return all lines
func ReadPuzzle(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var puzzle []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		puzzle = append(puzzle, scanner.Text())
	}
	return puzzle, scanner.Err()
}

// return content split by comma
func ReadPuzzleComma(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(strings.TrimSpace(string(content)), ",")

	return parts, nil
}

// return lines as ints
func ReadPuzzleInt(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var puzzle []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err // Return error if a line is not a number
		}
		puzzle = append(puzzle, number)
	}
	return puzzle, scanner.Err()
}

// Read grid and convert symbols
func ReadGrid[T any](path string, convert func(byte) T) ([][]T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]T
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]T, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = convert(line[i])
		}
		grid = append(grid, row)
	}
	return grid, scanner.Err()
}

func ReadGridSplitBySpaces(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Fields(line)
		grid = append(grid, row)
	}
	return grid, scanner.Err()
}

func PrintGrid[T any](grid [][]T) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(grid[i][j])
		}
		fmt.Println()
	}
}
