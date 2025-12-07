package utils

import (
	"bufio"
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
